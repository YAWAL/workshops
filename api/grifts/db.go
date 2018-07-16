package grifts

import (
	"github.com/gobuffalo/uuid"
	"github.com/markbates/grift/grift"

	"fmt"
	"github.com/YAWAL/workshops/api/models"
	"github.com/icrowley/fake"
	"errors"
)

var seed = func(c *grift.Context) error {
	// Add DB seeding stuff here
	fmt.Println("Err: ")

	for i := 0; i < 15; i++ {
		uu, _ := uuid.NewV4()
		fmt.Println(uu)
		for j := 0; j < 20; j++ {
			models.DB.RawQuery("INSERT INTO categories (id, alias, title, desc, logo, parent_id) VALUES ($1, $2, $3, $4, $5, $6)",
				uu, fake.FirstName(), fake.Title(), fake.Model(), fake.Continent(), uu)

		}

		for j := 0; j < 150; j++ {
			models.DB.RawQuery("INSERT INTO items (id, alias, title, desc, pictures, price, count, parent_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
				uu, fake.FirstName(), fake.Title(), fake.Model(), fake.Brand(), fake.MonthNum(), fake.WeekdayNum(), fake.Continent(), uu)
		}

	}

	return errors.New("was here")
}
var _ = grift.Namespace("db", func() {
	fmt.Println("Err: ")

	grift.Desc("seed", "Seeds a database")
	if err := grift.Add("seed", seed); err != nil {
		fmt.Println("Err: ", err)
	}

})

/*
type Category struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Alias     string    `json:"alias" db:"alias"`
	Title     string    `json:"title" db:"title"`
	Desc      string    `json:"desc" db:"desc"`
	Logo      string    `json:"logo" db:"logo"`
	ParentID  uuid.UUID `json:"parent_id" db:"parent_id"`
}
type Item struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Alias      string    `json:"alias" db:"alias"`
	Title      string    `json:"title" db:"title"`
	Desc       string    `json:"desc" db:"desc"`
	Pictures   string    `json:"pictures" db:"pictures"`
	Price      int       `json:"price" db:"price"`
	Count      int       `json:"count" db:"count"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`
}
*/
