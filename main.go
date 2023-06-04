package main

import (
	"fmt"
	"github.com/cubbit/cbs3/internal/datastore"
	"github.com/cubbit/cbs3/internal/delivery/router"
	"github.com/cubbit/cbs3/internal/registry"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		fmt.Print("sad .env file found")
	}
}

func main() {

	db := datastore.NewPostgreSQL()
	rg := registry.NewInteractor(db)
	h := rg.NewAppHandler()
	g := router.NewRouter(h)

	go g.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)
	<-stop

}
