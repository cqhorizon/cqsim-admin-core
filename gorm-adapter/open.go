//go:build !sqlite3
// +build !sqlite3

package gormadapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var opens = map[string]func(string) gorm.Dialector{
	"mysql": mysql.Open,
}
