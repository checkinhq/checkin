package app

import (
	"database/sql"

	proto "github.com/checkinhq/checkin/apis/checkin/user/v1alpha"
	"github.com/checkinhq/checkin/pkg/user/app"
	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/checkinhq/checkin/pkg/user/infrastructure"
	"github.com/checkinhq/checkin/pkg/user/infrastructure/database"
	"github.com/go-kit/kit/log"
	"github.com/goph/clock"
	"github.com/goph/emperror"
	"go.uber.org/dig"
	"go.uber.org/fx"
)

func User() fx.Option {
	return fx.Options(
		fx.Provide(
			NewUserRepository,
			NewUserAuthenticationService,
		),
		fx.Invoke(proto.RegisterAuthenticationServer),
	)
}

// UserAuthenticationServiceParams provides a set of dependencies for the service constructor.
type UserAuthenticationServiceParams struct {
	dig.In

	Repository   domain.UserRepository
	Logger       log.Logger       `optional:"true"`
	ErrorHandler emperror.Handler `optional:"true"`
}

// NewService returns a new service instance.
func NewUserAuthenticationService(params UserAuthenticationServiceParams) proto.AuthenticationServer {
	return app.NewAuthenticationService(
		infrastructure.NewAuthenticationService(params.Repository),
		app.Logger(params.Logger),
		app.ErrorHandler(params.ErrorHandler),
	)
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return database.NewUserRepository(db, clock.SystemClock)
}
