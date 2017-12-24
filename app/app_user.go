package app

import (
	"database/sql"

	user_api "github.com/checkinhq/checkin/apis/checkin/user/v1alpha"
	"github.com/checkinhq/checkin/pkg/adapters/database"
	user_app "github.com/checkinhq/checkin/pkg/app/user"
	"github.com/checkinhq/checkin/pkg/domain/user"
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
		fx.Invoke(user_api.RegisterAuthenticationServer),
	)
}

// UserAuthenticationServiceParams provides a set of dependencies for the service constructor.
type UserAuthenticationServiceParams struct {
	dig.In

	Repository   user.Repository
	Logger       log.Logger       `optional:"true"`
	ErrorHandler emperror.Handler `optional:"true"`
}

// NewService returns a new service instance.
func NewUserAuthenticationService(params UserAuthenticationServiceParams) user_api.AuthenticationServer {
	return user_app.NewAuthenticationService(
		user.NewAuthenticationService(params.Repository),
		user_app.Logger(params.Logger),
		user_app.ErrorHandler(params.ErrorHandler),
	)
}

func NewUserRepository(db *sql.DB) user.Repository {
	return database.NewUserRepository(db, clock.SystemClock)
}
