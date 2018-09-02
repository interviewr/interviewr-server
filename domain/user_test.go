package domain_test

import (
	"interviewr-server/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectUser(t *testing.T) {
	org := domain.User{
		ID:        "id",
		Login:     "login",
		Name:      "name",
		Email:     "email@text.com",
		Location:  "awesome location",
		Bio:       "some bio",
		AvatarURL: "http://avatar.com",
	}
	error := org.Validate()
	assert.Nil(t, error)
}

var invalidUsers = []domain.User{
	{ID: "", Name: "name", Login: "Login", Email: "email@text.com",
		Bio: "bio", Location: "location", AvatarURL: "http://avatar.com"},
	{ID: "ID", Name: "", Login: "Login", Email: "email@text.com",
		Bio: "bio", Location: "location", AvatarURL: "http://avatar.com"},
	{ID: "ID", Name: "name", Login: "Login", Email: "",
		Bio: "bio", Location: "location", AvatarURL: "http://avatar.com"},
	{ID: "ID", Name: "name", Login: "Login", Email: "non-email",
		Bio: "bio", Location: "location", AvatarURL: "http://avatar.com"},
	{ID: "ID", Name: "name", Login: "Login", Email: "non@valid",
		Bio: "bio", Location: "location", AvatarURL: "xss://avatar.com"},
}

func TestInvalidUsers(t *testing.T) {
	for _, user := range invalidUsers {
		error := user.Validate()
		assert.Error(t, error)
	}
}
