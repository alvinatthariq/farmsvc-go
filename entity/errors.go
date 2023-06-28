package entity

import "fmt"

const (
	// Code SQL Error from https://github.com/go-sql-driver/mysql/blob/master/errors.go
	CodeMySQLDuplicateEntry             = 1062
	CodeMySQLForeignKeyConstraintFailed = 1452
	CodeMySQLTableNotExist              = 1146
)

var (
	ErrorFarmNotFound             error = fmt.Errorf("Farm Not Found")
	ErrorFarmAlreadyExist         error = fmt.Errorf("Farm Already Exist")
	ErrorFarmIDRequired           error = fmt.Errorf("Farm ID Required")
	ErrorFarmIDMaxLength          error = fmt.Errorf("Farm ID Max Length is 36")
	ErrorFarmNameRequired         error = fmt.Errorf("Farm Name Required")
	ErrorFarmNameMaxLength        error = fmt.Errorf("Farm Name Max Length is 100")
	ErrorFarmDescriptionRequired  error = fmt.Errorf("Farm Description Required")
	ErrorFarmDescriptionMaxLength error = fmt.Errorf("Farm Description Max Length is 150")
	ErrorPondNotFound             error = fmt.Errorf("Pond Not Found")
	ErrorPondAlreadyExist         error = fmt.Errorf("Pond Already Exist")
	ErrorPondIDRequired           error = fmt.Errorf("Pond ID Required")
	ErrorPondIDMaxLength          error = fmt.Errorf("Pond ID Max Length is 36")
	ErrorPondNameRequired         error = fmt.Errorf("Pond Name Required")
	ErrorPondNameMaxLength        error = fmt.Errorf("Pond Name Max Length is 100")
	ErrorPondDescriptionRequired  error = fmt.Errorf("Pond Description Required")
	ErrorPondDescriptionMaxLength error = fmt.Errorf("Pond Description Max Length is 150")
)
