package model

type Product struct {
	Name         string `json:"name" bson:"name"`
	Owner        string `json:"owner" bson:"owner"`
	FlgAvailable bool   `json:"flg_available" bson:"flg_available"`
}
