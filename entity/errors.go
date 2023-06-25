package entity

import "fmt"

const (
	// Code SQL Error from https://github.com/go-sql-driver/mysql/blob/master/errors.go
	CodeMySQLDuplicateEntry             = 1062
	CodeMySQLForeignKeyConstraintFailed = 1452
	CodeMySQLTableNotExist              = 1146
)

var (
	// error farm
	ErrorFarmNotFound     error = fmt.Errorf("Farm Not Found")
	ErrorFarmAlreadyExist error = fmt.Errorf("Farm Already Exist")

	// error pond
	ErrorPondNotFound     error = fmt.Errorf("Pond Not Found")
	ErrorPondAlreadyExist error = fmt.Errorf("Pond Already Exist")
)
