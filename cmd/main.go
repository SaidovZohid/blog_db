package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/blog_db/api"
	"github.com/SaidovZohid/blog_db/config"
	"github.com/SaidovZohid/blog_db/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/SaidovZohid/blog_db/api/docs"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	fmt.Println("Configuration: ", cfg)
	fmt.Println("Connected Succesfully!")

	strg := storage.NewStoragePg(psqlConn)

	apiServer := api.New(&api.RoutetOptions{
		Cfg:     &cfg,
		Storage: strg,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %s", err)
	}

	log.Print("Server Stopped!")
}
