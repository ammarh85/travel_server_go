package tc

import (
	"net/http"
)

type Agent struct {

}

func CreateAgent(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

