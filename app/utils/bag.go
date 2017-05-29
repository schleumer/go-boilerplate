package utils

import (
	"net/http"
)

type BagError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

type BagMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

type Bag struct {
	Data     interface{} `json:"data"`
	Errors   []BagError `json:"errors"`
	Messages []BagMessage `json:"messages"`
}

func NewBag(data interface{}) *Bag {
	return &Bag{
		data,
		make([]BagError, 0),
		make([]BagMessage, 0),
	}
}

func (bag *Bag) AddError(code string, message string, field string) *Bag {

	bag.Errors = append(bag.Errors, BagError{
		code,
		message,
		field,
	})

	return bag
}

func (bag *Bag) WriteTo(rs http.ResponseWriter) {
	if len(bag.Errors) > 0 {
		JsonRespWithStatus(rs, bag, http.StatusBadRequest)
	} else {
		JsonRespWithStatus(rs, bag, http.StatusOK)
	}


}