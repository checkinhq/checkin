package database

import (
	"database/sql"

	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/goph/clock"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

type userRepository struct {
	db    *sql.DB
	clock clock.Clock
}

func NewUserRepository(db *sql.DB, c clock.Clock) domain.UserRepository {
	repo := &userRepository{
		db:    db,
		clock: c,
	}

	// Default clock
	if c == nil {
		c = clock.SystemClock
	}

	return repo
}

func (r *userRepository) FindByID(id int64) (*domain.User, error) {
	u := new(domain.User)

	err := r.db.QueryRow(
		`
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
			id = ?
		`,
		id,
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
		return nil, errors.Wrap(err, "user not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "searching for user failed")
	}

	return u, nil
}

func (r *userRepository) FindByUID(uid ksuid.KSUID) (*domain.User, error) {
	u := new(domain.User)

	err := r.db.QueryRow(
		`
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
			uid = ?
		`,
		uid,
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
		return nil, errors.Wrap(err, "user not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "searching for user failed")
	}

	return u, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	u := new(domain.User)

	err := r.db.QueryRow(
		`
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
			email = ?
		`,
		email,
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
		return nil, errors.Wrap(err, "user not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "searching for user failed")
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
