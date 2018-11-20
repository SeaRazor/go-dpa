package models

//Policy dpa policy
type Policy struct {
	ID           int32     `gorm:"column:PolicyId;primary_key"`
	Identifier   string    `gorm:"column:Identifier"`
	Tag          string    `gorm:"column:Tag"`
	TenantID     int32     `gorm:"column:TenantId"`
	CreationDate string    `gorm:"column:CreationDate"`
	Statement1   string    `gorm:"column:Statement1"`
	Statement2   string    `gorm:"column:Statement2"`
	Statement3   string    `gorm:"column:Statement3"`
	Statement4   string    `gorm:"column:Statement4"`
	Statement5   string    `gorm:"column:Statement5"`
	IsDeleted    bool      `gorm:"column:IsDeleted"`
	IsEditable   bool      `gorm:"column:IsEditable"`
}

func(Policy) TableName() string{
	return "TPolicy"
}

func (policy *Policy) UpdateValues(source *Policy){
	policy.Identifier = source.Identifier
	policy.Statement1 = source.Statement1
	policy.Statement2 = source.Statement2
	policy.Statement3 = source.Statement3
	policy.Statement4 = source.Statement4
	policy.Statement5 = source.Statement5
	policy.Tag = source.Tag
	policy.TenantID = source.TenantID
}




//StatementAnswer answer
type StatementAnswer struct {
	Text   string `json:"text"`
	Answer bool   `json:"answer"`
}

//Agreement client agreements
type Agreement struct {
	ID            int32           `json:"id"`
	TenantID      int32           `json:"tenantId"`
	CRMContactID  int32           `json:"CrmContactId"`
	AgreementDate string          `json:"agreement_date"`
	Statement1    StatementAnswer `json:"statement1"`
	Statement2    StatementAnswer `json:"statement2"`
	Statement3    StatementAnswer `json:"statement3"`
	Statement4    StatementAnswer `json:"statement4"`
	Statement5    StatementAnswer `json:"statement5"`
	Policy        Policy          `json:"policy"`
}

func (Agreement) TableName() string{
	return "TAgreement"
}