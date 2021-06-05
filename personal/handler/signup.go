package handler

import (
    "github.com/Projector-Solutions/Pharaon-tools/server"
    "log"
    "net/http"
    "pharaon-personal/model"
    "pharaon-personal/service"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    data := &model.Data{}
    err := server.ReadJson(r, data)
    if err != nil {
        log.Println(err)
        server.BadRequest(nil, w)
        return
    }
    
    result, err := service.Signup(data)
    if err != nil {
        log.Println(err)
        server.InternalServerError(w)
        return
    }
    
    if !result.IsValid() {
        server.BadRequest(result, w)
        return
    }
    
    server.Ok(data, w)
}
