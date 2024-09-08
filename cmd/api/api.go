package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/soufianiso/boxd/services/user"
	"github.com/soufianiso/boxd/services/movies"
	"github.com/soufianiso/boxd/utils"
)

func NewServer(logger *log.Logger, db *sql.DB) http.Handler{
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStorage(db)
	user.SetRoutes(apiRouter, userStore, logger)

	moviesStore := movies.NewStorage(db)
	movies.SetRoutes(apiRouter, moviesStore, logger)
	// here is top level middleware stuff
	var handler http.Handler
	handler = utils.CORSMiddleware(router)
	return handler
}




