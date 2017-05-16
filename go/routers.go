package tc

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Travel Concierge!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/tc/",
		Index,
	},

	Route{
		"CreateAgency",
		"POST",
		"/tc/agency",
		CreateAgency,
	},

	Route{
		"GetAgencies",
		"GET",
		"/tc/agency/getAgency",
		GetAgencies,
	},

	Route{
		"GetAgencyById",
		"GET",
		"/tc/agency/getAgencyById",
		GetAgencyById,
	},

	Route{
		"CreateAgent",
		"POST",
		"/tc/agent",
		CreateAgent,
	},

	Route{
		"CreateAgentsWithArrayInput",
		"POST",
		"/tc/agent/createWithArray",
		CreateAgentsWithArrayInput,
	},

	Route{
		"GetAgents",
		"GET",
		"/tc/agent/getAgents",
		GetAgents,
	},

	Route{
		"GetagentById",
		"GET",
		"/tc/agent/getAgentById",
		GetagentById,
	},

	Route{
		"LoginUser",
		"GET",
		"/tc/agent/login",
		LoginUser,
	},

	Route{
		"GetInventory",
		"GET",
		"/tc/package/inventory",
		GetInventory,
	},

	Route{
		"AddUser",
		"POST",
		"/tc/user",
		AddUser,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/tc/user",
		UpdateUser,
	},

}