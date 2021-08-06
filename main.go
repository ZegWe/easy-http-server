package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	str, _ := os.Getwd()
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	}else {
		port = ":8080"
	}
	log.Println(str)
	fs := http.FileServer(http.Dir(str))
	f := handlers.CustomLoggingHandler(log.Writer(), fs, func(writer io.Writer, params handlers.LogFormatterParams) {
		str:= fmt.Sprintf("%v %v %v %v\n",
			params.TimeStamp.Local().Format("2006/01/02 15:04:05"),
			params.StatusCode,
			params.Request.Method,
			params.Request.URL,
		)
		writer.Write([]byte(str))
	})
	http.Handle("/", f)
	log.Panic(http.ListenAndServe(port, nil))
}
