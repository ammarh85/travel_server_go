package tc

import (
	"net/http"
	data "./data"
)

type Package struct {

}

func GetInventory(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Foo bar! - Something Strange happed. Dont know what :/ !"))

	data.GetAgencyList()
}

