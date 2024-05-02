package api

import (
	"database/sql"
)

type Config struct {
	Db      *sql.DB
	BaseUrl string
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type jsonRequest struct {
	LongUrl string `json:"longUrl"`
}
