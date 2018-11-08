package models

//Policy dpa policy
type Policy struct {
	ID           int32     `json:"id"`
	Identifier   string    `json:"name"`
	Tag          string    `json:"tag"`
	TenantID     int32     `json:"tenantId"`
	CreationDate string    `json:"createdAt"`
	Statement1   Statement `json:"statement1"`
	Statement2   Statement `json:"statement2"`
	Statement3   Statement `json:"statement3"`
	Statement4   Statement `json:"statement4"`
	Statement5   Statement `json:"statement5"`
	IsDeleted    bool      `json:"isDeleted"`
	IsEditable   bool      `json:"isEditable"`
}

//Statement dfdf
type Statement struct {
	Text string `json:"text"`
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
