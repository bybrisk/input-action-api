// Package classification of Input-Action API
//
// Documentation for Input-Action API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/input-action-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Response for the order API
// swagger:response orderAPIResponse
type orderAPIResponseWrapper struct {
	// response for Order API
	// in: body
	Body data.OrderAPIResponse
}

// swagger:parameters order
type requestOrderParmsWrapper struct {
	// Data structure for the payload of the Order API
	// in: body
	// required: true
	Body data.OrderAPIRequest
}