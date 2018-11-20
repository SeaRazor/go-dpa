package models

import "strings"

type RequestParams struct {
	Filter []string
	Sorting []string
}


func CreateRequestParams(filterOptions []string, sortingOptions []string) RequestParams {
	requestParams := RequestParams{
		Filter:  filterOptions,
		Sorting: sortingOptions,
	}
	return requestParams
}

func (params *RequestParams) CreateSortingString() string{
	var result = strings.Join(params.Sorting, ",")
	return result
}


