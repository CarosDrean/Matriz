package data

import "github.com/CarosDrean/Matriz/models"

func SearchClient(id string) string {
	var a models.Client
	for _, d := range GetClients() {
		if d.ID == id {
			a = d
		}
	}
	return a.Name
}

func GetClients() []models.Client {
	return []models.Client{
		{
			ID:   "01",
			Name: "Oscar",
		},
		{
			ID:   "02",
			Name: "Yeison",
		},
	}
}
