package repository

import (
	"dpa/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type IPolicyRepository interface {
	GetPolicy(policyId int32) (models.Policy, error)
	GetPolicies(params models.RequestParams) ([]models.Policy, error)
	CreatePolicy(policy *models.Policy) (models.Policy, error)
	UpdatePolicy(id int, policy *models.Policy) (models.Policy, error)
	DeletePolicy(id int) error
}
type PolicyRepository struct {
	Config *models.Configuration

}
var db *gorm.DB

func init ()  {
	conn, err := gorm.Open("mssql", models.CurrentConfiguration.DatabaseConnectionString)
	if err != nil{
		panic(err)
	}
	db = conn
	//db.LogMode(true)
}

func (repo *PolicyRepository) GetPolicy(policyId int32) (models.Policy, error) {
	var policy models.Policy
	result := db.First(&policy, policyId)
	return policy, result.Error
}

func (repo *PolicyRepository) GetPolicies(params models.RequestParams) ([]models.Policy, error) {
	policies := make([]models.Policy, 0)
	var result *gorm.DB
	if len(params.Sorting) > 0 {
       result = db.Order(params.CreateSortingString()).Find(&policies)
		return policies, result.Error
	}
	result = db.Find(&policies)
	return policies, result.Error
}

func (repo *PolicyRepository) CreatePolicy(policy *models.Policy) (models.Policy, error) {
	db.NewRecord(policy)
	var newPolicy models.Policy
	result:= db.Create(&newPolicy)
	return newPolicy, result.Error
}

func (repo *PolicyRepository) UpdatePolicy(id int32, policy *models.Policy) (models.Policy, error) {
	policyToUpdate,err := repo.GetPolicy(id)
	if err != nil {
		return *policy, err
	}
	policyToUpdate.UpdateValues(policy)
	var updatedPolicy models.Policy
	result := db.Save(&updatedPolicy)
	return updatedPolicy, result.Error
}

func (repo *PolicyRepository) DeletePolicy(id int32) error {
	policyToDelete,err := repo.GetPolicy(id)
	if err != nil {
		return err
	}
	result := db.Delete(&policyToDelete)
	return result.Error
}
