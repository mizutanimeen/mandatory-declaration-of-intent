package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	tRoom, tStatus, tError := db.GetRoomByID(tID, aServer.Db)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	tRoomByte, tError := json.Marshal(tRoom)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	aResponseWriter.Header().Set("Content-Type", "application/json")
	aResponseWriter.WriteHeader(200)
	if _, tError := aResponseWriter.Write(tRoomByte); tError != nil {
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
	aResponseWriter.Header().Set("Content-Type", "application/json")
	if _, tError := aResponseWriter.Write([]byte(fmt.Sprintf("http://localhost:8123/mandatory-declaration-of-intent/api/v1/rooms/%s", tRoom.RoomID))); tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}
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
func (aServer *Server) GetAllGestMembers(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tID := chi.URLParam(aRequest, "ID")

	// クッキーチェック
	tCookie, err := aRequest.Cookie("mycookie")
	if err != nil {
		http.Error(aResponseWriter, "Access denied:cookie none", http.StatusForbidden)
		return
	}

	if tCookie.Value != "Hello, World!" {
		http.Error(aResponseWriter, fmt.Sprintf("Access denied: your cookie is %s", tCookie.Value), http.StatusForbidden)
		return
	}
	// クッキーチェック終わり

	tGestUsers, tStatus, tError := db.GetAllGestMembers(tID, aServer.Db)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	tUserByte, tError := json.Marshal(tGestUsers)
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
func (aServer *Server) CreateGestMember(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tGestUser, tStatus, tError := controllers.ParseRequestGestUser(aRequest)
	if tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	//リクエストの中身が要件を満たしているか確認

	//データベースに保存
	//データが存在していたときの例外処理、react側が欲しいなら渡すように
	if tStatus, tError := db.CreateGestUser(tGestUser, aServer.Db); tError != nil {
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	//tGestUser.RoomIDからRoomID.CookieName,RoomID.CookieValueを取得
	tRoom, tStatus, tError := db.GetRoomByID(tGestUser.RoomID, aServer.Db)
	tCookie := &http.Cookie{
		Name:  tRoom.CookieName,
		Value: tRoom.CookieValue,
	}
	http.SetCookie(aResponseWriter, tCookie)
	aResponseWriter.WriteHeader(http.StatusCreated)
	return
}
