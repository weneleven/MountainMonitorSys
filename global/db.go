package global

import (
	"database/sql"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
)
var (
	Db *sql.DB
)
