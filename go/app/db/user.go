package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mizutanimeen/mandatory-declaration-of-intent/models"
)

func GetUserByID(aID string, aDb *sql.DB) (*models.User, int, error) {
	return &models.User{}, 0, nil
}
