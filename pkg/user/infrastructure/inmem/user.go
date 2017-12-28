package inmem

import (
	"sync"

	"github.com/checkinhq/checkin/pkg/internal/inmem"
	"github.com/checkinhq/checkin/pkg/user/domain"
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
					Indexer: &inmem.IntFieldIndex{Field: "ID"},
				},
				"uid": {
					Name:    "uid",
					Unique:  true,
					Indexer: &inmem.KSUIDFieldIndex{Field: "UID"},
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

type Option func(repo *userRepository)

func Database(db *memdb.MemDB) Option {
	return func(repo *userRepository) {
		repo.db = db
	}
}

func Clock(c clock.Clock) Option {
	return func(repo *userRepository) {
		repo.clock = c
	}
}

func NewUserRepository(opts ...Option) domain.UserRepository {
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
	// Create read-only transaction
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", field, value)
	if err != nil || raw == nil {
		return nil, domain.ErrUserNotFound
	}

	u, ok := raw.(*domain.User)
	if !ok {
		return nil, errors.New("invalid data in the user record")
	}

	return u, nil
}

func (r *userRepository) Store(u *domain.User) error {
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
