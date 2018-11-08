package resources

import (
	"dpa/repository"
	"dpa/models"
)

type PolicyResource struct {
   PolicyRepository repository.PolicyRepository
}

func (res *PolicyResource) GetDpaPolicyById(id int) (models.Policy, error)  {
	policy, err := res.PolicyRepository.GetPolicy(id)
	if err != nil{
		panic(err)
	}
	return policy, nil

}