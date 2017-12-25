package inmem_test

import (
	"testing"

	"time"

	"github.com/checkinhq/checkin/pkg/checkin/domain"
	"github.com/checkinhq/checkin/pkg/checkin/infrastructure/inmem"
	"github.com/hashicorp/go-memdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckinRepository_Create(t *testing.T) {
	var db *memdb.MemDB
	db, err := inmem.NewCheckinDB()
	require.NoError(t, err)

	repo := inmem.NewCheckinRepository(inmem.CheckinDB(db))

	c := newCheckin(t)
	c.ID = 0

	err = repo.Store(c)
	require.NoError(t, err)

	txn := db.Txn(false)
	defer txn.Abort()

	c2, err := txn.First("checkins", "id", c.ID)
	require.NoError(t, err)
	assert.Equal(t, c, c2)
}

func TestCheckinRepository_Update(t *testing.T) {
	var db *memdb.MemDB
	db, err := inmem.NewCheckinDB()
	require.NoError(t, err)

	repo := inmem.NewCheckinRepository(inmem.CheckinDB(db))

	c := newCheckin(t)

	saveCheckin(t, db, c)

	c.GoalsReached = false

	err = repo.Store(c)
	require.NoError(t, err)

	txn := db.Txn(false)
	defer txn.Abort()

	c2, err := txn.First("checkins", "id", c.ID)
	require.NoError(t, err)
	assert.Equal(t, c, c2)
}

func TestCheckinRepository_CreateMore(t *testing.T) {
	var db *memdb.MemDB
	db, err := inmem.NewCheckinDB()
	require.NoError(t, err)

	repo := inmem.NewCheckinRepository(inmem.CheckinDB(db))

	c := newCheckin(t)
	c.ID = 0
	c2 := newCheckin(t)
	c2.ID = 0

	err = repo.Store(c)
	require.NoError(t, err)

	err = repo.Store(c2)
	require.NoError(t, err)

	txn := db.Txn(false)
	defer txn.Abort()

	cactual, err := txn.First("checkins", "id", c.ID)
	require.NoError(t, err)
	assert.Equal(t, c, cactual)

	c2actual, err := txn.First("checkins", "id", c2.ID)
	require.NoError(t, err)
	assert.Equal(t, c2, c2actual)
}

func newCheckin(t *testing.T) *domain.Checkin {
	c := &domain.Checkin{
		ID:     1,
		UserID: 1,
		Date:   time.Now().Truncate(time.Hour * 24),

		Previous:     "I finished doing something",
		GoalsReached: true,
		Next:         "I will start doing something",
		Blockers:     "None",

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return c
}

func saveCheckin(t *testing.T, db *memdb.MemDB, c *domain.Checkin) {
	txn := db.Txn(true)

	err := txn.Insert("checkins", c)

	require.NoError(t, err)

	txn.Commit()
}
