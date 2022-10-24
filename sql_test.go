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

	script := "INSERT INTO customer(id,name) VALUES('melza', 'Melza')"
	_, err := db.ExecContext(ctx, script) //gunakan ExecContext apabila tidak memerlukan hasil
	result := assert.Nil(t, err)
	if result {
		fmt.Println("Success Insert New Customer ")
	}
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)

	fmt.Println("===========")
	for rows.Next() {
		var id, name string

		err = rows.Scan(&id, &name)
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		fmt.Println("===========")
	}

	defer rows.Close()
	assert.Nil(t, err)
}
