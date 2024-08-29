package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"shorturl/database"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func (cfg *Config) getTime(w http.ResponseWriter, r *http.Request) {

	time, err := json.Marshal(time.Now())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "",
		Data:    string(time),
	}

	writeJSON(w, http.StatusOK, resp)

}

func (cfg *Config) shorten(w http.ResponseWriter, r *http.Request) {

	var requestPayload jsonRequest
	err := readJSON(w, r, &requestPayload)
	if err != nil {
		writeError(w, http.StatusNotAcceptable, errors.New("error unmarshaling Url"))
		return
	}

	if requestPayload.LongUrl == "" {
		writeError(w, http.StatusNotAcceptable, errors.New("url empty"))
		return
	}

	if !strings.HasPrefix(strings.ToLower(requestPayload.LongUrl), "http") {
		writeError(w, http.StatusNotAcceptable, errors.New("url must starts with http:// or https://"))
		return
	}

	shortUrl, err := database.CreateShortURL(cfg.Db, strings.ToLower(requestPayload.LongUrl))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "",
		Data:    string(cfg.BaseUrl + "/" + shortUrl),
	}

	writeJSON(w, http.StatusOK, resp)

}

func (cfg *Config) redirect(w http.ResponseWriter, r *http.Request) {

	shortUrl := chi.URLParam(r, "shortUrl")
	if shortUrl == "" {
		writeError(w, http.StatusNotAcceptable, errors.New("URL empty"))
		return
	}

	longUrl, err := database.GetLongURL(cfg.Db, shortUrl)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	http.Redirect(w, r, longUrl, http.StatusSeeOther)

}
