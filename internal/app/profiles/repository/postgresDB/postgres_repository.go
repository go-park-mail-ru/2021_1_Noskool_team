package postgresDB

import (
	"2021_1_Noskool_team/internal/app/profiles"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
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
	if err := u.ValidateForCreate(); err != nil {
		validationErr, _ := json.Marshal(err)
		return fmt.Errorf(string(validationErr))
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	_, err := r.con.Exec("INSERT INTO Profiles"+
		"(email, nickname, encrypted_password, avatar)"+
		"VALUES ($1, $2, $3, $4);",
		u.Email,
		u.Login,
		u.EncryptedPassword,
		u.Avatar)

	pgErr, ok := err.(*pq.Error)
	if ok {
		err = formattingDBerr(pgErr)
	}
	return err
}

// Update ...
func (r *ProfileRepository) Update(u *models.UserProfile) error {
	if err := u.ValidationForUpdate(); err != nil {
		fmt.Println(">>>", err)
		validationErr, _ := json.Marshal(err)
		return fmt.Errorf(string(validationErr))
	}

	_, err := r.con.Exec("UPDATE Profiles "+
		"SET email = $1, nickname = $2, first_name = $3, second_name = $4, encrypted_password = $5, favorite_genre = $6 "+
		"WHERE profiles_id = $7;",
		u.Email,
		u.Login,
		u.Name,
		u.Surname,
		u.EncryptedPassword,
		pq.Array(u.FavoriteGenre),
		u.ProfileID)

	pgErr, ok := err.(*pq.Error)
	if ok {
		err = formattingDBerr(pgErr)
	}
	return err
}

// FindByID ...
func (r *ProfileRepository) FindByID(id string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, first_name, second_name,"+
		"encrypted_password, avatar, favorite_genre  FROM Profiles"+
		" WHERE profiles_id = %s;", id)
	if err := r.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.Name,
		&u.Surname,
		&u.EncryptedPassword,
		&u.Avatar,
		pq.Array(&u.FavoriteGenre)); err != nil {
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
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, first_name, second_name, encrypted_password, avatar, favorite_genre  FROM Profiles"+
		" WHERE nickname = '%s';", nickname)
	if err := r.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.Name,
		&u.Surname,
		&u.EncryptedPassword,
		&u.Avatar,
		pq.Array(&u.FavoriteGenre)); err != nil {
		return nil, err
	}
	return u, nil
}

// Create ...
func (r *ProfileRepository) UpdatePassword(id int, newPass string) error {
	if err := models.ValidateForChangePass(newPass); err != nil {
		return err
	}
	encNewPass, err := models.BeforeUpdatePass(newPass)
	if err != nil {
		return err
	}

	_, err = r.con.Exec("UPDATE Profiles SET encrypted_password = $1 WHERE profiles_id = $2;", encNewPass, id)

	if err != nil {
		return err
	}
	return nil
}

func formattingDBerr(err *pq.Error) error {
	fmt.Println(err)
	var formatedErr error
	switch {
	case err.Code == "23505" && err.Constraint == "profiles_email_key":
		formatedErr = models.ErrConstraintViolationEmail
	case err.Code == "23505" && err.Constraint == "profiles_nickname_key":
		formatedErr = models.ErrConstraintViolationNickname
	default:
		formatedErr = models.ErrDefaultDB
	}
	return formatedErr
}

func (r *ProfileRepository) SubscribeMeToSomebody(myID, otherUserID int) error {
	query := `insert into subscriptions (who, on_whom) VALUES ($1, $2)`

	_, err := r.con.Exec(query, myID, otherUserID)
	return err
}

func (r *ProfileRepository) UnsubscribeMeToSomebody(myID, otherUserID int) error {
	query := `delete from subscriptions where who = $1 and on_whom = $2`

	_, err := r.con.Exec(query, myID, otherUserID)
	return err
}

func (r *ProfileRepository) CheckIsMySubscriber(myID, otherUserID int) bool {
	query := `select count(*) from subscriptions
			where who = $1 and on_whom = $2`

	amount := 0
	err := r.con.QueryRow(query, myID, otherUserID).Scan(&amount)
	if err != nil || amount != 1 {
		return false
	}
	fmt.Println(amount)
	return true
}

func (r *ProfileRepository) GetOtherUserPage(otherUserID int) (*models.OtherUser, error) {
	query := `select profiles_id, nickname, avatar from profiles
			where profiles_id = $1`
	otherUser := &models.OtherUser{}

	err := r.con.QueryRow(query, otherUserID).Scan(&otherUser.UserID, &otherUser.Nickname,
		&otherUser.Photo,
	)
	return otherUser, err
}

func (r *ProfileRepository) GetSubscriptions(userID int) ([]*models.OtherUser, error) {
	query := `select profiles_id, nickname, avatar from profiles
    		left join subscriptions s on profiles.profiles_id = s.on_whom
			where s.who = $1`

	rows, err := r.con.Query(query, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	subscriptions := make([]*models.OtherUser, 0)

	for rows.Next() {
		otherUser := &models.OtherUser{}
		err = rows.Scan(&otherUser.UserID, &otherUser.Nickname, &otherUser.Photo)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, otherUser)
	}
	return subscriptions, err
}

func (r *ProfileRepository) GetSubscribers(userID int) ([]*models.OtherUser, error) {
	query := `select profiles_id, nickname, avatar from profiles
			left join subscriptions s on profiles.profiles_id = s.who
			where s.on_whom = $1`

	rows, err := r.con.Query(query, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	subscriptions := make([]*models.OtherUser, 0)

	for rows.Next() {
		otherUser := &models.OtherUser{}
		err = rows.Scan(&otherUser.UserID, &otherUser.Nickname, &otherUser.Photo)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, otherUser)
	}
	return subscriptions, err
}

func (r *ProfileRepository) SearchTracks(searchQuery string) ([]*models.OtherUser, error) {
	query := `SELECT profiles_id, nickname, avatar FROM profiles
			WHERE nickname LIKE '%' || $1 || '%'`

	rows, err := r.con.Query(query, searchQuery)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	subscriptions := make([]*models.OtherUser, 0)

	for rows.Next() {
		otherUser := &models.OtherUser{}
		err = rows.Scan(&otherUser.UserID, &otherUser.Nickname, &otherUser.Photo)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, otherUser)
	}
	return subscriptions, err
}
