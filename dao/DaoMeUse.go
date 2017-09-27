package dao

import (
	"gopkg.in/mgo.v2"

	"fmt"
	"meuse/model"

	"gopkg.in/mgo.v2/bson"
)

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017/meuse")

	return session, err
}

func GetProduct() []model.Product {
	result := []model.Product{}

	session, err := getSession()

	if err != nil {
		fmt.Println("Erro ao recuperar sessao do mongo", err.Error())
		return nil
	}

	defer session.Close()

	c := session.DB("meuse").C("product")
	err = c.Find(bson.M{}).All(&result)

	return result
}

func GetAllocatedBox(boxCode int16) (model.ProductAllocated, error) {
	session, err := getSession()
	result := model.ProductAllocated{}

	if err != nil {
		fmt.Println("Erro ao recuperar sessao do mongo", err.Error())
		return result, err
	}

	defer session.Close()

	c := session.DB("meuse").C("product_allocated")

	err = c.Find(bson.M{"box_code": boxCode}).One(&result)

	return result, nil
}

func AllocateBox(productAllocated model.ProductAllocated) error {
	session, err := getSession()

	if err != nil {
		fmt.Println("Erro ao recuperar sessao do mongo", err.Error())
		return err
	}

	c := session.DB("meuse").C("product_allocated")

	err = c.Insert(productAllocated)

	if err != nil {
		fmt.Println("Erro ao inserir registro no MongoDB", err.Error())
		return err
	}

	return nil
}
