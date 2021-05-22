package data

import (
	"encoding/json"
	"net/http"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"bytes"
)

func OrderAPICrudOps(d *OrderAPIRequest) *OrderAPIResponse{
	
	var response OrderAPIResponse
	var emptyString string

	//isUserRegistered
	isRegistered,registeredErr := IsUserRegistered(d.UserID)

	if (!isRegistered || registeredErr!=nil) {
		response= OrderAPIResponse{
			Message:emptyString,
			Response:"Error! User not registered!",
		}
	} else{

		//Send the data to Maps API and save it.
		
		//infiltrate the data
		//"amount"
		//"paymentStatus"

		//prepare the payload
        postBody, _ := json.Marshal(map[string]interface{}{
			"BybID":  d.BusinessID,
			"CustomerAddress": d.CustomerAddress,
			"CustomerName" : d.CustomerName,
			"itemWeight" : d.ItemWeight,
			"latitude" : d.Latitude,
			"longitude" : d.Longitude,
			"phone" : d.Phone,
			"pincode" : d.Note,
		 })
		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post("https://developers.bybrisk.com/delivery/create/al", "application/json", responseBody)
		if err != nil {
			log.Error("OrderAPICrudOps ERROR:")
			log.Error(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
		}

		var betaResponseCrossOrigin ResponseFromMaps
		err = json.Unmarshal(body, &betaResponseCrossOrigin)

		if betaResponseCrossOrigin.Message == "Delivery added to ES Queue" {
			// Save deliveryID against UserID
			_=SaveDeliveryIDToMongo(d,betaResponseCrossOrigin.DeliveryID)

			//Get the message from mongo
			message,err := GetActionResponseMongo(d)
			if err!=nil{
				response= OrderAPIResponse{
					Message:emptyString,
					Response:"Error! Database error!",
				}
			} else {
				response= OrderAPIResponse{
					Message:message,
					Response:"success",
				}
			}
		} else {
			response= OrderAPIResponse{
				Message:emptyString,
				Response:"Error! Internal API error!",
			}
		}
	}
	return &response
}