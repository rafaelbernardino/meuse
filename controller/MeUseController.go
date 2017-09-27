package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"meuse/dao"
	"meuse/model"
	"meuse/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Teste(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
}

func GetListProduct(res http.ResponseWriter, req *http.Request) {
	response := utils.ProductResponse{}
	res.Header().Add("Content-Type", "application/json;charset=utf-8")
	listProduct := dao.GetProduct()

	response.Products = listProduct
	response.Description = "Consulta realizada com sucesso"

	json.NewEncoder(res).Encode(response)
}

func OpenBox(res http.ResponseWriter, req *http.Request) {
	response := utils.BoxResponse{}
	res.Header().Add("Content-Type", "application/json;charset=utf-8")

	//strBoxCode := req.URL.Query().Get("box_code")

	strPinCode := req.URL.Query().Get("pin_code")

	if strings.Trim(strPinCode, "") == "" {
		response.Description = "Código da caixa inválido"
		response.ResponseCode = 99
		json.NewEncoder(res).Encode(response)
		return
	}

	pinCode, err := strconv.ParseInt(strPinCode, 10, 16)

	if err != nil {
		response.Description = "Código da caixa deve ser composto apenas por números"
		response.ResponseCode = 99
		json.NewEncoder(res).Encode(response)
		return
	}

	prodAllocated, err := dao.GetAllocatedBox(171)

	if err != nil {
		response.ResponseCode = 99
		response.Description = "Erro na busca de dados"
		json.NewEncoder(res).Encode(response)
		return
	}

	if prodAllocated.PinCode == int16(pinCode) {
		response.ResponseCode = 0
		response.Description = "Pin válido"
	} else {
		response.ResponseCode = 2
		response.Description = "Pin incorreto"
	}

	json.NewEncoder(res).Encode(response)
}

func AllocateBox(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Acessou o AllocateBox")
	response := utils.AllocateResponse{}
	res.Header().Add("Content-Type", "application/json;charset=utf-8")

	decode := json.NewDecoder(req.Body)
	allocate := utils.AllocateRequest{}
	err := decode.Decode(&allocate)

	if err != nil {
		response.ResponseCode = 2
		response.Description = "Erro no parse dos dados do body"
		fmt.Println("Erro no parse dos dados do body", err.Error())
		return
	}

	prodAllocate := model.ProductAllocated{}
	prodAllocate.BoxCode = 171
	prodAllocate.PinCode = int16(randomPin())
	prodAllocate.Product = allocate.Product

	err = dao.AllocateBox(prodAllocate)

	if err != nil {
		response.ResponseCode = 99
		response.Description = "Erro na inserção do registro"
	} else {
		response.PinCode = prodAllocate.PinCode
		response.Description = "Registro inserido com sucesso"
	}

	json.NewEncoder(res).Encode(response)
}

func randomPin() int {
	max := 9999
	min := 1000
	rand.Seed(time.Now().Unix())
	pin := rand.Intn(int(max)-min) + min
	return pin
}
