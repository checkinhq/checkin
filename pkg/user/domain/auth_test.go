package domain_test

import (
	"testing"

	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/checkinhq/checkin/pkg/user/infrastructure/inmem"
	"github.com/dgrijalva/jwt-go"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationService_Login(t *testing.T) {
	repo := inmem.NewUserRepository()
	service := domain.NewAuthenticationService(repo)

	u := domain.NewUser()
	u.Password.Hash("password")
	u.Email = "john@doe.com"
	u.FirstName = "John"
	u.LastName = "Doe"

	repo.Store(u)

	token, err := service.Login("john@doe.com", "password")
	require.NoError(t, err)

	claims := new(jwt.StandardClaims)
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte("secret"), nil
	})
	require.NoError(t, err)

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	require.True(t, ok)

	assert.Equal(t, u.UID.String(), claims.Subject)
}
