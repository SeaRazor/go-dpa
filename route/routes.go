package route

import (
    "dpa/controllers"
    "dpa/resources"
    "dpa/repository"

    "net/http"

    "github.com/gorilla/mux"


    
)
var controller = &controllers.PolicyController {PolicyResource : resources.PolicyResource{PolicyRepository : repository.PolicyRepository{}}}

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route {
        "GetPolicy",
        "GET",
        "/v2/dpa_policies/{policyId}",
        controller.GetPolicy,
    },
    Route {
        "GetPolicies",
        "GET",
        "/v2/dpa_policies",
        controller.GetPolicies,
    },
    Route {
        "CreatePolicy",
        "POST",
        "/v2/dpa_policies",
        controller.PostPolicy,
    },
    Route {
        "UpdatePoliciy",
        "PUT",
        "/v2/dpa_policies/{policyId}",
        controller.PutPolicy,
    },
    Route {
        "DeletePolicy",
        "DELETE",
        "/v2/dpa_policies/{policyid}",
        controller.DeletePolicy,
    },

}

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    for _, route := range routes{
        var handler http.Handler
        handler = route.HandlerFunc
        
        router.
        Methods(route.Method).
        Path(route.Pattern).
        Name(route.Name).
        Handler(handler)
    }
    return router
}


