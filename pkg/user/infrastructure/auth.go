package infrastructure

import (
	"errors"
	"time"

	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/dgrijalva/jwt-go"
)

// ErrAuthenticationFailed is returned when a user is not found or the password doesn't match.
var ErrAuthenticationFailed = errors.New("authentication failed")

type AuthenticationService struct {
	userRepository domain.UserRepository
	secret         string
}

// NewAuthenticationService returns a new authentication service.
func NewAuthenticationService(userRepository domain.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
		secret:         "secret",
	}
}

func (s *AuthenticationService) Login(email string, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err == domain.ErrUserNotFound {
		return "", ErrAuthenticationFailed
	} else if err != nil {
		return "", err
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
