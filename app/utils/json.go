package utils

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func JsonResp(rs http.ResponseWriter, what interface{}) {
	js, _ := json.Marshal(what)
	rs.Header().Add("Content-Type", "application/json")
	rs.Write(js)
}

func JsonRespWithStatus(rs http.ResponseWriter, what interface{}, status int) {
	js, _ := json.Marshal(what)
	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(status)
	rs.Write(js)
}

func JsonReq(rq *http.Request, output interface{}) {
	b, _ := ioutil.ReadAll(rq.Body)

	json.Unmarshal(b, &output)
}