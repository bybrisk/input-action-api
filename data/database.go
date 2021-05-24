package data

import (
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func IsUserRegistered(docID string) (bool,error) {
	var isRegistered bool

	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}

	type getData struct {
		Phonenumber string `json:"phonenumber"`
		UserName string `json:"username"`
	}

	var document getData

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetInitialConversationMongo ERROR:")
		log.Error(err)
	}

	if err!=nil {
		isRegistered = false
	} else {
		if document.UserName == "" {
			isRegistered = false
		} else {
			isRegistered = true
		}
	}
	return isRegistered,err
}

func GetActionResponseMongo (d *OrderAPIRequest) (string,error) {
	collectionName := shashankMongo.DatabaseName.Collection("bot-schema")
	filter := bson.M{"businessid": d.BusinessID}

	var err error

	type OrderResultStruct struct{
		Message string `json:"message"`
	}
	type HandlerResponseWrapper struct{
		Order OrderResultStruct `json:"order"`
	}
	type ResponseResponseWrapper struct{
		Response HandlerResponseWrapper `json:"response"`
	}
	
	var document ResponseResponseWrapper

	err = collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("GetActionResponseMongo ERROR:")
		log.Error(err)
	}

	return document.Response.Order.Message,err
}

func SaveDeliveryIDToMongo(d *OrderAPIRequest, deliveryID string) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("input-user")
	id, _ := primitive.ObjectIDFromHex(d.UserID)
	filter := bson.M{"_id": id}

	type InputReferenceObject struct {
		DeliveryID string `json:"deliveryID"`
		InputStatus string `json:"inputStatus"`
		MapsStatus string `json:"mapsStatus"`
		BusinessID string `json:"businessID"`
	}

	val := InputReferenceObject{
		DeliveryID: deliveryID,
		InputStatus: "order",
		MapsStatus: "pending",
		BusinessID: d.BusinessID,
	}

	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.M{"$push":bson.M{"orderRef": val}})
	if err != nil {
		log.Error("SaveDeliveryIDToMongo ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}