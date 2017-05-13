package tc

import (
	data "./data"
	"net/http"
	"log"
	"encoding/json"
)

type Agency struct {
	AgencyId	string	`json:"agencyId"`
	AgencyName	string	`json:"agencyName"`
	AddressLine1	string	`json:"addressLine1"`
	AddressLine2	string	`json:"addressLine2"`
	City	string	`json:"city"`
	State	string	`json:"state"`
	Zip 	string	`json:"zip"`
	Country	string	`json:"country"`
	Email	string	`json:"email"`
	Phone	string	`json:"phone"`
	Website	string	`json:"website"`
}

func CreateAgency(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

// fetches the full list of available agencies which are not marked as deleted
// returns json string of Agency objects to caller
func GetAgencies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	itemList := data.GetItemsList("Agency")
	var agencies []Agency
	err := data.UnmarshalListOfMaps(itemList, &agencies)
	log.Println(err)
	json.NewEncoder(w).Encode(agencies)
}

func GetAgencyById(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

