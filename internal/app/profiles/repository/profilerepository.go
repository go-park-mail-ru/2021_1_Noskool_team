package repository

import (
	"2021_1_Noskool_team/internal/app/profiles/models"
	"fmt"
)

// ProfileRepository ...
type ProfileRepository struct {
	db *Store
}

// Create ...
func (r *ProfileRepository) Create(u *models.UserProfile) error {
	if err := u.Validate(true); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.db.con.QueryRow("INSERT INTO Profiles"+
		"(email, nickname, encrypted_password)"+
		"VALUES ($1, $2, $3)"+
		"RETURNING profiles_id;",
		u.Email,
		u.Login,
		u.EncryptedPassword).Scan(&u.ProfileID)
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

	return r.db.con.QueryRow("UPDATE Profiles "+
		"SET email = $1, nickname = $2, encrypted_password = $3 "+
		"WHERE profiles_id = $4 RETURNING profiles_id;",
		u.Email,
		u.Login,
		u.EncryptedPassword,
		u.ProfileID).Scan(&u.ProfileID)
}

// FindByID ...
func (r *ProfileRepository) FindByID(id string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, encrypted_password FROM Profiles"+
		" WHERE profiles_id = %s;", id)
	if err := r.db.con.QueryRow(sqlReq).Scan(&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByLogin ...
func (r *ProfileRepository) FindByLogin(nickname string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, encrypted_password FROM Profiles"+
		" WHERE nickname = '%s';", nickname)
	if err := r.db.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
