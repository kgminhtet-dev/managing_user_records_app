package handler

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
)

type PagingResp struct {
	Previous int `json:"previous"`
	Next     int `json:"next"`
}

type PaginatedResp struct {
	Data   []*data.User `json:"data"`
	Paging PagingResp   `json:"paging"`
}

type ErrorResp struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}
