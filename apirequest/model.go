package apirequest

import (
)
type RequestModel struct {
	Requestid int `json:"request_id"`

	Urlpackage []int `json:"url_package"`

	Ip string `json:"ip"`
}

type RequestPrice struct {
	Price float64 `json:"price"`
}



