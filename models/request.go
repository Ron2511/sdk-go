package models

//Payment
type OrderInput struct {
	amount          float32 `json:"amount"`
	description     string  `json:"description"`
	userName        string  `json:"userName"`
	callbackFail    string  `json:"callback_fail"`
	callbackSuccess string  `json:"callback_success"`
	origin          string  `json:"origin"`
	notificationUrl string  `json:"notification_url"`
}
