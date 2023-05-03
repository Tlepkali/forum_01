package post

import "math"

type Metadata struct {
	CurrentPage int
	FirstPage   int
	LastPage    int
	TotalPages  int
	TotalPosts  int
}

type Filters struct {
	Page     int
	PageSize int
	Sort     string
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage: page,
		FirstPage:   1,
		LastPage:    int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalPages:  pageSize,
		TotalPosts:  totalRecords,
	}
}

func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offset() int {
	return f.PageSize * (f.Page - 1)
}
