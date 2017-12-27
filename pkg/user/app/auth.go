package app

import (
	api "github.com/checkinhq/checkin/apis/checkin/user/v1alpha"
	"github.com/checkinhq/checkin/pkg/user/infrastructure"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/emperror"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthenticationServiceOption sets options in the AuthenticationService.
type AuthenticationServiceOption func(s *AuthenticationService)

// Logger returns a AuthenticationServiceOption that sets the logger for the service.
func Logger(l log.Logger) AuthenticationServiceOption {
	return func(s *AuthenticationService) {
		s.logger = l
	}
}

// ErrorHandler returns a AuthenticationServiceOption that sets the error handler for the service.
func ErrorHandler(l emperror.Handler) AuthenticationServiceOption {
	return func(s *AuthenticationService) {
		s.errorHandler = l
	}
}

// AuthenticationService contains the main controller logic.
type AuthenticationService struct {
	service *infrastructure.AuthenticationService

	logger       log.Logger
	errorHandler emperror.Handler
}

// NewAuthenticationService creates a new service object.
func NewAuthenticationService(service *infrastructure.AuthenticationService, opts ...AuthenticationServiceOption) *AuthenticationService {
	s := new(AuthenticationService)

	s.service = service

	for _, opt := range opts {
		opt(s)
	}

	// Default logger
	if s.logger == nil {
		s.logger = log.NewNopLogger()
	}

	// Default error handler
	if s.errorHandler == nil {
		s.errorHandler = emperror.NewNopHandler()
	}

	return s
}

func (s *AuthenticationService) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	token, err := s.service.Login(request.GetEmail(), request.GetPassword())
	if err == infrastructure.ErrAuthenticationFailed {
		level.Debug(s.logger).Log(
			"msg", "authentication failed",
			"email", request.GetEmail(),
		)

		return nil, status.Errorf(codes.Unauthenticated, "cannot authenticate user")
	} else if err != nil {
		s.errorHandler.Handle(err)

		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return &api.LoginResponse{
		Token: token,
	}, nil
}
