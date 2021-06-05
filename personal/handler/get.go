package handler

import (
    "github.com/Projector-Solutions/Pharaon-tools/server"
    "log"
    "net/http"
    "pharaon-personal/service"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
    id, err := server.ReadPathUUID(UserIdKey, r)
    if err != nil {
        log.Println(err)
        server.BadRequest(nil, w)
        return
    }
    
    data, err := service.GetById(id)
    if err != nil {
        log.Println(err)
        server.InternalServerError(w)
        return
    }
    
    if data == nil {
        server.NotFound(w)
        return
    }
    
    server.Ok(data, w)
}
