package data

import "github.com/CarosDrean/Matriz/models"

func SearchProduct(id string) string {
	var a models.Product
	for _, d := range GetProducts() {
		if d.ID == id {
			a = d
		}
	}
	return a.Name
}

func GetProducts() []models.Product {
	return []models.Product{
		{
			ID:   "01",
			Name: "Papa",
		},
		{
			ID:   "02",
			Name: "Cebolla",
		},
		{
			ID:   "03",
			Name: "Naranja",
		},
		{
			ID:   "04",
			Name: "Mandarina",
		},
	}
}
