package sql

import (
	"fmt"
	"github.com/percona/go-mysql/query"
)

func Fingerprint(sql string) {
	f := query.Fingerprint("select * from user where uuid = '123'")
	fmt.Println(f)
}
