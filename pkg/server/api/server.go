package api

import (
	"crypto/rand"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/wildnature/macaque/pkg/logger"
)

const (
	serverAddress string = "0.0.0.0:8080"
	basePath      string = "/api/v0"
)

func api(router *mux.Router) {
	notFoundHandler(router)
	apiRouter := router.PathPrefix(basePath).Subrouter()
	anonymousRoutes(apiRouter)
	securedRouter := apiRouter.PathPrefix("/admin").MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		adminToken := r.Header.Get("X-Authorization")
		if adminToken == "" {
			logger.Debug("Missing token.")
			return false
		}
		return adminToken == "ValidToken"
	}).Subrouter()
	securedRoutes(securedRouter)
}

func securedRoutes(router *mux.Router) {
	router.HandleFunc("/schedulers/{schedulerID:[0-9]+}", getSchedulerHandler).Methods(http.MethodGet).Name("GetScheduler")
}

func anonymousRoutes(anonymousRouter *mux.Router) {
	anonymousRouter.HandleFunc("/schedulers", createSchedulerHandler).Methods(http.MethodPost).Name("CreateScheduler")
}

//ConfigureServerAndRun initial function
func ConfigureServerAndRun() {
	logger.Info("Configuring server")
	router := mux.NewRouter().StrictSlash(false)
	api(router)
	k, err := randomKey()
	if err != nil {
		log.Fatal("Error while generating csrf key.")
	}
	CSRF := csrf.Protect(
		k,
	)
	server := &http.Server{
		Handler: CSRF(router),
		Addr:    serverAddress,
	}
	log.Printf("Launching server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func randomKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
