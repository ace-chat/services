package service

import (
	"ace/serializer"
	"strconv"
)

type Example struct {
	Code *string `form:"code" json:"code" binding:"required"`
}

func (e *Example) Example() serializer.Response {
	// some logic start
	number, err := strconv.Atoi(*e.Code)
	if err != nil {
		return serializer.RunError(err)
	}
	// some logic end

	example := serializer.Example{
		Code:   *e.Code,
		Number: number,
	}

	return example.Example(example)
}
