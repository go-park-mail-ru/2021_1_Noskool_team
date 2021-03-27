package postgresDB

import (
	"2021_1_Noskool_team/internal/app/profiles"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"database/sql"
	"fmt"
)

type ProfileRepository struct {
	con *sql.DB
}

func NewProfileRepository(con *sql.DB) profiles.Repository {
	return &ProfileRepository{
		con: con,
	}
}

// Create ...
func (r *ProfileRepository) Create(u *models.UserProfile) error {
	//const defaultAvatar = "/api/v1/data/img/default.png"
	if err := u.Validate(true); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.con.QueryRow("INSERT INTO Profiles"+
		"(email, nickname, first_name, second_name, encrypted_password, avatar)"+
		"VALUES ($1, $2, $3, $4, $5, $6)"+
		"RETURNING profiles_id;",
		u.Email,
		u.Login,
		u.Name,
		u.Surname,
		u.EncryptedPassword,
		u.Avatar).Scan(&u.ProfileID)
}

// Update ...
func (r *ProfileRepository) Update(u *models.UserProfile, withPassword bool) error {
	if withPassword {
		if err := u.Validate(true); err != nil {
			return err
		}
		if err := u.BeforeCreate(); err != nil {
			return err
		}
	}
	if err := u.Validate(false); err != nil {
		return err
	}

	return r.con.QueryRow("UPDATE Profiles "+
		"SET email = $1, nickname = $2, first_name = $3, second_name = $4, encrypted_password = $5 "+
		"WHERE profiles_id = $6 RETURNING profiles_id;",
		u.Email,
		u.Login,
		u.Name,
		u.Surname,
		u.EncryptedPassword,
		u.ProfileID).Scan(&u.ProfileID)
}

// FindByID ...
func (r *ProfileRepository) FindByID(id string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, first_name, second_name, encrypted_password, avatar FROM Profiles"+
		" WHERE profiles_id = %s;", id)
	if err := r.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.Name,
		&u.Surname,
		&u.EncryptedPassword,
		&u.Avatar); err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateAvatar ...
func (r *ProfileRepository) UpdateAvatar(userID string, newAvatar string) {
	r.con.QueryRow("UPDATE Profiles "+
		"SET avatar = $1 WHERE profiles_id = $2;",
		newAvatar, userID)
}

// FindByLogin ...
func (r *ProfileRepository) FindByLogin(nickname string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, first_name, second_name, encrypted_password FROM Profiles"+
		" WHERE nickname = '%s';", nickname)
	if err := r.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.Name,
		&u.Surname,
		&u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
