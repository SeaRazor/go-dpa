package validators

import (
	"dpa/models"
)

type Validator interface{
    Validate(v interface{})
}

type PolicyValidator struct{

}

func (policy *models.Policy) Validate() error  {
    
}