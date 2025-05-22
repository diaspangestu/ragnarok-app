package dto

import "reflect"

type Pagination struct {
	CurrentPage int         `json:"currentPage"`
	LastPage    int         `json:"lastPage"`
	PerPage     int         `json:"perPage"`
	Total       int         `json:"total"`
	Data        interface{} `json:"data"`
}

func SetupPagination(data interface{}, page, limit, totalRecords int) Pagination {
	var pagination Pagination

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Slice && val.Len() == 0 {
		dataType := val.Type()
		data = reflect.MakeSlice(dataType, 0, 0).Interface()
	}

	pagination.Data = data
	pagination.CurrentPage = page
	pagination.PerPage = limit
	pagination.Total = totalRecords

	lastPage := totalRecords / limit
	if totalRecords%limit != 0 {
		lastPage++
	}
	if lastPage == 0 {
		lastPage = 1
	}
	pagination.LastPage = lastPage

	return pagination
}
