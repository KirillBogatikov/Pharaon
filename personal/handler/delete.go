package handler

import (
    "github.com/Projector-Solutions/Pharaon-tools/server"
    "log"
    "net/http"
    "pharaon-personal/service"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    id, err := server.ReadPathUUID(UserIdKey, r)
    if err != nil {
        log.Println(err)
        server.BadRequest(nil, w)
        return
    }
    
    found, err := service.Delete(id)
    if err != nil {
        log.Println(err)
        server.InternalServerError(w)
        return
    }
    
    if !found {
        server.NotFound(w)
        return
    }
    
    server.Ok(nil, w)
}
