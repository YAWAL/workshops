package grifts

import (
	"github.com/markbates/grift/grift"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/YAWAL/workshops/api/models"
	"github.com/gobuffalo/uuid"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here

		for i := 0; i < 15; i++ {
			uu, _ := uuid.NewV4()
			fmt.Println(uu)
			fmt.Println(fake.Title())
			for j := 0; j < 20; j++ {
				alias:= fake.FirstName()
				title:= fake.Title()
				desc:= fake.Month()
				logo:= fake.Continent()
				models.DB.RawQuery("INSERT INTO categories (id, alias, title, descr, logo, parent_id) VALUES ($1, $2, $3, $4, $5, $6)",
					uu,alias, title,desc,logo , uu)
				fmt.Println(logo)

			}

			for j := 0; j < 150; j++ {
				models.DB.RawQuery("INSERT INTO items (id, alias, title, descr, pictures, price, count, parent_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
					uu, fake.FirstName(), fake.Title(), fake.Model(), fake.Brand(), fake.MonthNum(), fake.WeekdayNum(), uu)
			}

		}
		return nil
	})

})
