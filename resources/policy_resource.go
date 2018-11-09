package resources

import (
	"dpa/models"
	"dpa/repository"
)

type PolicyResource struct {
	PolicyRepository repository.PolicyRepository
}

func (res *PolicyResource) GetPolicyById(id int) (models.Policy, error) {
	policy, err := res.PolicyRepository.GetPolicy(id)
	if err != nil {
		panic(err)
	}
	return policy, nil

}

func (res *PolicyResource) GetPolicies(params models.RequestParams) ([]models.Policy, error) {
	policies, err := res.PolicyRepository.GetPolicies(params)
	if err != nil {
		panic(err)
	}
	return policies, nil

}

func (res *PolicyResource) CreatePolicy(policy *models.Policy) (models.Policy, error) {
	newPolicy, err := res.PolicyRepository.CreatePolicy(policy)
	if err != nil {
		panic(err)
	}
	return newPolicy, nil
}

func (res *PolicyResource) UpdatePolicy(id int, policy *models.Policy) (models.Policy, error) {
	updatedPolicy, err := res.PolicyRepository.UpdatePolicy(id, policy)
	if err != nil {
		panic(err)
	}
	return updatedPolicy, nil
}

func (res *PolicyResource) DeletePolicy(id int) error {
	err := res.PolicyRepository.DeletePolicy(id)
	if err != nil {
		panic(err)
	}
	return nil
}
