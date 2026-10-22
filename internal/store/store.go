// Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
// Package store is the only door to the database (FR67). Everything above it talks
// to Store; nothing above it knows which SQL engine is behind it.
//
// The probe behind D2: one table, written and read through Bun, on pure-Go SQLite
// and on Postgres, from the same code and the same tests.
package store

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	_ "modernc.org/sqlite" // registers the "sqlite" driver, no C toolchain needed
)

// Record is the smallest thing worth storing: an ID that is unique across every
// install, the install that made it, and a body. NFR13 is the two fields on top.
type Record struct {
	bun.BaseModel `bun:"table:records"`

	ID        string    `bun:"id,pk"`
	Origin    string    `bun:"origin,notnull"`
	Body      string    `bun:"body,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
}

// Store wraps one open database. Methods take a context so a request that is
// cancelled stops its queries too.
type Store struct {
	db *bun.DB
}

// OpenSQLite opens (or creates) the database file at path.
func OpenSQLite(path string) (*Store, error) {
	sqldb, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	// SQLite allows one writer at a time; a single connection makes that a queue
	// instead of a "database is locked" error.
	sqldb.SetMaxOpenConns(1)
	return &Store{db: bun.NewDB(sqldb, sqlitedialect.New())}, nil
}

// OpenPostgres connects with a postgres:// DSN. A DSN this function cannot read is
// returned as an error: the driver's own parser panics, and a typo in a guild's config
// must not crash the server at startup.
func OpenPostgres(dsn string) (s *Store, err error) {
	u, perr := url.Parse(dsn)
	if perr != nil || (u.Scheme != "postgres" && u.Scheme != "postgresql") {
		return nil, fmt.Errorf("store: postgres DSN must start with postgres:// or postgresql://, got %q", dsn)
	}
	defer func() {
		if r := recover(); r != nil {
			s, err = nil, fmt.Errorf("store: cannot read postgres DSN: %v", r)
		}
	}()
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	return &Store{db: bun.NewDB(sqldb, pgdialect.New())}, nil
}

// Migrate creates what is missing. Real migrations arrive with F10's pack schemas;
// for the probe, create-if-not-exists is the whole story.
func (s *Store) Migrate(ctx context.Context) error {
	_, err := s.db.NewCreateTable().Model((*Record)(nil)).IfNotExists().Exec(ctx)
	return err
}

// Put inserts one record. A second Put with the same ID is an error, by the
// primary key, on every engine.
func (s *Store) Put(ctx context.Context, r *Record) error {
	_, err := s.db.NewInsert().Model(r).Exec(ctx)
	return err
}

// Get returns the record with that ID, or sql.ErrNoRows.
func (s *Store) Get(ctx context.Context, id string) (*Record, error) {
	r := new(Record)
	if err := s.db.NewSelect().Model(r).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return r, nil
}

// reset drops the table so every test run starts clean. Unexported on purpose: the
// only door to the database does not hand every caller a one-line table drop.
func (s *Store) reset(ctx context.Context) error {
	_, err := s.db.NewDropTable().Model((*Record)(nil)).IfExists().Exec(ctx)
	return err
}

// Close releases the connection.
func (s *Store) Close() error { return s.db.Close() }
