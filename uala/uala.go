package uala

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Ron2511/sdk-go/constants"
	"github.com/Ron2511/sdk-go/models"
)

type api struct {
	AuthURL string
	BaseURL string
	Token   string
}

type Client interface {
	//Auth()
	GetOrder(orderId string) models.Order
	/// Validar el nombre
	Checkout()
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
	var order models.Order
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
	if err := json.Unmarshal(body, &order); err != nil {
		panic(err.Error())
	}
	return order
}

func (a *api) Checkout() {
	panic("Not implemented!")
}

func (a *api) GetFailedNotifications() {
	panic("Not implemented!")
}

//cli := api.New()
//ok, err := cli.Authenticate()
//cli.GetOrder()
