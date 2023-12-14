package sql

import (
	"fmt"
	"github.com/qx66/interlinks/pkg/sql/ast"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFingerprint(t *testing.T) {
	sqlQ := "SELECT * FROM user"
	f, id := Fingerprint(sqlQ)
	fmt.Printf("Fingerprint: %s, id: %s\n", f, id)
}

func TestParse(t *testing.T) {
	sqlQ := "SELECT * FROM user WHERE id = 1"
	err := Parse(sqlQ)
	require.NoError(t, err, "解析SQL失败")
}

func TestPretty(t *testing.T) {
	
	sqlQ := "SELECT * FROM user WHERE id = 1 limit 10"
	out := ast.Pretty(sqlQ, "builtin")
	fmt.Println("pretty out: ", out)
}
