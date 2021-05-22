package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/input-action-api/data"
)

func TestOrderAPICrudOps(t *testing.T) {

	payload := &data.OrderAPIRequest{
		ActionHandler:"order",
		UserID: "6083dafb171b889e90c5c7aa",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
		Phone:"9340212623",
		CustomerAddress: "MANIT Bhopal, MP",
		CustomerName: "Shashank Prakash",
		ItemWeight: 1,
		Latitude: 23.4043444,
		Longitude: 77.3493045,
	}

	res:= data.OrderAPICrudOps(payload) 

	fmt.Println(res)
}

	//Note string `json:"note"`

	//Amount float64 `json:"amount" validate:"required"`

	//PaymentStatus float64 `json:"paymentStatus"`