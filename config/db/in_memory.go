package db

import (
	"encoding/json"
	"io/ioutil"

	"viniti.us/hashout/config/log"

	"viniti.us/hashout/models/checkout"
)

type ProductsDataset struct {
	Products []checkout.Product
	ByID     map[int32]checkout.Product
	ByIsGift map[bool]checkout.Product
}

func NewConnection() (d ProductsDataset) {
	dbFile, err := ioutil.ReadFile("../config/db/products.json")
	if err != nil {
		log.Logger.Fatalw("Could not read DB File", "error", err.Error())
		return
	}

	json.Unmarshal(dbFile, &d.Products)

	d.ByID = make(map[int32]checkout.Product)
	d.ByIsGift = make(map[bool]checkout.Product)

	for _, p := range d.Products { // small known static collection
		d.ByID[p.ID] = p
		d.ByIsGift[p.IsGift] = p
	}

	log.Logger.Info("DB is up and running: dataset: ", d.Products)

	return
}
