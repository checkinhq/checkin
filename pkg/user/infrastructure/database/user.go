package database

import (
	"database/sql"

	"github.com/cathalgarvey/fmtless"
	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/goph/clock"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

// Option sets a value in userRepository instance.
type Option func(r *userRepository)

// Clock returns a Option that sets the clock.
func Clock(c clock.Clock) Option {
	return func(r *userRepository) {
		r.clock = c
	}
}

type userRepository struct {
	db    *sql.DB
	clock clock.Clock
}

func NewUserRepository(db *sql.DB, opts ...Option) domain.UserRepository {
	repo := &userRepository{
		db: db,
	}

	for _, opt := range opts {
		opt(repo)
	}

	// Default clock
	if repo.clock == nil {
		repo.clock = clock.SystemClock
	}

	return repo
}

func (r *userRepository) FindByID(id int64) (*domain.User, error) {
	return r.findBy("id", id)
}

func (r *userRepository) FindByUID(uid ksuid.KSUID) (*domain.User, error) {
	return r.findBy("uid", uid)
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	return r.findBy("email", email)
}

func (r *userRepository) findBy(field string, value interface{}) (*domain.User, error) {
	u := new(domain.User)

	err := r.db.QueryRow(
		fmt.Sprintf(`
		SELECT
			id,
			uid,
			email,
			password,
			first_name,
			last_name,
			created_at,
			updated_at
		FROM
			users/*t*/
		WHERE
			%s = ?
		`, field),
		value,
	).Scan(
		&u.ID,
		&u.UID,
		&u.Email,
		&u.Password,
		&u.FirstName,
		&u.LastName,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "finding user failed")
	}

	return u, nil
}

func (r *userRepository) Store(u *domain.User) error {
	if u.ID == 0 {
		return r.create(u)
	}

	return r.update(u)
}

func (r *userRepository) create(u *domain.User) error {
	u.CreatedAt = r.clock.Now()
	u.UpdatedAt = r.clock.Now()

	result, err := r.db.Exec(
		`
		INSERT INTO
			users/*t*/
		SET
			uid = ?,
			email = ?,
			password = ?,
			first_name = ?,
			last_name = ?,
			created_at = ?,
			updated_at = ?
		`,
		u.UID,
		u.Email,
		u.Password,
		u.FirstName,
		u.LastName,
		u.CreatedAt,
		u.UpdatedAt,
	)

	if err != nil {
		return errors.Wrap(err, "cannot create user")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "cannot determine user ID")
	}

	u.ID = id

	return nil
}

func (r *userRepository) update(u *domain.User) error {
	u.UpdatedAt = r.clock.Now()

	_, err := r.db.Exec(
		`
		UPDATE
			users/*t*/
		SET
			email = ?,
			password = ?,
			first_name = ?,
			last_name = ?,
			updated_at = ?
		WHERE
			id = ?
		`,
		u.Email,
		u.Password,
		u.FirstName,
		u.LastName,
		u.UpdatedAt,
		u.ID,
	)

	if err != nil {
		return errors.Wrap(err, "cannot update user")
	}

	return nil
}
