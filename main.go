package main

import (
	"Proyek-Akhir-Golang/db"
	"Proyek-Akhir-Golang/router"
)

func main() {
	db.StartDB()
	router.StartServer()
}
