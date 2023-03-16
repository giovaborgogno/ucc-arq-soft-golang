package main

import (
	"encoding/json"
	"fmt"
	"go-api/items"
)

func main() {

	item := items.GetItem()
	itemJSON, _ := json.Marshal(item)

	fmt.Println(string(itemJSON))
}
