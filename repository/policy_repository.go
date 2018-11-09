package repository

import (
	"dpa/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type PolicyRepository struct {
	Config *models.Configuration
}

func (repo *PolicyRepository) GetPolicy(policyId int32) (models.Policy, error) {
	db, err := gorm.Open("mssql", models.CurrentConfiguration.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var policy models.Policy
	result := db.Where(&models.Policy{ID: policyId}).First(&policy)
	if result.RecordNotFound() {
	}

	return policy, nil
}

func (repo *PolicyRepository) GetPolicies(params models.RequestParams) ([]models.Policy, error) {
	policies := make([]models.Policy, 10)
	return policies, nil
}

func (repo *PolicyRepository) CreatePolicy(policy *models.Policy) (models.Policy, error) {
	createdPolicy := models.Policy{}
	return createdPolicy, nil
}

func (repo *PolicyRepository) UpdatePolicy(id int, policy *models.Policy) (models.Policy, error) {
	updatedPolicy := models.Policy{}
	return updatedPolicy, nil
}

func (repo *PolicyRepository) DeletePolicy(id int) error {
	return nil
}
