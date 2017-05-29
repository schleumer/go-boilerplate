package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"./database"
	"./routes"
	"./settings"
	"./utils"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/namsral/flag"
	"github.com/rs/cors"
)

var port string

func availablePort() string {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	l.Close()

	return strconv.Itoa(l.Addr().(*net.TCPAddr).Port)
}

func getAddr(port string) string {
	return fmt.Sprintf("127.0.0.1:%s", port)
}

func main() {
	settings.BootSettings()

	flag.StringVar(&port, "port", "", "Port which server will run")

	if port == "" {
		port = availablePort()
	}

	flag.Parse()

	database.BootDb()

	r := mux.NewRouter()

	guardedRouter := utils.NewGuardedRouter(r)

	routes.Routes(guardedRouter)

	handler := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:     []string{"Authorization", "Accept", "Origin", "Content-Type"},
		OptionsPassthrough: false,
		Debug:              true,
	}).Handler(guardedRouter.MuxRouter)

	srv := &http.Server{
		Handler: handler,
		Addr:    getAddr(port),
	}

	fmt.Printf("Running on %s\n", getAddr(port))

	log.Fatal(srv.ListenAndServe())
}

func Run() {
	main()
}
