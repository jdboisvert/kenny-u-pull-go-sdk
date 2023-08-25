package kennyupullgosdk

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type InventoryListing struct {
	Year       string
	Make       string
	Model      string
	DateListed string
	RowID      string
	Branch     string
	ListingUrl string
}

type InventorySearch struct {
	Year   string
	Make   string
	Model  string
	Branch string
}

func getInventoryFromPage(inventorySearch *InventorySearch, page int16) ([]InventoryListing, error) {
	branch := "all-branches"
	if inventorySearch.Branch != "" {
		branch = inventorySearch.Branch
	}
	params := url.Values{
		"brand":                        {inventorySearch.Make}, // "brand" is the name of the form field on the website for Make (eg. Honda)
		"model":                        {inventorySearch.Model},
		"model_year":                   {inventorySearch.Year},
		"branch[]":                     {branch},
		"nb_items":                     {"42"},           // this is the number of items per page by default
		"input-select-sort-1621685231": {"New+arrivals"}, // Always search by new arrivals by default

	}
	url := fmt.Sprintf("https://kennyupull.com/auto-parts/our-inventory/page/%d/?%s", page, params.Encode())
	response, err := http.Get(url)

	inventoryListings := []InventoryListing{}
	if err != nil {
		return inventoryListings, fmt.Errorf("failed to get URL from Kenny U-Pull: %v", err)
	}
	defer response.Body.Close() // Close the response body when the function returns

	if response.StatusCode != http.StatusOK {
		return inventoryListings, fmt.Errorf("got non-200 status code: %d so returning None", response.StatusCode)
	}

	doc, _ := goquery.NewDocumentFromReader(response.Body)
	doc.Find(".product-info").Each(func(_ int, tag *goquery.Selection) {
		infosDateTag := tag.Find(".infos .infos--date .date.info")
		datedListed := strings.TrimSpace(infosDateTag.First().Text())
		rowID := strings.TrimSpace(infosDateTag.Last().Text())

		inventoryListing := InventoryListing{
			Year:       strings.TrimSpace(tag.Find(".title .year").Text()),
			Make:       strings.TrimSpace(tag.Find(".title .brand").Text()),
			Model:      strings.TrimSpace(tag.Find(".title .model").Text()),
			DateListed: datedListed,
			RowID:      rowID,
			Branch:     strings.TrimSpace(tag.Find(".infos .infos--branch .branch.info").Text()),
			ListingUrl: strings.TrimSpace(tag.Find(".btn-wrapper a").AttrOr("href", "")),
		}

		inventoryListings = append(inventoryListings, inventoryListing)
	})

	return inventoryListings, nil
}

func GetInventory(inventorySearch InventorySearch) ([]InventoryListing, error) {
	inventoryListings := []InventoryListing{}

	page := int16(1)
	for {
		inventoryListingsPage, err := getInventoryFromPage(&inventorySearch, page)
		if err != nil {
			return nil, err
		}

		inventoryListings = append(inventoryListings, inventoryListingsPage...)
		if len(inventoryListingsPage) == 0 {
			// No more pages to paginate through so exit loop
			break
		}

		page++
	}

	return inventoryListings, nil
}

func GetLatestListing(inventorySearch InventorySearch) (*InventoryListing, error) {
	inventoryListings, err := getInventoryFromPage(&inventorySearch, 1)
	if err != nil {
		return nil, err
	}

	if len(inventoryListings) == 0 {
		// No listings are present return nil
		return nil, nil
	}

	return &inventoryListings[0], nil
}
