package gomysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

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

	for rows.Next() {
		var id, name string

		err = rows.Scan(&id, &name)
		fmt.Println()
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}

	defer rows.Close()
	assert.Nil(t, err)
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married sql.NullBool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		fmt.Println()
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthDate.Valid {
			fmt.Println("Birthday :", birthDate.Time)
		}
		if married.Valid {
			fmt.Println("Married :", married.Bool)
		}
		fmt.Println("CreatedAt :", createdAt)
	}

	defer rows.Close()

	assert.Nil(t, err)
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)

	if rows.Next() {
		var username string
		err = rows.Scan(&username)

		fmt.Println("Login Berhasil :", username)
	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()

	assert.Nil(t, err)
}

func TestSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)

	if rows.Next() {
		var username string
		err = rows.Scan(&username)

		fmt.Println("Login Berhasil :", username)
	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()

	assert.Nil(t, err)
}
