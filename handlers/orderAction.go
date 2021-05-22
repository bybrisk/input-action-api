
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-action-api/data"
)

// swagger:route POST /action/order action order
// Call the order API from action module.
//
// responses:
//	200: orderAPIResponse
//  422: errorValidation
//  501: errorResponse

func (p *Input_Action) Make_Action_Order (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-action-api Module")
	request := &data.OrderAPIRequest{}

	err:=request.FromJSONToOrderAPIRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = request.ValidateOrderAPIRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-action-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//Take action for Order API
	response := data.OrderAPICrudOps(request)

	//writing to the io.Writer
	err = response.OrderAPIResponseToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}