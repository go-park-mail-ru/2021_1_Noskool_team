package postgresDB

import (
	"2021_1_Noskool_team/internal/app/profiles/models"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	otherUsers = []*models.OtherUser{
		{
			UserID:      1,
			Nickname:    "some nickname",
			Photo:       "photo",
			ISubscribed: false,
		},
		{
			UserID:      2,
			Nickname:    "some nickname",
			Photo:       "photo",
			ISubscribed: false,
		},
	}
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

func TestGetSubscriptions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewProfileRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"profiles_id", "nickname", "avatar",
	})
	for _, row := range otherUsers {
		rows.AddRow(row.UserID, row.Nickname, row.Photo)
	}

	query := "select profiles_id, nickname, avatar from profiles"
	mock.ExpectQuery(query).WillReturnRows(rows)
	track, err := tracRep.GetSubscriptions(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(otherUsers, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetSubscribers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewProfileRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"profiles_id", "nickname", "avatar",
	})
	for _, row := range otherUsers {
		rows.AddRow(row.UserID, row.Nickname, row.Photo)
	}

	query := "select profiles_id, nickname, avatar from profiles"
	mock.ExpectQuery(query).WillReturnRows(rows)
	track, err := tracRep.GetSubscribers(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(otherUsers, track) {
		t.Fatalf("Not equal")
	}
}

func TestSearchTracks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewProfileRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"profiles_id", "nickname", "avatar",
	})
	for _, row := range otherUsers {
		rows.AddRow(row.UserID, row.Nickname, row.Photo)
	}

	query := "SELECT profiles_id, nickname, avatar FROM profiles"
	mock.ExpectQuery(query).WillReturnRows(rows)
	track, err := tracRep.SearchTracks("search query")

	assert.NoError(t, err)
	if !reflect.DeepEqual(otherUsers, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetOtherUserPage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	profRep := NewProfileRepository(db)

	defer db.Close()
	rows := sqlmock.NewRows([]string{
		"profiles_id", "nickname", "avatar",
	}).AddRow(otherUsers[0].UserID, otherUsers[0].Nickname, otherUsers[0].Photo)
	query := "select profiles_id, nickname, avatar from profiles"

	mock.ExpectQuery(query).WillReturnRows(rows)

	p, err := profRep.GetOtherUserPage(1)

	fmt.Println(p)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}
