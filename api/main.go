package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/gobuffalo/uuid"
	"github.com/jinzhu/gorm"
	"github.com/YAWAL/ERP/back-end/tool/logging"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"time"
)

// main is the starting point to your Buffalo application.
// you can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)
func main() {
	conn, err := PGconn()

	if err != nil {
		fmt.Println(err)
	}
	conn.Exec("Truncate table categories")
	conn.Exec("Truncate table items")
	for i := 0; i < 15; i++ {
		uu, _ := uuid.NewV4()
		for j := 0; j < 20; j++ {
			uu_categories, _ := uuid.NewV4()

			alias:= fake.FirstName()
			title:= fake.Title()
				desc:= fake.Month()
				logo:= fake.Continent()
			conn.Exec("INSERT INTO categories (id, alias, title, descr, logo, parent_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
				uu_categories,alias, title,desc,logo , uu, time.Now(),time.Unix(1531747237,567))

		}

		for k := 0; k < 150; k++ {
			uu_items, _ := uuid.NewV4()

			conn.Exec("INSERT INTO items (id, alias, title, descr, pictures, price, count, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
				uu_items, fake.FirstName(), fake.Title(), fake.Model(), fake.Brand(), fake.MonthNum(), fake.WeekdayNum(), uu, time.Now(),time.Unix(1531741217,567))
		}

	}
}
func PGconn() (db *gorm.DB, err error) {
	// real configs -> config file
	db, err = gorm.Open("postgres", "user=postgres dbname=api_development sslmode=disable password=postgres")
	if err != nil {
		logging.Log.Errorf("Error during connection to Postgres has occurred %s", err.Error())
	} else {
		logging.Log.Info("Connection to Pg has been established")
	}
	return db, err
}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heaving lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/
