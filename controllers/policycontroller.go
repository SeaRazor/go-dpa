package controllers

import (
	"dpa/resources"
	"encoding/json"
	"dpa/helpers"
	"github.com/gorilla/mux"
  "net/http"
  "strconv"
)

type PolicyController struct {
    PolicyResource resources.PolicyResource
}

func (handler *PolicyController) GetPolicy(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  policyId := vars["policyId"]
  err := helpers.ValidateId(policyId)
  if err != nil  {
     panic(err)
  }
  id ,err := strconv.Atoi(policyId)
  policy,err := handler.PolicyResource.GetDpaPolicyById(id)
  if err != nil{
    panic(err)
  }
  json.NewEncoder(w).Encode(policy)
}

func (handler *PolicyController) GetPolicies(w http.ResponseWriter, r *http.Request) {

}

func (handler *PolicyController) PostPolicy(w http.ResponseWriter, r *http.Request) {

}

func (handler *PolicyController) PutPolicy(w http.ResponseWriter, r *http.Request) {

}

func (handler *PolicyController) DeletePolicy(w http.ResponseWriter, r *http.Request) {

}

