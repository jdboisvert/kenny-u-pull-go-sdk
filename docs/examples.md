# Functions and Usage

## GetInventory
Used to get a list of inventory listings based on the search criteria provided from the Kenny U-Pull Website/API.
This will include all the information about the vehicle and the branch it is located at. It will always be searching from latest to oldest listings.

### Example
```go
package main

import (
	"fmt"

	kennyupull "github.com/jdboisvert/kenny-u-pull-go-sdk"
)

func main() {
	inventorySearch := kennyupull.InventorySearch{
		Year:   "2010",
		Make:   "toyota",
		Model:  "corolla",
		Branch: "all-branches", // Optional by default will be set to "all-branches"
	}
	inventoryListings, err := kennyupull.GetInventory(inventorySearch)
	if err != nil {
		panic(err)
	}
	fmt.Println(inventoryListings) // Should display all the InventoryListing(s) matching your search criteria
}

```

## GetLatestListing
Used to get the latest inventory listing from the Kenny U-Pull Website/API.
This will include all the information about the vehicle and the branch it is located at.

### Example
```go
package main

import (
	"fmt"

	kennyupull "github.com/jdboisvert/kenny-u-pull-go-sdk"
)

func main() {
	inventorySearch := kennyupull.InventorySearch{
		Year:   "2010",
		Make:   "toyota",
		Model:  "corolla",
		Branch: "all-branches", // Optional by default will be set to "all-branches"
	}
	inventoryListings, err := kennyupull.GetLatestListing(inventorySearch)
	if err != nil {
		// Error occurred while getting the inventory listing
		panic(err)
	}
	if inventoryListings == nil {
		panic("No inventory listings found for the given inventory search criteria")
	}

	fmt.Println(inventoryListings) // Should display the latest InventoryListing matching your search criteria
}

```
