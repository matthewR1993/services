// This file consist of set of endpoints

package endpoints

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	log "github.com/sirupsen/logrus"
)

/*
  Echo test endpoint
*/
var TestEcho = Endpoint {
	"/echo",
	"GET",
	"<Echo>",
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "echo")
		log.Warning("some warning")
	},
	[]Middleware{},
	// e. g.:
	// []Middleware{EchoMiddle1, EchoMiddle2},
	// Evaluation order:  [ [SharedMiddlewares,...] --> Middleware2 --> Middleware1 --> FinalHandler ]
}

/*
  Main page
*/
var Main = Endpoint {
	"/",
	"GET",
	"<Main page>",
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w,
			`<!DOCTYPE html>
			<html lang="en">
			<head>
			  <meta charset="utf-8">
			  <title>Main</title>
			</head>
			<body>
			  <div id="root"></div>
			  <script src="/static/bundle.js"></script>
			</body> 
			</html>`)
		
		// Push static files back immediately
		pusher, ok := w.(http.Pusher)
		if ok {
			fmt.Println("Pushing assets")
			if err := pusher.Push("/static/bundle.js", nil); err != nil {
				log.Warning("Failed to push: %v", err)
			}
		}
	},
	[]Middleware{},
}


