package tc

import (
	"net/http"
)

type Package struct {

}

func GetInventory(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

