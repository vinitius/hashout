package db

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"viniti.us/hashout/models/checkout"
)

type ProductsDataset struct {
	products []checkout.Product
}

func NewConnection() (d ProductsDataset, err error) {
	dbFile, err := os.Open("products.json")
	if err != nil {
		return
	}

	defer dbFile.Close()

	bytes, err := ioutil.ReadAll(dbFile)
	if err != nil {
		return
	}

	json.Unmarshal(bytes, &d.products)

	return
}
