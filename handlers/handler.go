package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/IRFAN374/goRestApi/model"
	"github.com/gorilla/mux"
)

var groceries = []model.Grocery{
	{Name: "Milk", Quantity: 2},
	{Name: "Apple", Quantity: 10},
	{Name: "Biryani", Quantity: 1},
}

func AllGrocery(w http.ResponseWriter, r *http.Request) {

	 w.Header().Add("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(groceries); err != nil {
        fmt.Println(err)
        http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
	fmt.Println("EndPoint Hit : Allgrocery")
	// json.NewEncoder(w).Encode(groceries)

}

func GetGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for _, grocery := range groceries {
		if grocery.Name == name {
			json.NewEncoder(w).Encode(grocery)
		}
	}
}

func AddGrocery(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var grocery model.Grocery

	_ = json.Unmarshal(reqBody, &grocery)

	groceries = append(groceries, grocery)

	json.NewEncoder(w).Encode(groceries)
}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(groceries)
}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)

			var updatedGrocery model.Grocery 
			json.NewDecoder(r.Body).Decode(&updatedGrocery)

			groceries = append(groceries, updatedGrocery)
			fmt.Println("Updated Endpoint Hit")
			json.NewEncoder(w).Encode(updatedGrocery)

			return
		}
	}
}
