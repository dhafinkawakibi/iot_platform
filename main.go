package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kawakibireku/iot_platform/db"
	"github.com/kawakibireku/iot_platform/services"
)

var database *gorm.DB

func main() {

	fmt.Println("Starting Cozy POS backend service on port 8080!")

	services.StartMqtt()

	db.InitDb(&database)
	defer database.Close()

	r := mux.NewRouter()
	// InitRouters(&r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"}})

	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}
