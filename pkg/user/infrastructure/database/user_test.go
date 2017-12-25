// +build integration

package database_test

import (
	"testing"

	"database/sql"
	"time"

	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/checkinhq/checkin/pkg/user/infrastructure/database"
	_ "github.com/checkinhq/checkin/test/db"
	"github.com/goph/clock"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDbRepository_FindByID(t *testing.T) {
	db, err := sql.Open("txdb", "find by id")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())
	id := saveUser(t, db, expected)
	expected.ID = id

	actual, err := repo.FindByID(expected.ID)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestDbRepository_FindByID_NotFound(t *testing.T) {
	db, err := sql.Open("txdb", "user by id not found")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())

	_, err = repo.FindByID(expected.ID)
	assert.EqualError(t, err, "user not found: sql: no rows in result set")
}

func TestDbRepository_FindByUID(t *testing.T) {
	db, err := sql.Open("txdb", "find by uid")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())
	id := saveUser(t, db, expected)
	expected.ID = id

	actual, err := repo.FindByUID(expected.UID)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestDbRepository_FindByUID_NotFound(t *testing.T) {
	db, err := sql.Open("txdb", "user by uid not found")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())

	_, err = repo.FindByUID(expected.UID)
	assert.EqualError(t, err, "user not found: sql: no rows in result set")
}

func TestDbRepository_FindByEmail(t *testing.T) {
	db, err := sql.Open("txdb", "find by email")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())
	id := saveUser(t, db, expected)
	expected.ID = id

	actual, err := repo.FindByEmail(expected.Email)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestDbRepository_FindByEmail_NotFound(t *testing.T) {
	db, err := sql.Open("txdb", "user by email not found")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())

	_, err = repo.FindByEmail(expected.Email)
	assert.EqualError(t, err, "user not found: sql: no rows in result set")
}

func TestDbRepository_Create(t *testing.T) {
	db, err := sql.Open("txdb", "create user")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())

	err = repo.Store(expected)
	require.NoError(t, err)

	actual := getUser(t, db, expected.ID)
	assert.Equal(t, expected, actual)
}

func TestDbRepository_Update(t *testing.T) {
	db, err := sql.Open("txdb", "update user")
	require.NoError(t, err)
	defer db.Close()

	now := clock.StoppedAt(clock.SystemClock.Now().UTC().Truncate(time.Second))

	repo := database.NewUserRepository(db, now)

	expected := newUser(t, now.Now())
	id := saveUser(t, db, expected)
	expected.ID = id

	expected.Email = "john.new@example.com"

	err = repo.Store(expected)
	require.NoError(t, err)

	actual := getUser(t, db, expected.ID)
	assert.Equal(t, expected, actual)
}

func newUser(t *testing.T, now time.Time) *domain.User {
	uid, err := ksuid.NewRandom()
	require.NoError(t, err)

	u := &domain.User{
		UID: uid,

		Email:     "john@example.com",
		FirstName: "John",
		LastName:  "Doe",

		CreatedAt: now,
		UpdatedAt: now,
	}

	err = u.Password.Hash("password")
	require.NoError(t, err)

	return u
}

func saveUser(t *testing.T, db *sql.DB, user *domain.User) int64 {
	result, err := db.Exec(
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
		user.UID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.CreatedAt,
		user.UpdatedAt,
	)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	return id
}

func getUser(t *testing.T, db *sql.DB, id int64) *domain.User {
	u := new(domain.User)

	err := db.QueryRow(
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
	require.NoError(t, err)

	return u
}
