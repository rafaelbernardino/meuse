package utils

import (
	"meuse/model"
)

type ProductResponse struct {
	ResponseCode int8            `json:"response_code"`
	Description  string          `json:"description"`
	Products     []model.Product `json:"produto"`
}
