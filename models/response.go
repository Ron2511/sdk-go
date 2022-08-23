/*
client-api-checkout SDK GO
User API.
API version: 1.0.0
Contact: "MAIL CONTACT"
         "APICHECKOUTWEB"
         "GITHUB"
*/

package models

// APIResponse stores the API response models returned by the server.

type (
	//Order
	Order struct {
		OrderId   string  `json:"order_id"`
		Status    string  `json:"status"`
		RefNumber string  `json:"ref_number"`
		Amount    float32 `json:"amount"`
	}

	//Checkout
	OrderOutput struct {
		Id          string  `json:"id"`
		Type        string  `json:"type"`
		Uuid        string  `json:"uuid"`
		OrderNumber string  `json:"orderNumber"`
		Currency    string  `json:"currency"`
		Status      string  `json:"status"`
		RefNumber   string  `json:"refNumber"`
		Amount      float32 `json:"amount"`
		Links       Links   `json:"links"`
	}

	Links struct {
		CheckoutLink string `json:"checkoutLink"`
		Success      string `json:"success"`
		Failed       string `json:"failed"`
	}
)
