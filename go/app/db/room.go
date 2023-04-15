package db

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mizutanimeen/mandatory-declaration-of-intent/models"
)

func GetRoomByID(aID string, aDb *sql.DB) (*models.Room, int, error) {
	tWhere := fmt.Sprintf("WHERE %s = ?", os.Getenv("MYSQL_ROOMS_TABLE_ID"))
	tQuery := fmt.Sprintf("SELECT * FROM %s %s", os.Getenv("MYSQL_ROOMS_TABLE"), tWhere)

	tRoom := models.Room{}
	if tError := aDb.QueryRow(tQuery, aID).Scan(&tRoom.RoomID, &tRoom.Name, &tRoom.Description, &tRoom.CookieName, &tRoom.CookieValue, &tRoom.CreateAt, &tRoom.UpdateAt); tError != nil {
		return nil, http.StatusInternalServerError, tError
	}
	return &tRoom, 0, nil
}

func CreateRoom(aRoom *models.Room, aDb *sql.DB) (int, error) {
	tColumns := fmt.Sprintf("(%s,%s,%s,%s,%s)",
		os.Getenv("MYSQL_ROOMS_TABLE_ID"), os.Getenv("MYSQL_ROOMS_TABLE_NAME"), os.Getenv("MYSQL_ROOMS_TABLE_DESCRIPTION"), os.Getenv("MYSQL_ROOMS_TABLE_COOKIE_NAME"), os.Getenv("MYSQL_ROOMS_TABLE_COOKIE_VALUE"))
	tQuery := fmt.Sprintf("INSERT INTO %s %s VALUES (?,?,?,?,?)",
		os.Getenv("MYSQL_ROOMS_TABLE"), tColumns)

	tStmt, tError := aDb.Prepare(tQuery)
	if tError != nil {
		return http.StatusInternalServerError, tError
	}
	defer tStmt.Close()

	tUUID, tStatus, tError := createUUID()
	if tError != nil {
		return tStatus, tError
	}
	aRoom.RoomID = tUUID

	tCookieName, tCookieValue, tStatus, tError := createCookie(aRoom)
	if tError != nil {
		return tStatus, tError
	}

	if _, tError := tStmt.Exec(aRoom.RoomID, aRoom.Name, aRoom.Description, tCookieName, tCookieValue); tError != nil {
		return http.StatusInternalServerError, tError
	}
	return 0, nil
}

func createUUID() (string, int, error) {
	tUuid := make([]byte, 16)
	if n, tError := rand.Read(tUuid); n != len(tUuid) || tError != nil {
		return "", http.StatusInternalServerError, tError
	}
	tUuid[8] = tUuid[8]&^0xc0 | 0x80
	tUuid[6] = tUuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x", tUuid), 0, nil
}

func createCookie(aRoom *models.Room) (string, string, int, error) {
	tUuid := make([]byte, 16)
	if n, tError := rand.Read(tUuid); n != len(tUuid) || tError != nil {
		return "", "", http.StatusInternalServerError, tError
	}
	tUuid[8] = tUuid[8]&^0xc0 | 0x80
	tUuid[6] = tUuid[6]&^0xf0 | 0x40
	tCookieName := fmt.Sprintf("mandatory-declaration-of-intent-room-%s", aRoom.RoomID)
	tCookieValue := fmt.Sprintf("%x", tUuid)

	return tCookieName, tCookieValue, 0, nil
}
