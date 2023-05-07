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
		tRouter.Get("/users/{ID}", tServer.GetUserByID)                    // 永久ユーザーの情報をIDから取得
		tRouter.Post("/users", tServer.CreateUser)                         // 永久ユーザー作成
		tRouter.Get("/rooms/{ID}", tServer.GetRoomByID)                    // IDから部屋の情報を取得
		tRouter.Post("/rooms", tServer.CreateRoom)                         // 部屋作成
		tRouter.Get("/rooms/{ID}/members/gest", tServer.GetAllGestMembers) // IDから部屋のゲストメンバーを全て取得
		tRouter.Post("/rooms/members/gest", tServer.CreateGestMember)      // 送信された部屋IDに送信された情報のゲストユーザーを作成 // idをURLに含めたい
		tRouter.Get("/rooms/{ID}/check", tServer.PasswordCheck)            // 送信されたパスワードが正しいかチェック
		// tRouter.Post("/rooms/members/user", tServer.JoinUserMember)
	})

	log.Println("--------------Server Start--------------")
	if tError := http.ListenAndServe(":8080", tRouter); tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}
}
