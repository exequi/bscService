package model

type jsonError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type myError jsonError
