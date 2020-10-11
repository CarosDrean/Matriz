package helper

import (
	"github.com/CarosDrean/Matriz/data"
	"github.com/CarosDrean/Matriz/models"
)

func DataOrdered() []models.DataNeat {
	var dataN []models.DataNeat

	for _, e := range data.GetClients() {
		sales := data.SalesClient(e.ID)
		var q models.DataNeat
		q.Client = data.SearchClient(e.ID)
		q.Product = saleProduct(sales)
		dataN = append(dataN, q)
	}
	return dataN
}

func saleProduct(sales []models.Sale) []models.ProductCount {
	var pc []models.ProductCount
	for _, a := range sales {
		var item models.ProductCount
		item.Product = data.SearchProduct(a.IdProduct)
		item.Count = a.Count
		var b = true
		for i, q := range pc {
			if q.Product == item.Product {
				aux := pc[i]
				aux.Count = q.Count + item.Count
				pc[i] = aux
				b = false
			}
		}
		if b {
			pc = append(pc, item)
		}

	}
	return pc
}
