package repository

import (
	"dpa/models"
)

type PolicyRepository struct {
	//Config *models.Configuration
}

func (repository *PolicyRepository) GetPolicy(policyId int) (models.Policy, error){
    policy := models.Policy{}
    return policy, nil
}