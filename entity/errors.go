package entity

import "fmt"

const (
	// Code SQL Error from https://github.com/go-sql-driver/mysql/blob/master/errors.go
	CodeMySQLDuplicateEntry             = 1062
	CodeMySQLForeignKeyConstraintFailed = 1452
	CodeMySQLTableNotExist              = 1146
)

var (
	ErrorFarmNotFound     error = fmt.Errorf("Farm Not Found")
	ErrorFarmAlreadyExist error = fmt.Errorf("Farm Already Exist")
)
