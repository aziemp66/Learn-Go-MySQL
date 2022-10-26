package gomysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/aziemp66/Learn-Go-MySQL/entity"
	"github.com/aziemp66/Learn-Go-MySQL/repository"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "Yahuu@gmail.com",
		Comment: "Wadaw anjay",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 90)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(comments)

}
