package kennyupullgosdk

import (
	"fmt"
	"testing"

	"github.com/h2non/gock"

	"github.com/stretchr/testify/assert"
)

const HTMLWithVehicleExample string = `
<div class="product-info">
	<div class="title">
		<h5>
			<span class="brand">LEXUS</span>
			<span class="model"> IS 250</span>
			<span class="year"> 2010</span>
		</h5>
	</div>
	<div class="infos">
		<div class="infos--branch">
			<p class="subtitle">Branch</p>
			<p class="branch info">Ottawa</p>
		</div>
		<div class="infos--date">
			<p class="subtitle">Date listed</p>
			<p class="date info">2023-01-06</p>
		</div>
		<div class="infos--date">
			<p class="subtitle">Row</p>
			<p class="date info">26A</p>
		</div>
	</div>
		<div class="btn-wrapper">
		<a href="https://kennyupull.com/part/lexus-is250-2010-93/"
			target="_self" class="btn-arrow is-red full-width-with-text">
			More Details
			<span class="picto reset-font">
			<svg xmlns="http://www.w3.org/2000/svg" width="9" height="15" viewBox="0 0 9 15">
			<g fill="none" fill-rule="evenodd" transform="translate(-8 -5)">
			<polygon points="0 0 20 0 20 20 0 20" />
			<polygon points="0 0 25 0 25 25 0 25" />
			<polyline stroke-linecap="round" stroke-linejoin="round" stroke-width="1.25" points="9.375 6.25 15.625 12.5 9.375 18.75" />
			</g>
			</svg>
			</span>
		</a>
	</div>
</div>
`

const HTMLWIthMultipleVehiclesExample string = `
<div class="product-info">
	<div class="title">
		<h5>
			<span class="brand">LEXUS</span>
			<span class="model"> IS 250</span>
			<span class="year"> 2010</span>
		</h5>
	</div>
	<div class="infos">
		<div class="infos--branch">
			<p class="subtitle">Branch</p>
			<p class="branch info">Ottawa</p>
		</div>
		<div class="infos--date">
			<p class="subtitle">Date listed</p>
			<p class="date info">2023-01-06</p>
		</div>
		<div class="infos--date">
			<p class="subtitle">Row</p>
			<p class="date info">26A</p>
		</div>
	</div>
		<div class="btn-wrapper">
		<a href="https://kennyupull.com/part/lexus-is250-2010-93/"
			target="_self" class="btn-arrow is-red full-width-with-text">
			More Details
			<span class="picto reset-font">
			<svg xmlns="http://www.w3.org/2000/svg" width="9" height="15" viewBox="0 0 9 15">
			<g fill="none" fill-rule="evenodd" transform="translate(-8 -5)">
			<polygon points="0 0 20 0 20 20 0 20" />
			<polygon points="0 0 25 0 25 25 0 25" />
			<polyline stroke-linecap="round" stroke-linejoin="round" stroke-width="1.25" points="9.375 6.25 15.625 12.5 9.375 18.75" />
			</g>
			</svg>
			</span>
		</a>
	</div>
</div>
<div class="product-info">
	<div class="title">
		<h5>
			<span class="brand">LEXUS</span>
			<span class="model"> IS 250</span>
			<span class="year">2006</span>
		</h5>
	</div>
	<div class="infos">
		<div class="infos--branch">
			<p class="subtitle">Branch</p>
			<p class="branch info">Ste-Sophie</p>
		</div>
		<div class="infos--date">
			<p class="subtitle">Date listed</p>
			<p class="date info">2023-01-01</p>
		</div>
		<div class="infos--date">
			<p class="subtitle">Row</p>
			<p class="date info">25B</p>
		</div>
	</div>
		<div class="btn-wrapper">
		<a href="https://kennyupull.com/part/lexus-is250-2006-1/"
			target="_self" class="btn-arrow is-red full-width-with-text">
			More Details
			<span class="picto reset-font">
			<svg xmlns="http://www.w3.org/2000/svg" width="9" height="15" viewBox="0 0 9 15">
			<g fill="none" fill-rule="evenodd" transform="translate(-8 -5)">
			<polygon points="0 0 20 0 20 20 0 20" />
			<polygon points="0 0 25 0 25 25 0 25" />
			<polyline stroke-linecap="round" stroke-linejoin="round" stroke-width="1.25" points="9.375 6.25 15.625 12.5 9.375 18.75" />
			</g>
			</svg>
			</span>
		</a>
	</div>
</div>
`

const HTMLWithoutVehicleExample string = `
<div>This is a test</div>
`

func MockPagingThroughInventory(numberOfPages int) {
	// Mock HTTP Request for N amount of pages
	for i := 1; i < numberOfPages; i++ {
		url := fmt.Sprintf("auto-parts/our-inventory/page/%d/", i)
		gock.New("https://kennyupull.com").
			Get(url).
			Reply(200).
			BodyString(HTMLWithVehicleExample)
	}

	// Ensure to have a page that has no items
	urlEnd := fmt.Sprintf("auto-parts/our-inventory/page/%d/", numberOfPages)
	gock.New("https://kennyupull.com").
		Get(urlEnd).
		Reply(200).
		BodyString(HTMLWithoutVehicleExample)
}

