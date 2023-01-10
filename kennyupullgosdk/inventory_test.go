package kennyupullgosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetInventory_NotEmpty(t *testing.T) {
	inventorySearch := InventorySearch{
		Year:   "2010",
		Make:   "lexus",
		Model:  "is-250",
		Branch: "all-branches",
	}

	inventoryListings, err := GetInventory(inventorySearch)
	assert.Nil(t, err)
	assert.NotEmpty(t, inventoryListings)
}
