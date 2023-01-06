package router

import "net/http"

type Router interface {
	HandleReq(uri string, f func(writer http.ResponseWriter, request *http.Request))
	Serve(port string)
}
