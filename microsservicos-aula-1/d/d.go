package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coupon struct {
	Code string
	isWinner bool
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) CheckWinner(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code && item.isWinner{
			return "winner"
		}
	}
	return "valid"
}

type Result struct {
	Status string
}

var coupons Coupons

func main() {

	// Cria um cupom unico
	coupon := Coupon{
		Code: "def",
		isWinner: true,
	}

	// adiciona na lista de Coupons
	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")

	winner := coupons.CheckWinner(coupon)

	result := Result{Status: winner}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}

