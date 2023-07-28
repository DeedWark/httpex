/* Author: DeedWark <github.com/DeedWark>
 * @2023
 * v1
 */

package httpex

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
)

type ServeOpt struct {
	/*
		Uri:       "/",
		Addr:      ":8080",
		WebPage:   myFunc,
		LogMethod: os.Stdout,
		LogFormat: "CombinedLoggingHandler",
	*/
	Uri       string
	Addr      string
	WebPage   func(w http.ResponseWriter, req *http.Request)
	LogMethod io.Writer
	LogFormat string
}

/* __ from github.com/gorilla/handlers -> __
 * LogginHandler:          Apache Common Log Format.
 * CombinedLoggingHandler: Apache Combined Log Format (commonly used by both Apache and NGINX).
 * CompressHandler:        GZIP.
 */

func (s ServeOpt) Serve() {
	mux := http.NewServeMux()
	mux.Handle(s.Uri, handlers.LoggingHandler(s.LogMethod, http.HandlerFunc(s.WebPage)))

	var logHandler http.Handler
	switch s.LogFormat {
	case "LoggingHandler", "CommonLogFormat":
		logHandler = handlers.LoggingHandler(s.LogMethod, mux)
	case "CombinedLoggingHandler", "CombinedLogFormat":
		logHandler = handlers.CombinedLoggingHandler(s.LogMethod, mux)
	case "CompressHandler":
		logHandler = handlers.CompressHandler(mux)
	default:
		logHandler = handlers.CompressHandler(mux)
	}

	http.ListenAndServe(s.Addr, logHandler)
}

// Example:
// func thisIsAnExample() {
// 	opt := httpex.ServeOpt{
// 		Uri:       "/",
// 		Addr:      ":8080",
// 		WebPage:   myPage,
// 		LogMethod: os.Stderr, // Can be a filename to send log to***
//      LogFormat: "CompressHandler"
// 	}
// 	opt.Serve()
// }
// *** logFile, _ := os.OpenFile("file.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// *** defer logFile.Close()
// *** then replace os.Stderr by logFile
