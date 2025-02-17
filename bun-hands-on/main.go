package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func main() {
	dsn := "postgres://user:password@localhost:5432/testdb?sslmode=disable"
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())
	ctx := context.Background()

	fmt.Println("Database connected!")
	// Create table
	db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)

	// テーブルの中身を削除
	_, err = db.NewDelete().Model((*User)(nil)).Exec(ctx)

	// ユーザーの作成
	user := &User{Name: "Alice", Age: 25}
	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted:", user)

	// ユーザーの取得
	var users []User
	db.NewSelect().Model(&users).Scan(ctx)
	fmt.Println("Users:", users)

	// ユーザーの更新
	user.Age = 26
	_, err = db.NewUpdate().Model(user).WherePK().Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated:", user)

	// ユーザーの再取得
	users = nil
	db.NewSelect().Model(&users).Scan(ctx)
	fmt.Println("Users after update:", users)

	// ユーザーの削除
	_, err = db.NewDelete().Model(user).WherePK().Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted user ID:", user.ID)

	// ユーザーの最終確認
	users = nil
	db.NewSelect().Model(&users).Scan(ctx)
	fmt.Println("Users after delete:", users)
}

type User struct {
	ID   int64  `bun:",pk,autoincrement"`
	Name string `bun:",unique"`
	Age  int    `bun:",notnull"`
}
