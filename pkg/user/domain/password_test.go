package domain_test

import (
	"testing"

	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword_Hash(t *testing.T) {
	rawPassword := "password"

	var password domain.Password

	err := password.Hash(rawPassword)
	require.NoError(t, err)

	err = bcrypt.CompareHashAndPassword(password, []byte(rawPassword))
	require.NoError(t, err)
}

func TestPassword_Verify(t *testing.T) {
	rawPassword := "password"
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	require.NoError(t, err)

	var password domain.Password = hash

	err = password.Verify(rawPassword)
	assert.NoError(t, err)
}

func TestPassword_String(t *testing.T) {
	rawPassword := "password"
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	require.NoError(t, err)

	var password domain.Password = hash

	assert.Equal(t, string(hash), password.String())
}
