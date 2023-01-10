package kennyupullgosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetInventory(t *testing.T) {
	GetInventory(InventorySearch{
		Year:   "2010",
		Make:   "honda",
		Model:  "civic",
		Branch: "all-branches",
	})

	assert.Equal(t, "3", 4, "Somehow 2 + 2 does not equal 4")
}
