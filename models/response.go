/*
uala-api-checkout SDK GO
User API.
API version: 1.0.0
Contact: "MAIL CONTACT"
         "APICHECKOUTWEB"
         "GITHUB"
*/

package models

// APIResponse stores the API response models returned by the server.

//Order
type Order struct {
	orderId   string
	status    string
	refNumber string
	amount    float32
}

// APIResponse stores the API response returned by the server.
type Checkout struct {
}
