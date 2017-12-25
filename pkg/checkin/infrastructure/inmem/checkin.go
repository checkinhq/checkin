package inmem

import (
	"sync"

	"github.com/checkinhq/checkin/pkg/checkin/domain"
	"github.com/checkinhq/checkin/pkg/internal/inmem"
	"github.com/goph/clock"
	"github.com/hashicorp/go-memdb"
	"github.com/pkg/errors"
)

var checkinSchema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"checkins": {
			Name: "checkins",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &inmem.IntFieldIndex{Field: "ID"},
				},
				"checkin": {
					Name:   "checkin",
					Unique: true,
					Indexer: &memdb.CompoundIndex{
						Indexes: []memdb.Indexer{
							&inmem.IntFieldIndex{Field: "ID"},
							&memdb.StringFieldIndex{Field: "Date"},
						},
					},
				},
			},
		},
	},
}

// NewCheckinDB creates a new, in-memory checkin database.
func NewCheckinDB() (*memdb.MemDB, error) {
	return memdb.NewMemDB(checkinSchema)
}

type checkinRepository struct {
	db             *memdb.MemDB
	lastInsertedId int64

	clock clock.Clock

	mu sync.Mutex
}

type CheckinRepositoryOption func(repo *checkinRepository)

func CheckinDB(db *memdb.MemDB) CheckinRepositoryOption {
	return func(repo *checkinRepository) {
		repo.db = db
	}
}

func CheckinClock(c clock.Clock) CheckinRepositoryOption {
	return func(repo *checkinRepository) {
		repo.clock = c
	}
}

func NewCheckinRepository(opts ...CheckinRepositoryOption) domain.Repository {
	repo := new(checkinRepository)

	for _, opt := range opts {
		opt(repo)
	}

	// Default DB
	if repo.db == nil {
		db, err := NewCheckinDB()
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

func (r *checkinRepository) Store(c *domain.Checkin) error {
	txn := r.db.Txn(true)

	// Decide if this is a new user
	if c.ID == 0 {
		r.mu.Lock()
		r.lastInsertedId++
		c.ID = r.lastInsertedId
		r.mu.Unlock()

		c.CreatedAt = r.clock.Now()
	}

	c.UpdatedAt = r.clock.Now()

	if err := txn.Insert("checkins", c); err != nil {
		return errors.Wrap(err, "cannot insert checkin")
	}

	txn.Commit()

	return nil
}
