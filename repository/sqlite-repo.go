package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ea3hsp/test/api"
	"github.com/ea3hsp/test/models"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

type sqlite struct {
	db *sql.DB
}

// NewSqlite creates sqlite repo
func NewSqlite() (api.Repository, error) {
	// creates sqlite
	s := new(sqlite)
	// connection
	err := s.connect()
	if err != nil {
		return nil, err
	}
	// execute migrations
	s.migrations()
	// return
	return s, nil
}

func (s *sqlite) connect() error {
	if s.db != nil {
		return nil
	}
	db, err := sql.Open("sqlite3", "things.sqlite")
	if err != nil {
		return api.ErrorRepoConnect
	}
	s.db = db
	return nil
}

func (s *sqlite) migrations() error {
	// // Query
	q := `CREATE TABLE IF NOT EXISTS things (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		owner VARCHAR(100) NOT NULL,
		name VARCHAR(200) NOT NULL,
		key VARCHAR(50) NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NULL);`
	_, err := s.db.Exec(q)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *sqlite) insertThing(ctx context.Context, thing models.Thing) error {
	// uuid
	key := uuid.Must(uuid.NewV4(), nil).String()
	q := `INSERT INTO things (owner, name, key) VALUES (?, ?, ?)`
	stmt, err := s.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(thing.Owner, thing.Name, key)
	if err != nil {
		return err
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return api.ErrorAffectedRowsExpected
	}
	return nil
}

func (s *sqlite) readThing(ctx context.Context, key string) (*models.Thing, error) {
	thing := &models.Thing{}
	q := fmt.Sprintf(`SELECT id,owner,name,key,created_at,updated_at from things where key='%s'`, key)
	rows, err := s.db.Query(q)
	if err != nil {
		return &models.Thing{}, api.ErrorThingRead
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&thing.ID,
			&thing.Owner,
			&thing.Name,
			&thing.Key,
			&thing.CreatedAt,
			&thing.UpdatedAt,
		)
	}
	return thing, nil
}

func (s *sqlite) CreateThing(ctx context.Context, things ...models.Thing) error {
	for _, thing := range things {
		err := s.insertThing(ctx, thing)
		if err != nil {
			return api.ErrorThingInsert
		}
	}
	return nil
}

func (s *sqlite) ReadThing(ctx context.Context, key string) (*models.Thing, error) {
	return s.readThing(ctx, key)
}

func (s *sqlite) UpdateThing(ctx context.Context, thing models.Thing) error {
	q := `UPDATE things SET INTO things (owner, name, key) VALUES (?, ?, ?)`
	stmt, err := s.db.Prepare(q)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlite) DeleteThing(ctx context.Context, id string) error {
	return nil
}

func (s *sqlite) Close(ctx context.Context) error {
	return s.db.Close()
}
