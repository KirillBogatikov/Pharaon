package main

import (
	"encoding/json"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"github.com/Projector-Solutions/Pharaon-tools/server"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"pharaon-proxy/header"
	mongoLog "pharaon-proxy/log"
	settings "pharaon-proxy/settings"
	"regexp"
)

type Proxy struct {
	logger   *mongoLog.Logger
	services map[string]string
	settings settings.Settings
}

func NewProxy(logger *mongoLog.Logger, set settings.Settings) *Proxy {
	return &Proxy{
		logger:   logger,
		services: make(map[string]string),
		settings: set,
	}
}

func (p *Proxy) Register(name, addr string) {
	p.services[name] = addr
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	clientIp := header.GetClientAddr(r)
	serviceName, serviceURI := header.GetProxyService(r)
	log.Println(serviceName, serviceURI)

	record := &mongoLog.Record{
		Id:      uuid.New(),
		Client:  header.GetClientAddr(r),
		Remote:  r.RemoteAddr,
		Service: serviceName,
		Request: &mongoLog.RequestInfo{
			Method:  r.Method,
			Length:  header.GetContentLength(r.Header),
			URI:     r.RequestURI,
			Headers: header.GetHeadersMap(r.Header),
		},
	}

	defer func() {
		data, _ := json.Marshal(record)
		log.Println(string(data))
		go p.logger.Log(record)
	}()

	httpClient := &http.Client{}
	r.RequestURI = ""

	header.ClearHeader(r.Header)
	header.AppendProxyHost(r.Header, clientIp)

	addr, ok := p.services[serviceName]
	if !ok {
		record.Unknown = true
		log.Printf("Unknown service %s\n", serviceName)
		server.NotFound(w)
		return
	}

	settings, ok := p.settings[serviceName]
	if ok {
		for _, pattern := range settings.Hidden {
			result, _ := regexp.MatchString(pattern, serviceURI)

			if result {
				record.Forbidden = true
				log.Printf("Forbidden url %s\n", serviceName)
				server.NotFound(w)
			}
		}
	}

	r.URL = client.MustFormatURL("%s/%s", addr, serviceURI)

	resp, err := httpClient.Do(r)
	if err != nil {
		log.Println(err)
		server.InternalServerError(w)
		return
	}

	header.ClearHeader(resp.Header)
	header.CopyHeader(resp.Header, w.Header())

	record.Response = &mongoLog.ResponseInfo{
		Length:  header.GetContentLength(resp.Header),
		Status:  resp.StatusCode,
		Headers: header.GetHeadersMap(resp.Header),
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println(err)
	}
}
