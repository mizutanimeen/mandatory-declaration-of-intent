package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mizutanimeen/mandatory-declaration-of-intent/models"
)

func CreateRoom(aRoom *models.Room, aDb *sql.DB) (int, error) {
	tColumns := fmt.Sprintf("(%s,%s)",
		os.Getenv("MYSQL_ROOMS_TABLE_NAME"), os.Getenv("MYSQL_ROOMS_TABLE_DESCRIPTION"))
	tQuery := fmt.Sprintf("INSERT INTO %s %s VALUES (?,?)",
		os.Getenv("MYSQL_ROOMS_TABLE"), tColumns)

	tStmt, tError := aDb.Prepare(tQuery)
	if tError != nil {
		return http.StatusInternalServerError, tError
	}
	defer tStmt.Close()

	if _, tError := tStmt.Exec(aRoom.Name, aRoom.Description); tError != nil {
		return http.StatusInternalServerError, tError
	}

	return 0, nil
}

func GetRoomByID(aID string, aDb *sql.DB) (*models.Room, int, error) {
	tWhere := fmt.Sprintf("WHERE %s = %s", os.Getenv("MYSQL_ROOMS_TABLE_ID"), aID)
	tQuery := fmt.Sprintf("SELECT * FROM %s %s", os.Getenv("MYSQL_ROOMS_TABLE"), tWhere)

	tRoom := models.Room{}
	if tError := aDb.QueryRow(tQuery).Scan(&tRoom.RoomID, &tRoom.Name, &tRoom.Description, &tRoom.CreateAt, &tRoom.UpdateAt); tError != nil {
		return nil, http.StatusInternalServerError, tError
	}
	return &tRoom, 0, nil
}
