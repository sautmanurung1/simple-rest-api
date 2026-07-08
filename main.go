package main

import (
	"fmt"
	"log"

	"simple-rest-api/config"
	"simple-rest-api/database"
	"simple-rest-api/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbConf := config.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password", // Sesuaikan dengan password database Anda
		DBName:   "bioskop_db",
	}

	db := database.ConnectDB(dbConf)
	defer db.Close()

	router := gin.Default()

	bioskopHandler := handlers.NewBioskopHandler(db)
	router.POST("/bioskop", bioskopHandler.AddBioskop)

	fmt.Println("Server berjalan di port 3000...")
	if err := router.Run(":3000"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
