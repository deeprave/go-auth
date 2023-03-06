package api

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func logIfError(err error) bool {
	pc, file, line, _ := runtime.Caller(1)
	prefix := fmt.Sprintf("%s(%d):%v", file, line, pc)
	if err != nil {
		log.Printf("%s: %v", prefix, err)
		return true
	}
	return false
}

func isLocalRequest(r *http.Request) bool {
	return r.RemoteAddr == "localhost"
}
