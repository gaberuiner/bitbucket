package database

import (
	storage "bot/database"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

var (
	statementCreate    = `CREATE TABLE IF NOT EXISTS files(username TEXT,type TEXT, file TEXT)`
	statementInsert    = `INSERT INTO files (username, file, type) VALUES(?, ?, ?)`
	statementSelect    = `SELECT file FROM files WHERE type = ? ORDER BY RANDOM() LIMIT 1`
	statementDelete    = `DELETE FROM files WHERE username = ? AND type = ? AND file = ?`
	statementIsExist   = `SELECT COUNT(*) FROM files WHERE type = ? AND file = ?`
	statementSelectAll = `SELECT file FROM files WHERE type = ?`
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

func (d *Storage) Save(f *storage.File) error {
	_, err := d.db.Exec(statementInsert, f.UserSended, f.FileName, f.Type)
	if err != nil {
		return fmt.Errorf("can't insert in database: %w", err)
	}
	return nil
}

func (d *Storage) SelectALL(Qtype string) (string, error) {
	var ret string
	rows, err := d.db.Query(statementSelectAll, Qtype)
	defer rows.Close()
	if Qtype != "link" && Qtype != "image" {
		for rows.Next() {
			var file string
			err = rows.Scan(&file)
			if err != nil {
				return "", fmt.Errorf("can't scan row: %w", err)
			}
			ret += file + ", " // append the file value to the string with a space
		}

		if err != nil {
			return "", fmt.Errorf("can't select from database: %w", err)
		}
		if ret == "" {
			return ret, fmt.Errorf("no data in storage")
		}
		ret = ret[:len(ret)-2]

		return ret, nil
	} else {
		for rows.Next() {
			var file string
			err = rows.Scan(&file)
			if err != nil {
				return "", fmt.Errorf("can't scan row: %w", err)
			}
			ret += file + " " // append the file value to the string with a space
		}
		return ret, nil
	}
}

func (d *Storage) PickRandom(Qtype string) (string, error) {
	var ret string
	err := d.db.QueryRow(statementSelect, Qtype).Scan(&ret)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("can't select from database %w", err)
	}
	return ret, nil
}

func (d *Storage) IsExist(f *storage.File) (bool, error) {
	var count int
	err := d.db.QueryRow(statementIsExist, f.Type, f.FileName).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("can't check if exist: %w", err)
	}
	return count > 0, nil
}

func (d *Storage) Delete(f *storage.File) error {
	_, err := d.db.Exec(statementDelete, f.UserSended, f.Type, f.FileName)
	if err != nil {
		return fmt.Errorf("can't delete from database %w", err)
	}
	return nil
}
