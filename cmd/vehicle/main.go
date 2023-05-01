package main

import (
	"github.com/DamirLuketic/vehicle/config"
	"github.com/DamirLuketic/vehicle/db"
	"github.com/DamirLuketic/vehicle/handlers"
	"log"
	"net/http"
)

func main() {
	c := config.NewServerConfig()
	mdb := db.NewMariaDBDataStore(c)
	apiHandler := handlers.NewApiHandler(mdb, c)
	http.HandleFunc("/api/vehicles", apiHandler.GetVehicles)
	http.HandleFunc("/api/create_vehicle", apiHandler.CreateVehicles)
	http.HandleFunc("/api/delete_vehicle", apiHandler.DeleteVehicles)
	log.Println("Listing for requests at PORT: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
