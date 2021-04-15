package postgresDB

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	profRep := NewProfileRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"profiles_id", "email", "nickname", "encrypted_password", "avatar", "first_name",
		"second_name", "favorite_genre",
	}).AddRow("1", "some email", "some nickname", "4324322", "fdsfd",
		"alex", "alex", "{rok, pop}")
	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)

	p, err := profRep.FindByID("1")

	fmt.Println(p)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}

func TestFindByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	profRep := NewProfileRepository(db)

	defer db.Close()
	rows := sqlmock.NewRows([]string{
		"profiles_id", "email", "nickname", "encrypted_password", "avatar", "first_name",
		"second_name", "favorite_genre",
	}).AddRow("1", "some email", "some nickname", "4324322", "fdsfd",
		"alex", "alex", "{rok, pop}")
	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)

	p, err := profRep.FindByLogin("nickname")

	fmt.Println(p)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}
