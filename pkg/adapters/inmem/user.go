package inmem

import (
	"sync"

	"github.com/checkinhq/checkin/pkg/adapters/inmem/internal"
	"github.com/checkinhq/checkin/pkg/domain/user"
	"github.com/goph/clock"
	"github.com/hashicorp/go-memdb"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

var userSchema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"users": {
			Name: "users",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &internal.IntFieldIndex{Field: "ID"},
				},
				"uid": {
					Name:    "uid",
					Unique:  true,
					Indexer: &internal.KSUIDFieldIndex{Field: "UID"},
				},
				"email": {
					Name:    "email",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Email"},
				},
			},
		},
	},
}

// NewUserDB creates a new, in-memory user database.
func NewUserDB() (*memdb.MemDB, error) {
	return memdb.NewMemDB(userSchema)
}

type userRepository struct {
	db             *memdb.MemDB
	lastInsertedId int64

	clock clock.Clock

	mu sync.Mutex
}

type UserRepositoryOption func(repo *userRepository)

func UserDB(db *memdb.MemDB) UserRepositoryOption {
	return func(repo *userRepository) {
		repo.db = db
	}
}

func UserClock(c clock.Clock) UserRepositoryOption {
	return func(repo *userRepository) {
		repo.clock = c
	}
}

func NewUserRepository(opts ...UserRepositoryOption) user.Repository {
	repo := new(userRepository)

	for _, opt := range opts {
		opt(repo)
	}

	// Default DB
	if repo.db == nil {
		db, err := NewUserDB()
		if err != nil {
			panic(err)
		}

		repo.db = db
	}

	// Default clock
	if repo.clock == nil {
		repo.clock = clock.SystemClock
	}

	return repo
}

func (r *userRepository) FindByID(id int64) (*user.User, error) {
	// Create read-only transaction
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", "id", id)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	u, ok := raw.(*user.User)
	if !ok {
		return nil, errors.New("invalid data in the user record")
	}

	return u, nil
}

func (r *userRepository) FindByUID(uid ksuid.KSUID) (*user.User, error) {
	// Create read-only transaction
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", "uid", uid)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	u, ok := raw.(*user.User)
	if !ok {
		return nil, errors.New("invalid data in the user record")
	}

	return u, nil
}

func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	// Create read-only transaction
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", "email", email)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	u, ok := raw.(*user.User)
	if !ok {
		return nil, errors.New("invalid data in the user record")
	}

	return u, nil
}

func (r *userRepository) Store(u *user.User) error {
	txn := r.db.Txn(true)

	// Decide if this is a new user
	if u.ID == 0 {
		r.mu.Lock()
		r.lastInsertedId++
		u.ID = r.lastInsertedId
		r.mu.Unlock()

		u.CreatedAt = r.clock.Now()
	}

	u.UpdatedAt = r.clock.Now()

	if err := txn.Insert("users", u); err != nil {
		return errors.Wrap(err, "cannot insert user")
	}

	txn.Commit()

	return nil
}