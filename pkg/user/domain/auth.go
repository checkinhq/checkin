package domain

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ErrAuthenticationFailed is returned when a user is not found or the password doesn't match.
var ErrAuthenticationFailed = errors.New("authentication failed")

// AuthenticationService authenticates a user.
type AuthenticationService interface {
	// Login authenticates a user with an email-password pair.
	Login(email string, password string) (string, error)
}

type authenticationService struct {
	userRepository UserRepository
	secret         string
}

// NewAuthenticationService returns a new authentication service.
func NewAuthenticationService(userRepository UserRepository) AuthenticationService {
	return &authenticationService{
		userRepository: userRepository,
		secret:         "secret",
	}
}

func (s *authenticationService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", ErrAuthenticationFailed
	}

	err = user.Password.Verify(password)
	if err != nil {
		return "", ErrAuthenticationFailed
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		Subject:   user.UID.String(),
		ExpiresAt: time.Now().Add(1500 * time.Second).Unix(),
		Issuer:    "checkin-login",
		Audience:  "general",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secret))
}
