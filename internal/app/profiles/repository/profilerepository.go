package repository

import "2021_1_Noskool_team/internal/app/profiles/models"

// ProfileRepository ...
type ProfileRepository struct {
	db *Store
}

// Create ...
func (r *ProfileRepository) Create(u *models.UserProfile) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.db.con.QueryRow("INSERT INTO profiles"+
		"(email, nickname, encrypted_password)"+
		"VALUES ($1, $2, $3)"+
		"RETURNING profiles_id",
		u.Email,
		u.Login,
		u.Encrypted_password).Scan(&u.ProfileID)
}

// FindByEmail ...
func (r *ProfileRepository) FindByEmail(email string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	if err := r.db.con.QueryRow("SELECT email, nickname, encrypted_password FROM profiles"+
		"WHERE email = $1", email).Scan(
		&u.Email,
		&u.Login,
		&u.Encrypted_password); err != nil {
		return nil, err
	}
	return u, nil
}
