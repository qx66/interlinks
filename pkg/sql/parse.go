package sql

import (
	"fmt"
	"github.com/percona/go-mysql/query"
	
	"strings"
	"vitess.io/vitess/go/vt/sqlparser"
)

func Fingerprint(sqlQuery string) (fingerprint string, id string) {
	fingerprint = strings.TrimSpace(query.Fingerprint(sqlQuery))
	id = query.Id(fingerprint)
	return
}

func Parse(sqlQuery string) error {
	statement, err := sqlparser.Parse(sqlQuery)
	if err != nil {
		return err
	}
	
	switch t := statement.(type) {
	//case *sqlparser.Select:
	//	fmt.Println("t: select")
	case sqlparser.Statement:
		fmt.Println("Statement")
	case sqlparser.SelectStatement:
		fmt.Println("SelectStatement")
	case sqlparser.DBDDLStatement:
		fmt.Println("DBDDLStatement")
	case sqlparser.DDLStatement:
		fmt.Println("DDLStatement")
	default:
		x := sqlparser.String(t)
		fmt.Println("x: ", x)
	}
	
	newSQL := sqlparser.String(statement)
	fmt.Println("newSQL: ", newSQL)
	return nil
}
