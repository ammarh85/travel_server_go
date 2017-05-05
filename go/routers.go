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
	fmt.Fprintf(w, "Hello World From Router!!!")
}

func ApiInfo(w http.ResponseWriter, r *http.Request) {
	const indexPage = "api/swagger.yaml"
	http.ServeFile(w, r, indexPage)
}

var routes = Routes{
	Route{
		"ApiInfo",
		"GET",
		"/api/",
		ApiInfo,
	},

	Route{
		"Index",
		"GET",
		"/v2/",
		Index,
	},

	Route{
		"CreateAgent",
		"POST",
		"/v2/agent",
		CreateAgent,
	},

	Route{
		"GetInventory",
		"GET",
		"/v2/package/inventory",
		GetInventory,
	},

	Route{
		"AddUser",
		"POST",
		"/v2/user",
		AddUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		"POST",
		"/v2/agent/createWithArray",
		CreateUsersWithArrayInput,
	},

	Route{
		"CreateUsersWithListInput",
		"POST",
		"/v2/agent/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"LoginUser",
		"GET",
		"/v2/agent/login",
		LoginUser,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/v2/user",
		UpdateUser,
	},

}