package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/wildnature/macaque/pkg/logger"
	"github.com/wildnature/macaque/pkg/rest"
)

func complete(res http.ResponseWriter, httpStatus int, content interface{}) {
	body, err := json.Marshal(content)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set(rest.HeaderContentType, rest.JSONContentType)
	res.WriteHeader(httpStatus)
	res.Write(body)
}

func notFoundHandler(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		content := rest.HTTPError("ERR00001", "Invalid url")
		complete(res, http.StatusNotFound, content)
	})
}

func missingCSRFHandler(res http.ResponseWriter, req *http.Request) {
	content := rest.HTTPError("ERR00001", "Missing CSRF")
	complete(res, http.StatusNotFound, content)

}

func createSchedulerHandler(res http.ResponseWriter, req *http.Request) {
	complete(res, http.StatusNotImplemented, nil)
}

func getSchedulerHandler(res http.ResponseWriter, req *http.Request) {
	logger.Info("getSchedulerHandler")
	vars := mux.Vars(req)
	schedulerID, _ := strconv.Atoi(vars["schedulerID"])
	logger.Debugf("Searching for scheduler with id %s", schedulerID)
	if schedulerID == 0 {
		content := rest.HTTPError("ERR00003", fmt.Sprintf("There's not a scheduler with ID  %x", schedulerID))
		complete(res, http.StatusNotFound, content)
	} else {
		content := &Scheduler{
			Name: "my-scheduler",
		}
		complete(res, http.StatusOK, content)
	}
}
