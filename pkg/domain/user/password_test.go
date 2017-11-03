package user_test

import (
	"testing"

	"github.com/checkinhq/checkin/pkg/domain/user"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/assert"
)

func TestPassword_Hash(t *testing.T) {
	rawPassword := "password"

	var password user.Password

	err := password.Hash(rawPassword)
	require.NoError(t, err)

	err = bcrypt.CompareHashAndPassword(password, []byte(rawPassword))
	require.NoError(t, err)
}

func TestPassword_Verify(t *testing.T) {
	rawPassword := "password"
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	require.NoError(t, err)

	var password user.Password = hash

	err = password.Verify(rawPassword)
	assert.NoError(t, err)
}

func TestPassword_String(t *testing.T) {
	rawPassword := "password"
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	require.NoError(t, err)

	var password user.Password = hash

	assert.Equal(t, string(hash), password.String())
}
