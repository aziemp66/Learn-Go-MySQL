package gomysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id,name) VALUES('melza', 'Melza')"
	_, err := db.ExecContext(ctx, query)
	result := assert.Nil(t, err)
	if result {
		fmt.Println("Success Insert New Customer ")
	}
}
