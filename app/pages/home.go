package pages

import "net/http"
import "../models"

func Home(user models.User, rs http.ResponseWriter, rq *http.Request) {
	rs.Write([]byte(user.Name))
}
