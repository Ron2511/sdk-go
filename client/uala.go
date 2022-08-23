package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Ron2511/sdk-go/constants"
	"github.com/Ron2511/sdk-go/models"
	"io/ioutil"
	"net/http"
)

type api struct {
	AuthURL string
	BaseURL string
	Token   string
}

type Client interface {
	//Auth()
	GetOrder(orderId string) models.Order
	///Validar el nombre
	Checkout(body models.OrderInput) models.OrderOutput
	GetFailedNotifications()
}

func New(token string) Client {
	return &api{
		AuthURL: constants.AUTH_URL,
		BaseURL: constants.BASE_URL,
		Token:   token,
	}
}

func (a *api) GetOrder(orderId string) models.Order {
	endpoint := fmt.Sprintf("%s/1/order/%s", a.BaseURL, orderId)
	bearerToken := fmt.Sprintf("Bearer %s", a.Token)
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		panic(err.Error())
	}
	request.Header.Add("Authorization", bearerToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	var order models.Order
	if err := json.Unmarshal(body, &order); err != nil {
		panic(err.Error())
	}
	return order
}

func (a *api) Checkout(body models.OrderInput) models.OrderOutput {
	endpoint := fmt.Sprintf("%s/1/checkout", a.BaseURL)
	bearerToken := fmt.Sprintf("Bearer %s", a.Token)
	bodyPost, err := json.Marshal(body)
	if err != nil {
		panic(err.Error())
	}
	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(bodyPost))
	request.Header.Add("Authorization", bearerToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	bodyReturn, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	var orderOutput models.OrderOutput
	if err := json.Unmarshal(bodyReturn, &orderOutput); err != nil {
		panic(err.Error())
	}
	return orderOutput
}

func (a *api) GetFailedNotifications() {
	endpoint := fmt.Sprintf("%s/notifications/", a.BaseURL)
	bearerToken := fmt.Sprintf("Bearer %s", a.Token)
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		panic(err.Error())
	}
	request.Header.Add("Authorization", bearerToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	//No se que carajo es el objego FailedNotifications
}
