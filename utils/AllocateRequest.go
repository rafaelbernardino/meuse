package utils

import (
	"meuse/model"
)

type AllocateRequest struct {
	Product model.Product `json:"product"  bson:"product"`
	PinCode int16         `json:"pin_code" bson:"pin_code"`
	BoxCode int16         `json:"box_code" bson:"box_code"`
}
