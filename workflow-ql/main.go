// main.go
package main

import (
	"context"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

var db *bun.DB

func init() {
	pgURL := os.Getenv("DATABASE_URL")
	if pgURL == "" {
		pgURL = "postgres://user:password@localhost:5432/yourdbname"
	}

	driver, err := bun.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	db = driver.NewDB()
	db.AddQueryHook(bun.DebugHook)
}

func main() {
	err := createSchema()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/graphql", func(c *gin.Context) {
		handler := graphql.Handler{
			Schema:   getSchema(),
			Resolver: &Resolver{},
		}
		handler.ServeHTTP(c.Writer, c.Request)
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func createSchema() error {
	ctx := context.Background()
	model := &Model{}

	// AutoMigrate will create tables based on the defined Go struct (Model).
	if err := db.Schema().CreateTable(ctx, model); err != nil {
		return err
	}

	return nil
}
