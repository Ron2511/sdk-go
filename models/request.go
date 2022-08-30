package models

//Payment
type OrderInput struct {
	Amount          float32 `json:"amount"`
	Description     string  `json:"description"`
	UserName        string  `json:"userName"`
	CallbackFail    string  `json:"callback_fail"`
	CallbackSuccess string  `json:"callback_success"`
	Origin          string  `json:"origin"`
	NotificationUrl string  `json:"notification_url"`
}
