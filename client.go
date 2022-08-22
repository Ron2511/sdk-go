package client

import "github.com/Ron2511/sdk-go/models"

type api struct {
	token string
}

type UalaBis interface {
	//Auth()
	GetOrder(orderId string) *models.Order
	Checkout() //Validar el nombre
	GetFailedNotifications()
}

//cli := api.New()
//ok, err := cli.Authenticate()
//cli.GetOrder()
