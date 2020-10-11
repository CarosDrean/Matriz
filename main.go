package main

import "log"

func main() {
	var datac []fin

	for _, e := range allClients{
		sales := salesClient(e.ID)
		var q fin
		q.Client = searchClient(e.ID)
		q.Product = saleProduct(sales)
		datac = append(datac, q)
	}
	log.Println(datac)
}

func saleProduct(sales []sale) []productCount{
	var pc []productCount
	for _, a := range sales{
		var item productCount
		item.Product = searchProduct(a.IdProduct)
		item.Count = a.Count
		var b = true
		for i, q := range pc{
			if q.Product == item.Product{
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

func salesClient(id string) []sale{
	var a []sale
	for _, d := range allSales{
		if d.IdClient == id{
			a = append(a, d)
		}
	}
	return a
}

func searchClient(id string) string{
	var a client
	for _, d := range allClients{
		if d.ID == id{
			a = d
		}
	}
	return a.Name
}

func searchProduct(id string) string{
	var a product
	for _, d := range allProducts{
		if d.ID == id{
			a = d
		}
	}
	return a.Name
}

type fin struct {
	Client string
	Product []productCount
}

type productCount struct {
	Product string
	Count int32
}

type client struct {
	ID string
	Name string
}

type product struct {
	ID string
	Name string
}

type sale struct {
	ID string
	IdProduct string
	IdClient string
	Count int32
}

type clients []client
type products []product
type sales []sale

var allClients = clients{
	{
		ID: "01",
		Name: "Oscar",
	},
	{
		ID: "02",
		Name: "Yeison",
	},
}

var allProducts = products{
	{
		ID: "01",
		Name: "Papa",
	},
	{
		ID: "02",
		Name: "Cebolla",
	},
	{
		ID: "03",
		Name: "Naranja",
	},
	{
		ID: "04",
		Name: "Mandarina",
	},
}

var allSales = sales{
	{
		ID: "01",
		IdClient: "01",
		IdProduct: "01",
		Count: 3,
	},
	{
		ID: "02",
		IdClient: "02",
		IdProduct: "01",
		Count: 3,
	},
	{
		ID: "03",
		IdClient: "01",
		IdProduct: "02",
		Count: 3,
	},
	{
		ID: "04",
		IdClient: "02",
		IdProduct: "02",
		Count: 3,
	},
	{
		ID: "05",
		IdClient: "01",
		IdProduct: "03",
		Count: 3,
	},
	{
		ID: "06",
		IdClient: "02",
		IdProduct: "03",
		Count: 3,
	},
	{
		ID: "07",
		IdClient: "02",
		IdProduct: "01",
		Count: 5,
	},
	{
		ID: "08",
		IdClient: "02",
		IdProduct: "01",
		Count: 5,
	},
	{
		ID: "09",
		IdClient: "01",
		IdProduct: "01",
		Count: 3,
	},
}