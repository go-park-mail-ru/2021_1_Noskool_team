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
	//const defaultAvatar = "/api/v1/data/img/default.png"
	if err := u.Validate(true); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.db.Con.QueryRow("INSERT INTO Profiles"+
		"(email, nickname, encrypted_password, avatar)"+
		"VALUES ($1, $2, $3, $4)"+
		"RETURNING profiles_id;",
		u.Email,
		u.Login,
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

	return r.db.Con.QueryRow("UPDATE Profiles "+
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
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, encrypted_password, avatar FROM Profiles"+
		" WHERE profiles_id = %s;", id)
	if err := r.db.Con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.EncryptedPassword,
		&u.Avatar); err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateAvatar ...
func (r *ProfileRepository) UpdateAvatar(userID string, newAvatar string) {
	r.db.Con.QueryRow("UPDATE Profiles "+
		"SET avatar = $1 WHERE profiles_id = $2;",
		newAvatar, userID)
}

// FindByLogin ...
func (r *ProfileRepository) FindByLogin(nickname string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, encrypted_password FROM Profiles"+
		" WHERE nickname = '%s';", nickname)
	if err := r.db.Con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
