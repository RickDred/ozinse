package models

import (
	"math"
	"strings"
)

const (
	Asc  = "ASC"
	Desc = "DESC"
)

type Metadata struct {
	CurrentPage  int
	PageSize     int
	FirstPage    int
	LastPage     int
	TotalRecords int
}

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func (f Filters) SortColumn() (string, error) {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.Trim(f.Sort, "-"), nil
		}
	}

	return "", nil
}

func (f Filters) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return Desc
	}
	return Asc
}

func (f Filters) Limit() int {
	return f.PageSize
}

func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func CalculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}
