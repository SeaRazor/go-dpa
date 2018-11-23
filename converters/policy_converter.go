package converters

import (
	"dpa/contracts"
	"dpa/models"
)

func ConvertDomainToContract(domain models.Policy) (contracts.PolicyDocument, error){
	var document = contracts.PolicyDocument{}
	document.ID = domain.ID
	document.TenantID = domain.TenantID
	document.Tag = domain.Tag
	document.Identifier = domain.Identifier
	document.CreationDate = domain.CreationDate
	document.IsDeleted = domain.IsDeleted
	document.IsEditable = domain.IsEditable
	document.Statement1 = contracts.Statement{Text:domain.Statement1}
	document.Statement2 = contracts.Statement{Text:domain.Statement2}
	document.Statement3 = contracts.Statement{Text:domain.Statement3}
	document.Statement4 = contracts.Statement{Text:domain.Statement4}
	document.Statement5 = contracts.Statement{Text:domain.Statement5}
	return document,nil
}

func ConvertDomainToContractCollection(domains []models.Policy)(contracts.PolicyDocumentCollection, error){
	documentCollection := contracts.PolicyDocumentCollection{}

	for _,domain := range domains{
		var document,err = ConvertDomainToContract(domain)
		if err != nil{
			return contracts.NewCollection(),err
		}
		documentCollection.Documents = append(documentCollection.Documents, document)
	}
	return documentCollection, nil
}

func ConvertContractToDomain(document contracts.PolicyDocument) (models.Policy, error) {
	domain := models.Policy{}
	domain.ID = document.ID
	domain.CreationDate = document.CreationDate
	domain.Identifier = document.Identifier
	domain.Tag = document.Tag
	domain.TenantID = document.TenantID
	domain.IsEditable = document.IsEditable
	domain.IsDeleted = document.IsDeleted
	domain.Statement1 = document.Statement1.Text
	domain.Statement2 = document.Statement2.Text
	domain.Statement3 = document.Statement3.Text
	domain.Statement4 = document.Statement4.Text
	domain.Statement5 = document.Statement5.Text
	return domain, nil
}

