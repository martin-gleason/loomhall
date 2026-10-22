// Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
package store

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/martin-gleason/loomhall/internal/id"
)

// FR67: one suite, every engine. A backend is a name and a way to open it.
type backend struct {
	name string
	open func(t *testing.T) *Store
}

func backends(t *testing.T) []backend {
	t.Helper()
	return []backend{
		{"sqlite", func(t *testing.T) *Store {
			s, err := OpenSQLite(filepath.Join(t.TempDir(), "probe.db"))
			if err != nil {
				t.Fatal(err)
			}
			return s
		}},
		{"postgres", func(t *testing.T) *Store {
			dsn := os.Getenv("LOOMHALL_TEST_POSTGRES_DSN")
			if dsn == "" {
				t.Skip("LOOMHALL_TEST_POSTGRES_DSN unset: the Postgres half of FR67 did not run here")
			}
			s, err := OpenPostgres(dsn)
			if err != nil {
				t.Fatal(err)
			}
			return s
		}},
	}
}

func fresh(t *testing.T, b backend) *Store {
	t.Helper()
	s := b.open(t)
	t.Cleanup(func() { _ = s.Close() })
	ctx := context.Background()
	if err := s.reset(ctx); err != nil {
		t.Fatal(err)
	}
	if err := s.Migrate(ctx); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestStore_PutThenGet(t *testing.T) {
	for _, b := range backends(t) {
		t.Run(b.name, func(t *testing.T) {
			s := fresh(t, b)
			ctx := context.Background()
			want := &Record{ID: id.New(), Origin: "test-install", Body: "a story", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
			if err := s.Put(ctx, want); err != nil {
				t.Fatal(err)
			}
			got, err := s.Get(ctx, want.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != want.ID || got.Origin != want.Origin || got.Body != want.Body || !got.CreatedAt.Equal(want.CreatedAt) {
				t.Fatalf("round trip changed the record:\n got %+v\nwant %+v", got, want)
			}
		})
	}
}

func TestStore_DuplicateIDIsRefused(t *testing.T) {
	for _, b := range backends(t) {
		t.Run(b.name, func(t *testing.T) {
			s := fresh(t, b)
			ctx := context.Background()
			r := &Record{ID: id.New(), Origin: "test-install", Body: "x", CreatedAt: time.Now()}
			if err := s.Put(ctx, r); err != nil {
				t.Fatal(err)
			}
			if err := s.Put(ctx, r); err == nil {
				t.Fatal("second Put with the same ID succeeded; the primary key is not enforced")
			}
		})
	}
}

func TestStore_MissingIsErrNoRows(t *testing.T) {
	for _, b := range backends(t) {
		t.Run(b.name, func(t *testing.T) {
			s := fresh(t, b)
			_, err := s.Get(context.Background(), id.New())
			if !errors.Is(err, sql.ErrNoRows) {
				t.Fatalf("Get of a missing id: err = %v, want sql.ErrNoRows", err)
			}
		})
	}
}

func TestOpenPostgres_BadDSNIsAnErrorNotAPanic(t *testing.T) {
	for _, dsn := range []string{"localhost:5433/loomhall_test", "", "postgres://%zz"} {
		s, err := OpenPostgres(dsn)
		if err == nil || s != nil {
			t.Fatalf("OpenPostgres(%q) = %v, %v; want nil store and an error", dsn, s, err)
		}
	}
}

// Timing for D2 and RR1. Not an assertion: it prints what happened on this machine.
func TestStore_Timing(t *testing.T) {
	const n = 1000
	for _, b := range backends(t) {
		t.Run(b.name, func(t *testing.T) {
			s := fresh(t, b)
			ctx := context.Background()
			ids := make([]string, n)
			start := time.Now()
			for i := range ids {
				ids[i] = id.New()
				if err := s.Put(ctx, &Record{ID: ids[i], Origin: "timing", Body: "lorem ipsum", CreatedAt: time.Now()}); err != nil {
					t.Fatal(err)
				}
			}
			wrote := time.Since(start)
			start = time.Now()
			for _, x := range ids {
				if _, err := s.Get(ctx, x); err != nil {
					t.Fatal(err)
				}
			}
			read := time.Since(start)
			t.Logf("%s: %d single-row inserts in %v (%v each), %d reads by id in %v (%v each)",
				b.name, n, wrote, wrote/n, n, read, read/n)
		})
	}
}
