package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	"github.com/mizutanimeen/mandatory-declaration-of-intent/controllers"
	"github.com/mizutanimeen/mandatory-declaration-of-intent/db"
)

type Server struct {
	Db *sql.DB
}

func New(aDb *sql.DB) *Server {
	return &Server{
		Db: aDb,
	}
}

func (aServer *Server) GetRoomByID(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tID := chi.URLParam(aRequest, "ID")

	tUser, tStatus, tError := db.GetRoomByID(tID, aServer.Db)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	tUserByte, tError := json.Marshal(tUser)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	aResponseWriter.Header().Set("Content-Type", "application/json")
	aResponseWriter.WriteHeader(200)
	if _, tError := aResponseWriter.Write(tUserByte); tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (aServer *Server) CreateRoom(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tRoom, tStatus, tError := controllers.ParseRequestRoom(aRequest)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	//リクエストの中身が要件を満たしているか確認

	//データベースに保存
	//データが存在していたときの例外処理、react側が欲しいなら渡すように
	if tStatus, tError := db.CreateRoom(tRoom, aServer.Db); tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}
	aResponseWriter.WriteHeader(http.StatusCreated)
	return
}

func (aServer *Server) GetUserByID(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tID := chi.URLParam(aRequest, "ID")

	tUser, tStatus, tError := db.GetUserByID(tID, aServer.Db)
	if tError != nil {
		http.Error(aResponseWriter, tError.Error(), tStatus)
	}

	tUserByte, tError := json.Marshal(tUser)
	if tError != nil {
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	aResponseWriter.Header().Set("Content-Type", "application/json")
	aResponseWriter.WriteHeader(200)
	if _, tError := aResponseWriter.Write(tUserByte); tError != nil {
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (aServer *Server) CreateUser(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tUser, tStatus, tError := controllers.ParseRequestUser(aRequest)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	//リクエストの中身が要件を満たしているか確認

	//データベースに保存
	//データが存在していたときの例外処理、react側が欲しいなら渡すように
	if tStatus, tError := db.CreateUser(tUser, aServer.Db); tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}
	aResponseWriter.WriteHeader(http.StatusCreated)
	return
}
