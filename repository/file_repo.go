package repository

// sqlx doc
// http://jmoiron.github.io/sqlx/

import (
	"fmt"
	"zupload/config"
	"zupload/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type FileRepo struct {
	FileTable string
}

func NewFileRpo() *FileRepo {
	return &FileRepo{
		FileTable: model.File{}.TableName(),
	}
}

func (f *FileRepo) CreateOne(row *model.File) (int64, error) {
	db, err := sqlx.Connect("sqlite3", config.Conf.DataBase.Sqlite3.DBName)
	if err != nil {
		return 0, errors.WithMessage(err, "connect the DB")
	}

	defer db.Close()

	sql := fmt.Sprintf(`INSERT INTO %s (file_name, created_at) VALUES (?, ?)`, f.FileTable)
	if result, err := db.Exec(sql, row.FileName, row.CreatedAt); err != nil {
		return 0, errors.WithMessage(err, "db.Exec Insert")
	} else {
		return result.LastInsertId()
	}

}

func (f *FileRepo) List(limit int) ([]model.File, error) {
	db, err := sqlx.Connect("sqlite3", config.Conf.DataBase.Sqlite3.DBName)
	if err != nil {
		return nil, errors.WithMessage(err, "connect the DB")
	}

	defer db.Close()
	list := make([]model.File, 0, limit)

	err = db.Select(&list, fmt.Sprintf(`SELECT * FROM %s Order By id DESC LIMIT ?`, f.FileTable), limit)
	if err != nil {
		return nil, errors.WithMessage(err, "db.Select")
	}

	return list, nil
}
