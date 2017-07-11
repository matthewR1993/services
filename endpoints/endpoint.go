// Description of endpoint type

package endpoints

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct {
	Path string
	Method string
	Name string
	Handler Handler
	EndpointMiddlewares []Middleware  // middleware handlers act before final Handler
}

type Handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)

type EndpointUtils interface {
	AttachMiddleware(mldwres []Middleware)
}

func (e *Endpoint) AttachMiddleware(mdlwres ...Middleware) {
	e.EndpointMiddlewares = append(e.EndpointMiddlewares, mdlwres...)
}
