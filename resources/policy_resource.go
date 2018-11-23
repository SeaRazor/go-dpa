package resources

import (
	"dpa/contracts"
	"dpa/converters"
	"dpa/models"
	"dpa/repository"
)

type PolicyResource struct {
	PolicyRepository repository.PolicyRepository
}

func (res *PolicyResource) GetPolicyById(id int32) (contracts.PolicyDocument, error) {
	policy, err := res.PolicyRepository.GetPolicy(id)
	if err != nil {
		panic(err)
	}
	policyDocument,err := converters.ConvertDomainToContract(policy)
	if err != nil{
		return policyDocument, err
	}
	return policyDocument, nil

}

func (res *PolicyResource) GetPolicies(params models.RequestParams) (contracts.PolicyDocumentCollection, error) {
	policies, err := res.PolicyRepository.GetPolicies(params)
	if err != nil {
		panic(err)
	}
	policyDocumentCollection,err := converters.ConvertDomainToContractCollection(policies)
	if err != nil{
		return policyDocumentCollection,err
	}
	return policyDocumentCollection, nil

}

func (res *PolicyResource) CreatePolicy(policyDocument contracts.PolicyDocument) (contracts.PolicyDocument, error) {

	newPolicy,err := converters.ConvertContractToDomain(policyDocument)
	if err != nil {
		panic(err)
	}
	createdPolicy, err := res.PolicyRepository.CreatePolicy(newPolicy)
	if err != nil {
		panic(err)
	}
	newPolicyDocument,err := converters.ConvertDomainToContract(createdPolicy)
	if err != nil {
		panic(err)
	}
	return newPolicyDocument, nil
}

func (res *PolicyResource) UpdatePolicy(id int32, policyDocument contracts.PolicyDocument) (contracts.PolicyDocument, error) {
	policyToUpdate,err := converters.ConvertContractToDomain(policyDocument)
	if err != nil{
		panic(err)
	}
	updatedPolicy, err := res.PolicyRepository.UpdatePolicy(id, policyToUpdate)
	if err != nil {
		panic(err)
	}
	updatedPolicyDocument, err := converters.ConvertDomainToContract(updatedPolicy)
	if err != nil {
		panic(err)
	}
	return updatedPolicyDocument, nil
}

func (res *PolicyResource) DeletePolicy(id int32) error {
	err := res.PolicyRepository.DeletePolicy(id)
	if err != nil {
		panic(err)
	}
	return nil
}