func MockPagingThroughInventoryWithMultipleVehicles(numberOfPages int) {
	// Mock HTTP Request for N amount of pages
	for i := 1; i < numberOfPages; i++ {
		url := fmt.Sprintf("auto-parts/our-inventory/page/%d/", i)
		gock.New("https://kennyupull.com").
			Get(url).
			Reply(200).
			BodyString(HTMLWIthMultipleVehiclesExample)
	}

	// Ensure to have a page that has no items
	urlEnd := fmt.Sprintf("auto-parts/our-inventory/page/%d/", numberOfPages)
	gock.New("https://kennyupull.com").
		Get(urlEnd).
		Reply(200).
		BodyString(HTMLWithoutVehicleExample)
}

func MockFirstPageBlank() {
	// Mock HTTP Request for the first page so it stops at the first page
	url1 := fmt.Sprintf("auto-parts/our-inventory/page/%d/", 1)

	gock.New("https://kennyupull.com").
		Get(url1).
		Reply(200).
		BodyString(HTMLWithoutVehicleExample)
}

func Test_GetInventory_SingleInventoryItem(t *testing.T) {
	defer gock.Off()
	MockPagingThroughInventory(2)
	inventorySearch := InventorySearch{
		Year:   "2010",
		Make:   "lexus",
		Model:  "is-250",
		Branch: "all-branches",
	}

	inventoryListings, err := GetInventory(inventorySearch)
	assert.Nil(t, err)
	assert.NotEmpty(t, inventoryListings)

	assert.Equal(t, 1, len(inventoryListings))
	assert.Equal(t, InventoryListing{"2010", "LEXUS", "IS 250", "2023-01-06", "26A", "Ottawa", "https://kennyupull.com/part/lexus-is250-2010-93/"}, inventoryListings[0])
}

func Test_GetInventory_MultipleInventoryItems(t *testing.T) {
	defer gock.Off()
	MockPagingThroughInventory(3)
	inventorySearch := InventorySearch{
		Year:   "2010",
		Make:   "lexus",
		Model:  "is-250",
		Branch: "all-branches",
	}

	inventoryListings, err := GetInventory(inventorySearch)
	assert.Nil(t, err)
	assert.NotEmpty(t, inventoryListings)

	assert.Equal(t, 2, len(inventoryListings))

	// For these tests, we are only checking the first item since the second item is the same
	inventoryListing := InventoryListing{"2010", "LEXUS", "IS 250", "2023-01-06", "26A", "Ottawa", "https://kennyupull.com/part/lexus-is250-2010-93/"}
	assert.Equal(t, inventoryListing, inventoryListings[0])
	assert.Equal(t, inventoryListing, inventoryListings[1])
}

func Test_GetInventory_NoBranch(t *testing.T) {
	defer gock.Off()
	MockPagingThroughInventory(2)
	inventorySearch := InventorySearch{
		Year:  "2010",
		Make:  "lexus",
		Model: "is-250",
	}

	inventoryListings, err := GetInventory(inventorySearch)

	assert.Nil(t, err)
	assert.NotEmpty(t, inventoryListings)

	assert.Equal(t, 1, len(inventoryListings))

	inventoryListing := InventoryListing{"2010", "LEXUS", "IS 250", "2023-01-06", "26A", "Ottawa", "https://kennyupull.com/part/lexus-is250-2010-93/"}
	assert.Equal(t, inventoryListing, inventoryListings[0])
}

func Test_GetInventory_NoMatchReturnEmpty(t *testing.T) {
	defer gock.Off()
	MockFirstPageBlank()
	inventorySearch := InventorySearch{
		Make: "jeffrey",
	}

	inventoryListings, err := GetInventory(inventorySearch)

	assert.Nil(t, err)
	assert.Empty(t, inventoryListings)
}

func Test_GetLatestListing_ListingFound(t *testing.T) {
	defer gock.Off()
	MockPagingThroughInventoryWithMultipleVehicles(2)
	inventorySearch := InventorySearch{
		Year:   "2010",
		Make:   "lexus",
		Model:  "is-250",
		Branch: "all-branches",
	}

	inventoryListing, err := GetLatestListing(inventorySearch)

	assert.Nil(t, err)
	assert.NotNil(t, inventoryListing)

	assert.Equal(t, &InventoryListing{"2010", "LEXUS", "IS 250", "2023-01-06", "26A", "Ottawa", "https://kennyupull.com/part/lexus-is250-2010-93/"}, inventoryListing)
}

func Test_GetLatestListing_NoListingFound(t *testing.T) {
	defer gock.Off()
	MockFirstPageBlank()
	inventorySearch := InventorySearch{
		Make: "jeffrey",
	}

	inventoryListing, err := GetLatestListing(inventorySearch)

	assert.Nil(t, err)
	assert.Nil(t, inventoryListing)
}
