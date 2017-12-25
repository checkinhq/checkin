package acceptance

import (
	"database/sql"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

type DbFeatureContext struct {
	driverName string

	db *sql.DB

	frozen bool
}

func NewDbFeatureContext(driverName string) *DbFeatureContext {
	return &DbFeatureContext{
		driverName: driverName,
	}
}

func (c *DbFeatureContext) FeatureContext(s *godog.Suite) {
	if c.frozen {
		panic("trying to use a frozen feature context")
	}
	c.frozen = true

	s.BeforeScenario(c.beforeScenario)
	s.AfterScenario(c.afterScenario)
}

func (c *DbFeatureContext) beforeScenario(scenario interface{}) {
	var name string

	if s, ok := scenario.(*gherkin.Scenario); ok {
		name = s.Name
	} else if s, ok := scenario.(*gherkin.ScenarioOutline); ok {
		name = s.Name
	}

	// This works with https://github.com/DATA-DOG/go-txdb
	db, err := sql.Open(c.driverName, name)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	c.db = db
}

func (c *DbFeatureContext) afterScenario(scenario interface{}, err error) {
	c.db.Close()
}

func (c *DbFeatureContext) DB() *sql.DB {
	return c.db
}
