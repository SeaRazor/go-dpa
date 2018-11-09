package repository

import (
	"dpa/models"
)

type PolicyRepository struct {
	//Config *models.Configuration
}

func (repository *PolicyRepository) GetPolicy(policyId int) (models.Policy, error) {
	policy := models.Policy{}
	return policy, nil
}

func (repository *PolicyRepository) GetPolicies(params models.RequestParams) ([]models.Policy, error) {
	policies := make([]models.Policy, 10)
	return policies, nil
}

func (repository *PolicyRepository) CreatePolicy(policy *models.Policy) (models.Policy, error) {
	createdPolicy := models.Policy{}
	return createdPolicy, nil
}

func (repository *PolicyRepository) UpdatePolicy(id int, policy *models.Policy) (models.Policy, error) {
	updatedPolicy := models.Policy{}
	return updatedPolicy, nil
}

func (repository *PolicyRepository) DeletePolicy(id int) error {
	return nil
}
