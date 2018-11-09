package models

type RequestParams struct {
	Filter  Filter
	Sorting Sorting
}

type Sorting struct {
	SortingOptions []string
}

type Filter struct {
	FilterOptions []string
}

func CreateRequestParams(filterOptions []string, sortingOptions []string) RequestParams {
	requestParams := RequestParams{
		Filter:  CreateFilter(filterOptions),
		Sorting: CreateSorting(sortingOptions),
	}
	return requestParams
}

func CreateFilter(filterOptions []string) Filter {
	filter := Filter{FilterOptions: filterOptions}
	return filter
}

func CreateSorting(sortingOptions []string) Sorting {
	sorting := Sorting{SortingOptions: sortingOptions}
	return sorting
}
