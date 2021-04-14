package postgresDB

import (
	"2021_1_Noskool_team/internal/app/profiles/models"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	expectedUsers = []models.UserProfile{
		{
			ProfileID:         1,
			Email:             "test1@gmail.com",
			Login:             "test1",
			Name:              "Name1",
			Surname:           "Surname1",
			Password:          "Password1",
			EncryptedPassword: "EncryptedPassword1",
			Avatar:            "Avatar1",
			FavoriteGenre:     []string{"pop"},
		},
		{
			ProfileID:         2,
			Email:             "test2@gmail.com",
			Login:             "test2",
			Name:              "Name2",
			Surname:           "Surname2",
			Password:          "Password2",
			EncryptedPassword: "EncryptedPassword2",
			Avatar:            "Avatar2",
			FavoriteGenre:     []string{"rock"},
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

// func TestCreate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("cant create mock '%s'", err)
// 	}
// 	profilesRep := NewProfileRepository(db)
// 	defer db.Close()

// 	createUser := models.UserProfile{
// 		Email:         "test222@gmail.com",
// 		Login:         "test222",
// 		Name:          "Name222",
// 		Surname:       "Surname222",
// 		Password:      "Password222",
// 		Avatar:        "Avatar222",
// 		FavoriteGenre: []string{"rock"},
// 	}

// 	createUserExpected := models.UserProfile{
// 		Email:         "test222@gmail.com",
// 		Login:         "test222",
// 		Name:          "Name222",
// 		Surname:       "Surname222",
// 		Password:      "Password222",
// 		Avatar:        "Avatar222",
// 		FavoriteGenre: []string{"rock"},
// 	}

// 	if err := createUserExpected.Validate(true); err != nil {
// 		t.Error("validate")
// 	}
// 	if err := createUserExpected.BeforeCreate(); err != nil {
// 		t.Error("BeforeCreate")
// 	}

// 	query := regexp.QuoteMeta(`INSERT INTO Profiles(email, nickname, first_name, second_name, encrypted_password, avatar, favorite_genre)VALUES ($1, $2, $3, $4, $5, $6, $7);`)

// 	mock.ExpectQuery(query).WithArgs(&createUser)

// 	err = profilesRep.Create(&createUser)

// 	assert.NoError(t, err)
// }
