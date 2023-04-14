package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mizutanimeen/mandatory-declaration-of-intent/models"
)

func GetUserByID(aID string, aDb *sql.DB) (*models.User, int, error) {
	tWhere := fmt.Sprintf("WHERE %s = %s", os.Getenv("MYSQL_USERS_TABLE_ID"), aID)
	tQuery := fmt.Sprintf("SELECT * FROM %s %s", os.Getenv("MYSQL_USERS_TABLE"), tWhere)

	tUser := models.User{}
	if tError := aDb.QueryRow(tQuery).Scan(&tUser.UserID, &tUser.Name, &tUser.Text, &tUser.CreateAt, &tUser.UpdateAt); tError != nil {
		return nil, http.StatusInternalServerError, tError
	}
	return &tUser, 0, nil
}

func CreateUser(aUser *models.User, aDb *sql.DB) (int, error) {
	tColumns := fmt.Sprintf("(%s,%s)",
		os.Getenv("MYSQL_USERS_TABLE_NAME"), os.Getenv("MYSQL_USERS_TABLE_TEXT"))
	tQuery := fmt.Sprintf("INSERT INTO %s %s VALUES (?,?)",
		os.Getenv("MYSQL_USERS_TABLE"), tColumns)

	tStmt, tError := aDb.Prepare(tQuery)
	if tError != nil {
		return http.StatusInternalServerError, tError
	}
	defer tStmt.Close()

	if _, tError := tStmt.Exec(aUser.Name, aUser.Text); tError != nil {
		return http.StatusInternalServerError, tError
	}

	return 0, nil
}

func CreateGestUser(aGestUser *models.GestUser, aDb *sql.DB) (int, error) {
	tColumns := fmt.Sprintf("(%s,%s,%s)",
		os.Getenv("MYSQL_GEST_USERS_TABLE_NAME"), os.Getenv("MYSQL_GEST_USERS_TABLE_TEXT"), os.Getenv("MYSQL_ROOMS_TABLE_ID"))
	tQuery := fmt.Sprintf("INSERT INTO %s %s VALUES (?,?,?)",
		os.Getenv("MYSQL_GEST_USERS_TABLE"), tColumns)

	tStmt, tError := aDb.Prepare(tQuery)
	if tError != nil {
		return http.StatusInternalServerError, tError
	}
	defer tStmt.Close()

	if _, tError := tStmt.Exec(aGestUser.Name, aGestUser.Text, aGestUser.RoomID); tError != nil {
		return http.StatusInternalServerError, tError
	}

	return 0, nil
}
