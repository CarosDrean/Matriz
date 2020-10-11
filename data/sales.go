package data

import "github.com/CarosDrean/Matriz/models"

func SalesClient(id string) []models.Sale {
	var a []models.Sale
	for _, d := range GetSales() {
		if d.IdClient == id {
			a = append(a, d)
		}
	}
	return a
}

func GetSales() []models.Sale {
	return []models.Sale{
		{
			ID:        "01",
			IdClient:  "01",
			IdProduct: "01",
			Count:     3,
		},
		{
			ID:        "02",
			IdClient:  "02",
			IdProduct: "01",
			Count:     3,
		},
		{
			ID:        "03",
			IdClient:  "01",
			IdProduct: "02",
			Count:     3,
		},
		{
			ID:        "04",
			IdClient:  "02",
			IdProduct: "02",
			Count:     3,
		},
		{
			ID:        "05",
			IdClient:  "01",
			IdProduct: "03",
			Count:     3,
		},
		{
			ID:        "06",
			IdClient:  "02",
			IdProduct: "03",
			Count:     3,
		},
		{
			ID:        "07",
			IdClient:  "02",
			IdProduct: "01",
			Count:     5,
		},
		{
			ID:        "08",
			IdClient:  "02",
			IdProduct: "01",
			Count:     5,
		},
		{
			ID:        "09",
			IdClient:  "01",
			IdProduct: "01",
			Count:     3,
		},
	}
}
