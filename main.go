package main

import (
	"Assignment_2_CST_412/firestore"
	"html/template"
	"log"
	"net/http"
)

func main() {

	firestore.InitDB("firebase_key.json")
	defer firestore.CloseDB()

	http.HandleFunc("/", mainPage)

	log.Println("Running on port", "80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("html_files\\layout_1.html"))
	tmpl.Execute(w, nil)

	if r.Method == http.MethodPost {
		firestore.FindAsset(w, r.FormValue("asset"), tmpl)
	}
}

/*
func addNewAsset() {

	var asset Asset

	//The finished webhook with all data structured as it will be added to the DB.
	var NewAsset = map[string]interface{}{
		"asset_number":         asset.AssetNumber,
		"manufacturer":         asset.Manufacturer,
		"manufacturer_address": asset.ManufacturerAddress,
		"manufacturer_phone":   asset.ManufacturerPhone,
		"manufacturer_web":     asset.ManufacturerWeb,
		"model":                asset.Model,
		"date_purchased":       asset.DatePurchased,
		"purchase_price":       asset.PurchasePrice,
		"warranty_date":        asset.WarrantyDate,
		"retired_date":         asset.RetiredDate,
		"description":          asset.Description,
	}

	fmt.Println(NewAsset)
}*/
