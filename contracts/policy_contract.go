package contracts

//Policy dpa policy
type PolicyDocument struct {
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

type PolicyDocumentCollection struct {
	  Documents []PolicyDocument
}

func NewCollection() PolicyDocumentCollection{
    result := PolicyDocumentCollection{}
    result.Documents = []PolicyDocument{}
    return result
}