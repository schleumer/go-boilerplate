package routes

import (
	"net/http"
)

func Hello(rs http.ResponseWriter, rq *http.Request) {
	rs.Write([]byte("Hello World"))
}

