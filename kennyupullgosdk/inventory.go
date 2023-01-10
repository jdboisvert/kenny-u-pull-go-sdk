package kennyupullgosdk

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type InventoryListing struct {
	Year       string
	Make       string
	Model      string
	DateListed string
	RowID      string
	Branch     string
}

type InventorySearch struct {
	Year   string
	Make   string
	Model  string
	Branch string
}

func GetInventory(inventorySearch InventorySearch) ([]InventoryListing, error) {
	branch := "all-branches"
	if inventorySearch.Branch != "" {
		branch = inventorySearch.Branch
	}
	params := url.Values{
		"brand":      {inventorySearch.Make}, // "brand" is the name of the form field on the website for Make (eg. Honda)
		"model":      {inventorySearch.Model},
		"model_year": {inventorySearch.Year},
		"branch[]":   {branch},
		"nb_items":   {"42"}, // this is the number of items per page by default

	}
	response, _ := http.Get("https://kennyupull.com/auto-parts/our-inventory/page/1/?" + params.Encode())
	// fmt.Println("HERE")
	// fmt.Println(response.StatusCode)
	// fmt.Println(response)

	// defer response.Body.Close()

	// if response.StatusCode == http.StatusOK {
	// 	bodyBytes, err := io.ReadAll(response.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bodyString := string(bodyBytes)

	// 	fmt.Println(bodyString)
	// }

	doc, _ := goquery.NewDocumentFromReader(response.Body)
	doc.Find(".product-info .title").Each(func(_ int, tag *goquery.Selection) {

		text := tag.Text()
		fmt.Printf("%s\n", text)
	})

	doc.Find(".product-info .infos").Each(func(_ int, tag *goquery.Selection) {

		text := tag.Text()
		fmt.Printf("%s\n", text)
	})

	return []InventoryListing{}, nil
}
