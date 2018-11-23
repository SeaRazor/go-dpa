package controllers

import (
	"dpa/contracts"
	"dpa/models"
	"dpa/resources"
	"encoding/json"
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
	id, err := strconv.Atoi(policyId)
	policy, err := handler.PolicyResource.GetPolicyById(int32(id))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(policy)
}

func (handler *PolicyController) GetPolicies(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	sorting := vars["sorting"]
	filterOptions := []string{filter}
	sortingOptions := []string{sorting}
	requestParams := models.CreateRequestParams(filterOptions, sortingOptions)
	policies, err := handler.PolicyResource.GetPolicies(requestParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(policies.Documents)
}

func (handler *PolicyController) PostPolicy(w http.ResponseWriter, r *http.Request) {
	requestBody := r.Body
	if requestBody == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var policyDocument *contracts.PolicyDocument
	decoder := json.NewDecoder(requestBody)
	err := decoder.Decode(&policyDocument)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newPolicy, err := handler.PolicyResource.CreatePolicy(*policyDocument)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPolicy)
}

func (handler *PolicyController) PutPolicy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	policyId := vars["policyId"]
	var policyDocument contracts.PolicyDocument
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&policyDocument)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(policyId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedPolicy, err := handler.PolicyResource.UpdatePolicy(int32(id), policyDocument)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPolicy)
}

func (handler *PolicyController) DeletePolicy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	policyId := vars["policyid"]
	id, err := strconv.Atoi(policyId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PolicyResource.DeletePolicy(int32(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}
