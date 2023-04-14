package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mizutanimeen/mandatory-declaration-of-intent/models"
)

func GetAllGestMembers(aID string, aDb *sql.DB) ([]*models.GestUser, int, error) {
	tWhere := fmt.Sprintf("WHERE %s = %s", os.Getenv("MYSQL_ROOMS_TABLE_ID"), aID)
	tQuery := fmt.Sprintf("SELECT * FROM %s %s", os.Getenv("MYSQL_GEST_USERS_TABLE"), tWhere)

	tRows, tError := aDb.Query(tQuery)
	if tError != nil {
		return nil, http.StatusInternalServerError, tError
	}
	defer tRows.Close()

	tGestUsers := []*models.GestUser{}
	for tRows.Next() {
		tGestUser := models.GestUser{}
		if tError := tRows.Scan(&tGestUser.GestUserID, &tGestUser.Name, &tGestUser.Text, &tGestUser.RoomID, &tGestUser.CreateAt, &tGestUser.UpdateAt); tError != nil {
			return nil, http.StatusInternalServerError, tError
		}
		tGestUsers = append(tGestUsers, &tGestUser)
	}

	return tGestUsers, 0, nil
}
