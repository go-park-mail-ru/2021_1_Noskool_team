package repository

import (
	"2021_1_Noskool_team/configs"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	profileStore := New(configs.NewConfig())
	profileStore.Con = db

	defer db.Close()
	rows := sqlmock.NewRows([]string{
		"profiles_id", "email", "nickname", "encrypted_password", "avatar",
	}).AddRow("1", "some email", "some nickname", "4324322", "fdsfd")
	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)

	track, err := profileStore.User().FindByID("1")

	fmt.Println(track)
	assert.NoError(t, err)
	assert.NotNil(t, track)
}

func TestFindByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	profileStore := New(configs.NewConfig())
	profileStore.Con = db

	defer db.Close()
	rows := sqlmock.NewRows([]string{
		"profiles_id", "email", "nickname", "encrypted_password",
	}).AddRow("1", "some email", "some nickname", "4324322")
	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)

	track, err := profileStore.User().FindByLogin("nickname")

	fmt.Println(track)
	assert.NoError(t, err)
	assert.NotNil(t, track)
}
