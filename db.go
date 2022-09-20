package gromer

import (
	"context"
	"database/sql"
)

type SQLInterface[T any] interface {
	SqlOne(ctx context.Context, query string, args ...any) (*T, error)
	SqlMany(ctx context.Context, query string, args ...any) ([]*T, error)
	SqlExecute(ctx context.Context, query string, args ...any) error
}

type DBSQLInterface[T any] struct {
	db *sql.DB
	t  T
}

func (p *DBSQLInterface[T]) SqlOne(ctx context.Context, query string, args ...any) (*T, error) {
	conn, err := p.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.QueryContext(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// cols, _ := rows.Columns()
	for rows.Next() {
		// rows.Scan()
		// for _, c := range cols {
		// if rows.Err()
		// }
	}

	// row.Scan()
	// switch {
	// case err == sql.ErrNoRows:
	// 	log.Fatalf("no user with id %d", id)
	// case err != nil:
	// 	log.Fatal(err)
	// default:
	// 	log.Printf("username is %s\n", username)
	// }
	return nil, nil
}

func (p *DBSQLInterface[T]) SqlMany(ctx context.Context, query string, args ...any) ([]*T, error) {
	return nil, nil
}

func (p *DBSQLInterface[T]) SqlExecute(ctx context.Context, query string, args ...any) error {
	return nil
}

// type User struct {
// }

// func UsersTable(ctx context.Context) SQLInterface[User] {
// 	return &DBSQLInterface[User]{t: User{}}
// }

// func GetNote(ctx context.Context, id string) (*User, error) {
// 	return UsersTable(ctx).SqlOne(ctx, "select * from users where id = :id", id)
// }

// func GetNotes2(ctx context.Context) ([]*User, error) {
// 	return UsersTable(ctx).SqlMany(ctx, "select * from users")
// }
