package apiserver

import (
	"api/internal/app/stores/sqlStore"
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)
func Start(param *Api) error {
	db, err := newDB(param.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlStore.NewDB(db)
	sessionvalue := sessions.NewCookieStore([]byte(param.SessionKey))
	srv := NewServer(store, sessionvalue)
	return http.ListenAndServe(param.BindAdress,srv)
}

func newDB(database string) (*sql.DB, error) {
	db, err := sql.Open("postgres", database)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}