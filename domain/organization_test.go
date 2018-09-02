package domain_test

import (
	"interviewr-server/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectOrganization(t *testing.T) {
	org := domain.Organization{
		ID:          "id",
		Name:        "name",
		Email:       "email@text.com",
		Description: "some description",
		Location:    "awesome location",
	}
	error := org.Validate()
	assert.Nil(t, error)
}

var invalidOrgs = []domain.Organization{
	{ID: "", Name: "name", Email: "email@text.com", Description: "desc", Location: "location"},
	{ID: "ID", Name: "", Email: "email@text.com", Description: "desc", Location: "location"},
	{ID: "ID", Name: "name", Email: "", Description: "desc", Location: "location"},
	{ID: "ID", Name: "name", Email: "non-email", Description: "desc", Location: "location"},
	{ID: "ID", Name: "name", Email: "non@valid", Description: "desc", Location: "location"},
}

func TestInvalidOrganizations(t *testing.T) {
	for _, org := range invalidOrgs {
		error := org.Validate()
		assert.Error(t, error)
	}
}
