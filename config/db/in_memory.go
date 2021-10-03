package db

import (
	"encoding/json"
	"io/ioutil"

	"viniti.us/hashout/config/log"

	"viniti.us/hashout/models/checkout"
)

type Database interface {
	FindProductByID(id int32) (checkout.Product, bool)
	FindLastProductByIsGift(isGift bool) (checkout.Product, bool)
}

type ProductDataset struct {
	Database Database
	products []checkout.Product
	byID     map[int32]checkout.Product
	byIsGift map[bool]checkout.Product
}

func (d ProductDataset) FindProductByID(id int32) (p checkout.Product, found bool) {
	p, found = d.byID[id]
	return
}

func (d ProductDataset) FindLastProductByIsGift(isGift bool) (p checkout.Product, found bool) {
	p, found = d.byIsGift[isGift]
	return
}

func NewConnection() (d ProductDataset) {
	dbFile, err := ioutil.ReadFile("../config/db/products.json")
	if err != nil {
		log.Logger.Fatalw("Could not read DB File", "error", err.Error())
		return
	}

	json.Unmarshal(dbFile, &d.products)

	d.byID = make(map[int32]checkout.Product)
	d.byIsGift = make(map[bool]checkout.Product)

	for _, p := range d.products { // small known static collection
		d.byID[p.ID] = p
		d.byIsGift[p.IsGift] = p
	}

	log.Logger.Info("DB is up and running: dataset: ", d.products)

	return
}
