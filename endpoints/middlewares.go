package endpoints

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

var EchoMiddle1 Middleware = func (h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("Echo from example middleware 1")
		fmt.Fprint(w, "Additional content from middleware 1\n")
		
		h(w, r, ps)
	}
}

var EchoMiddle2 Middleware = func (h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("Echo from example middleware 2")
		fmt.Fprint(w, "Additional content from middleware 2\n")
		
		h(w, r, ps)
	}
}

var RequestConsoleLog Middleware = func (h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println(r)

		h(w, r, ps)
	}
}
