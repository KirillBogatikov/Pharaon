package handler

import (
	tool "github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pharaon-migration/data"
	"pharaon-migration/service"
)

func RollbackMigrationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	migration := data.Migration{}
	serviceError := tool.ReadJson(r, &migration)
	if serviceError != nil {
		log.Println(serviceError)
		tool.InternalServerError(w)
		return
	}

	serviceName := vars["service"]
	v, databaseError, serviceError := service.Rollback(serviceName, migration)
	if serviceError == data.ServiceNotFound {
		tool.NotFound(w)
		return
	}

	if serviceError != nil {
		log.Println(serviceError)
		tool.InternalServerError(w)
		return
	}

	result := data.MigrationResult{Service: serviceName, Version: v}
	if databaseError != nil {
		errText := databaseError.Error()
		result.Error = &errText
		tool.BadRequest(result, w)
		return
	}

	tool.Ok(result, w)
}
