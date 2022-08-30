package client

import (
	"bytes"
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
	Auth(credentials string)
	GetOrder(orderId string) models.Order
	///Validar el nombre
	Checkout(body models.OrderInput) models.OrderOutput
	GetFailedNotifications() models.FailedNotifications
}

func New(token *string) Client {
	tk := func(token *string) string {
		if token != nil {
			return *token
		}
		return ""
	}(token)
	return &api{
		AuthURL: constants.AUTH_URL,
		BaseURL: constants.BASE_URL,
		Token:   tk,
	}
}

func (a *api) Auth(credentials string) {
	type authToken struct {
		AccessToken string `json:"access_token"`
	}

	var (
		endpoint   string = fmt.Sprintf("%s/1/auth/token", a.AuthURL)
		token      authToken
		payload    []byte        = []byte(credentials)
		bodyReader *bytes.Reader = bytes.NewReader(payload)
	)

	request, err := http.NewRequest(http.MethodPost, endpoint, bodyReader)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(body, &token); err != nil {
		panic(err)
	}

	a.Token = token.AccessToken
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
	if err != nil {
		panic(err)
	}
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

func (a *api) GetFailedNotifications() models.FailedNotifications {
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	var failedNotifications models.FailedNotifications
	if err := json.Unmarshal(body, &failedNotifications); err != nil {
		panic(err.Error())
	}
	return failedNotifications
}
