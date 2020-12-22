package utils

type PageInfo struct {
	Size int `json:"size"`
	Page int `json:"page"`
}

type Page struct {
	Count int `json:"count"`
	Page int `json:"page"`
	Size int `json:"size"`
}

const DefaultSize = 10
const DefaultPage = 1