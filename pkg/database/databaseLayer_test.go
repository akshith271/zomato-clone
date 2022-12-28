package database

import (
	"database/sql"
	"fmt"
	model "mock-grpc/models"
	"mock-grpc/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func SetupDBClient() (*gorm.DB, sqlmock.Sqlmock, error) {
	var (
		Db *sql.DB
	)
	Db, mock, _ := sqlmock.New()
	db, err := gorm.Open("postgres", Db)
	return db, mock, err
}
func TestCreateUser(t *testing.T) {
	db, mock, err := SetupDBClient()
	if err != nil {
		t.Fatalf("failed to create mock db client: %v", err)
	}
	defer db.Close()
	mockClient := DBClient{
		Db: db,
	}
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" (.+)`).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1),
	).WillReturnError(err)
	mock.ExpectCommit()
	createUserError := mockClient.CreateUser(model.User{})
	utils.CheckError(createUserError)
	if err != nil {
		fmt.Println("failed to create user")
	}
}
