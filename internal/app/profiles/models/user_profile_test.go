package models

import (
	"fmt"
	"testing"

	"github.com/microcosm-cc/bluemonday"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name            string
		userCridentials UserProfile
		withPassword    bool
		want            error
	}{
		{
			name:            "normal with password",
			userCridentials: UserProfile{Email: "test123@mail.ru", Login: "test123", Password: "test123pas"},
			withPassword:    true,
			want:            fmt.Errorf("not error"),
		},
		{
			name:            "bad email (not @)",
			userCridentials: UserProfile{Email: "test123mail.ru", Login: "test123", Password: "test123pas"},
			withPassword:    true,
			want:            nil,
		},
		{
			name:            "bad login (too short)",
			userCridentials: UserProfile{Email: "test123@mail.ru", Login: "t", Password: "test123pas"},
			withPassword:    true,
			want:            nil,
		},
		{
			name:            "bad password (too short)",
			userCridentials: UserProfile{Email: "test123@mail.ru", Login: "test123", Password: "123"},
			withPassword:    true,
			want:            nil,
		},
		{
			name: "bad password (too long)",
			userCridentials: UserProfile{Email: "test123@mail.ru",
				Login:    "test123",
				Password: "171711askdfjlaksdjkfjalksdjf9384398439848578394dfnvmldfdgrd56677778888899KJHKJHL"},
			withPassword: true,
			want:         nil,
		},
		{
			name:            "normal without password",
			userCridentials: UserProfile{Email: "alah_babah@mail.ru", Login: "testTEST"},
			withPassword:    false,
			want:            fmt.Errorf("not error"),
		},
		{
			name:            "bad email without password",
			userCridentials: UserProfile{Email: "alah_babah", Login: "testTEST"},
			withPassword:    false,
			want:            nil,
		},
		{
			name:            "bad login without password",
			userCridentials: UserProfile{Email: "alah_babah@mail.ru", Login: "1"},
			withPassword:    false,
			want:            nil,
		},
	}

	for _, testCase := range testCases {
		user := testCase.userCridentials
		resErr := user.ValidateForCreate()

		assert.NotEqual(t, testCase.want, resErr)
	}
}

func TestRequestLoginSanitize(t *testing.T) {
	req := &RequestLogin{Password: "<a onblur=\"alert(secret)\" href=\"http://www.google.com\">Google</a>"}
	sanitize := bluemonday.UGCPolicy()

	req.Sanitize(sanitize)
	assert.Equal(t, req.Password, "<a href=\"http://www.google.com\" rel=\"nofollow\">Google</a>")
}

func TestProfileRequestSanitize(t *testing.T) {
	xssString := "<a onblur=\"alert(secret)\" href=\"http://www.google.com\">Google</a>"
	req := &ProfileRequest{
		Email:         xssString,
		Password:      "",
		Nickname:      "",
		Name:          "",
		Surname:       "",
		FavoriteGenre: []string{xssString},
	}
	sanitize := bluemonday.UGCPolicy()
	req.Sanitize(sanitize)

	assert.Equal(t, req.Email, "<a href=\"http://www.google.com\" rel=\"nofollow\">Google</a>")
	assert.Equal(t, req.FavoriteGenre[0], "<a href=\"http://www.google.com\" rel=\"nofollow\">Google</a>")
}
