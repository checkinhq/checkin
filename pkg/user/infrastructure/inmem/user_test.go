package inmem_test

import (
	"testing"

	"time"

	"github.com/checkinhq/checkin/pkg/user/domain"
	"github.com/checkinhq/checkin/pkg/user/infrastructure/inmem"
	"github.com/hashicorp/go-memdb"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_FindByID(t *testing.T) {
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	u := newUser(t)

	saveUser(t, db, u)

	u2, err := repo.FindByID(u.ID)
	require.NoError(t, err)
	assert.Equal(t, u, u2)
}

func TestUserRepository_FindByID_NotFound(t *testing.T) {
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	_, err = repo.FindByID(1)
	assert.Equal(t, err, domain.ErrUserNotFound)
}

func TestUserRepository_FindByUID(t *testing.T) {
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	u := newUser(t)

	saveUser(t, db, u)

	u2, err := repo.FindByUID(u.UID)
	require.NoError(t, err)
	assert.Equal(t, u, u2)
}

func TestUserRepository_FindByUID_NotFound(t *testing.T) {
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	_, err = repo.FindByUID(ksuid.New())
	assert.Equal(t, err, domain.ErrUserNotFound)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	u := newUser(t)

	saveUser(t, db, u)

	u2, err := repo.FindByEmail(u.Email)
	require.NoError(t, err)
	assert.Equal(t, u, u2)
}

func TestUserRepository_FindByEmail_NotFound(t *testing.T) {
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	_, err = repo.FindByEmail("john.doe@example.com")
	assert.Equal(t, err, domain.ErrUserNotFound)
}

func TestUserRepository_Create(t *testing.T) {
	var db *memdb.MemDB
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	u := newUser(t)
	u.ID = 0

	err = repo.Store(u)
	require.NoError(t, err)

	txn := db.Txn(false)
	defer txn.Abort()

	u2, err := txn.First("users", "id", u.ID)
	require.NoError(t, err)
	assert.Equal(t, u, u2)
}

func TestUserRepository_Update(t *testing.T) {
	var db *memdb.MemDB
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	u := newUser(t)

	saveUser(t, db, u)

	u.Email = "john.new@doe.com"

	err = repo.Store(u)
	require.NoError(t, err)

	txn := db.Txn(false)
	defer txn.Abort()

	u2, err := txn.First("users", "id", u.ID)
	require.NoError(t, err)
	assert.Equal(t, u, u2)
}

func TestUserRepository_CreateMore(t *testing.T) {
	var db *memdb.MemDB
	db, err := inmem.NewUserDB()
	require.NoError(t, err)

	repo := inmem.NewUserRepository(inmem.UserDB(db))

	u := newUser(t)
	u.ID = 0
	u2 := newUser(t)
	u2.ID = 0

	err = repo.Store(u)
	require.NoError(t, err)

	err = repo.Store(u2)
	require.NoError(t, err)

	txn := db.Txn(false)
	defer txn.Abort()

	uactual, err := txn.First("users", "id", u.ID)
	require.NoError(t, err)
	assert.Equal(t, u, uactual)

	u2actual, err := txn.First("users", "id", u2.ID)
	require.NoError(t, err)
	assert.Equal(t, u2, u2actual)
}

func newUser(t *testing.T) *domain.User {
	uid, err := ksuid.NewRandom()
	require.NoError(t, err)

	u := &domain.User{
		ID:  1,
		UID: uid,

		Email:     "john@doe.com",
		FirstName: "John",
		LastName:  "Doe",

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.Password.Hash("password")
	require.NoError(t, err)

	return u
}

func saveUser(t *testing.T, db *memdb.MemDB, u *domain.User) {
	txn := db.Txn(true)

	err := txn.Insert("users", u)

	require.NoError(t, err)

	txn.Commit()
}
