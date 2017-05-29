package pages

import "net/http"
import "../models"
import "../utils"

func User(user models.User, rs http.ResponseWriter, rq *http.Request) {
	bag := utils.NewBag(user)

	utils.JsonResp(rs, bag)

	rs.Write([]byte(user.Name))
}
