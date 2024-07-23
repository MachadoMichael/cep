package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type InfoCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	cep := os.Args[1]

	url := "https://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data InfoCEP
	json.Unmarshal(res, &data)
	fmt.Println(data)

}
