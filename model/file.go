package model

type File struct {
	Id        uint   `db:"id"`
	FileName  string `db:"file_name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (f File) TableName() string {
	return "file"
}
