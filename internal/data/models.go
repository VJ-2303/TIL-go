package data

import "database/sql"

type Models struct {
	TILs TILModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		TILs: TILModel{DB: db},
	}
}
