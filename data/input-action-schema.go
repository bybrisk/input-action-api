package data

import (
	"github.com/go-playground/validator/v10"
)

//response for calling Order API
type OrderAPIResponse struct{
	// Message the business wants the user to see.
	//
	Message string `json: "message"`

	//Response message
	Response string `json:"response"`

	//Status code
	//
	Status int64 `json:"status"`
}

//post request for calling the Order API
type OrderAPIRequest struct{
	// UserID of the user 
	//
	// required: true
	// max length: 1000
	UserID string `json: "userID" validate:"required"`

	// BusinessID of the business under context
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessid" validate:"required"`

	// User Address 
	//
	// required: true
	CustomerAddress string `json:"customerAddress" validate:"required"`

	// Name of the user
	//
	// required: true
	CustomerName string `json:"customerName" validate:"required"`

	// Weight of the order (if not provided then use 1 as default)
	//
	// required: true
	ItemWeight float64 `json:"itemWeight" validate:"required"`

	// Latitude of the user
	//
	// required: true
	Latitude float64 `json:"latitude" validate:"required"`

	// Longitude of the user
	//
	// required: true
	Longitude float64 `json:"longitude" validate:"required"`

	// Phone number of the user
	//
	// required: true
	Phone string `json:"phone" validate:"required"`

	// Note from the user
	//
	// required: false
	Note string `json:"note"`

	// Amount of the transaction (if any)
	//
	// required: false
	Amount float64 `json:"amount" validate:"required"`

	// Status of payment (if amount is not provided then set this to true)
	//
	// required: false
	PaymentStatus float64 `json:"paymentStatus"`

}

type ResponseFromMaps struct {
	DeliveryID string `json:"deliveryID"`
	Message string `json:"message"` 
}

func (d *OrderAPIRequest) ValidateOrderAPIRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}