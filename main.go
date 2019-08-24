package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Number struct {
	Decimal int    `json:"decimal"`
	Roman   string `json:"roman"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func decimalToRoman(num int) string {
	roman := ""
	Decimals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1,}
	Symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",}

	i := 0
	for num > 0 {
		k := num / Decimals[i]
		for j := 0; j < k; j++ {
			roman += Symbols[i]
			num -= Decimals[i]
		}
		i++
	}
	return roman
}

func GetRoman(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	decimal, err := strconv.Atoi(mux.Vars(r)["decimal"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	roman := decimalToRoman(decimal)
	response := Number{Decimal: decimal, Roman: roman}
	_ = json.NewEncoder(w).Encode(response)
}

func main() {
	log.Println("Start application")
	router := mux.NewRouter()
	router.HandleFunc("/decimal/{decimal}", GetRoman).Methods("GET")
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
