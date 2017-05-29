package routes

import (
	"../pages"
	"../utils"
)

func Routes(r utils.GuardedRouter) {
	r.Func("/", Hello)

	r.ProtectFunc("/api/user", ShowUser)

	r.Func("/api/access_token", SignInUser, "POST")

	r.ProtectFunc("/", pages.Home)
}
