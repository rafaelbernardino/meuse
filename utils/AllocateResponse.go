package utils

type AllocateResponse struct {
	ResponseCode int8   `json:"response_code"`
	Description  string `json:"description"`
	PinCode      int16  `json:"pin_code"`
}
