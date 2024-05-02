package database

import (
	"database/sql"
	"shorturl/app"
	"sync"
)

var mutex sync.Mutex

func CreateShortUrl(conn *sql.DB, longUrl string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var existingShortUrl string

	// check if longurl already exists
	queryExists := `SELECT shorturl FROM "UrlStore" WHERE longUrl = $1`
	_ = conn.QueryRow(queryExists, longUrl).Scan(&existingShortUrl)

	if existingShortUrl != "" {
		return existingShortUrl, nil
	}

	// take next value
	query := `SELECT nextval('urlstore_id_seq')`

	var nextValId int

	err := conn.QueryRow(query).Scan(&nextValId)
	if err != nil {
		return "", err
	}

	shortUrl := app.EncodeBase58(int64(nextValId))

	query = `INSERT INTO "UrlStore" (id, longurl, shorturl) VALUES ($1, $2, $3)`

	_, err = conn.Query(query, nextValId, longUrl, shortUrl)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func GetLongUrl(conn *sql.DB, shortUrl string) (string, error) {

	var longUrl string

	query := `SELECT longUrl FROM "UrlStore" WHERE shortUrl = $1`

	err := conn.QueryRow(query, shortUrl).Scan(&longUrl)
	if err != nil {
		return "", err
	}

	return longUrl, nil
}
