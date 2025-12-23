package main

import (
	"context"
	"fmt"
	"os"

	bugRepo "github.com/airats/bug_and_fix/internal/repository/bug"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// urlExample := "postgres://user:password@localhost:5432/bug_and_fix?sslmode=disable"
	// conn, err := pgx.Connect(context.Background(), urlExample)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	dbpool, err := pgxpool.New(context.Background(), "postgres://user:password@localhost:5432/bug_and_fix?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	fmt.Println("ok")

	r := bugRepo.NewBugRepository(dbpool)
	d, err := r.Get(2)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(d.ErrorMessage)

	// app.Run()
}
