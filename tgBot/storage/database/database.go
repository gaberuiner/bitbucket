package database

import (
	"database/sql"
	"fmt"
	"testbot/storage"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

var (
	statementCreate  = `CREATE TABLE IF NOT EXISTS films(username TEXT, film TEXT)`
	statementInsert  = `INSERT INTO films (username, film) VALUES(?, ?)`
	statementSelect  = `SELECT film FROM films WHERE username = ? ORDER BY RANDOM() LIMIT 1`
	statementDelete  = `DELETE FROM films WHERE username = ? AND film = ?`
	statementIsExist = `SELECT COUNT(*) FROM pages WHERE username = ? AND film = ?`
)

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open db : %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to db: %w", err)
	}
	return &Storage{
		db: db,
	}, nil
}

func (d *Storage) Init() error {
	_, err := d.db.Exec(statementCreate)
	if err != nil {
		return fmt.Errorf("can't create db: %w", err)
	}
	return nil
}

// type Storage interface {
// 	Save(f *Film) error
// 	PickRandom(userName string) (*Film, error)
// 	IsExist(f *Film) (bool, error)
// 	Delete(f *Film) error
// }

func (d *Storage) Save(f *storage.Film) error {
	_, err := d.db.Exec(statementInsert, f.UserSended, f.FilmName)
	if err != nil {
		return fmt.Errorf("can't insert in database: %w", err)
	}
	return nil
}

func (d *Storage) PickRandom(userName string) (*storage.Film, error) {
	var ret string
	err := d.db.QueryRow(statementSelect, userName).Scan(&ret)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can't select from database %w", err)
	}
	return &storage.Film{UserSended: userName, FilmName: ret}, nil
}

func (d *Storage) IsExist(f *storage.Film) (bool, error) {
	var count int
	err := d.db.QueryRow(statementIsExist, f.UserSended, f.FilmName).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("can't check if exist: %w", err)
	}
	return count > 0, nil
}

func (d *Storage) Delete(f *storage.Film) error {
	_, err := d.db.Exec(statementDelete, f.UserSended, f.FilmName)
	if err != nil {
		return fmt.Errorf("can't delete from database %w", err)
	}
	return nil
}
