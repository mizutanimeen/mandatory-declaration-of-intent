package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mizutanimeen/mandatory-declaration-of-intent/db"
	logFile "github.com/mizutanimeen/mandatory-declaration-of-intent/logs"
	"github.com/mizutanimeen/mandatory-declaration-of-intent/server"
)

func main() {
	//ログ出力先設定
	tLogFile, tError := logFile.Settings("./logs/data/go.log")
	if tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}
	defer func(aLogFile *os.File) {
		if tError := aLogFile.Close(); tError != nil {
			log.Fatalf("%s\n", tError.Error())
		}
	}(tLogFile)

	//DB接続
	tDb, tError := db.Open()
	if tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}
	defer func(aDb *sql.DB) {
		if tError := aDb.Close(); tError != nil {
			log.Fatalf("%s\n", tError.Error())
		}
	}(tDb)

	tServer := server.New(tDb)

	//ルーティング
	tRouter := chi.NewRouter()
	tRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("REACT_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "X-Requested-With", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"ExposedHeaders", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	tRouter.Use(middleware.Logger)
	tRouter.Route(os.Getenv("REACT_APP_GO_PATH"), func(tRouter chi.Router) {
		tRouter.Get("/users/{ID}", tServer.GetUserByID)
		tRouter.Post("/users", tServer.CreateUser)
		tRouter.Get("/rooms/{ID}", tServer.GetRoomByID)
		tRouter.Post("/rooms", tServer.CreateRoom)
		tRouter.Get("/rooms/{ID}/members/gest", tServer.GetAllGestMembers)
		tRouter.Get("/rooms/{ID}/check", tServer.PasswordCheck)
		// tRouter.Post("/rooms/members/user", tServer.JoinUserMember)
		tRouter.Post("/rooms/members/gest", tServer.CreateGestMember)
	})

	log.Println("--------------Server Start--------------")
	if tError := http.ListenAndServe(":8080", tRouter); tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}
}
