package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/icrowley/fake"
	"time"
	"github.com/gobuffalo/uuid"
	"github.com/YAWAL/workshops/api/models"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here

		models.DB.RawQuery("Truncate table categories")
		models.DB.RawQuery("Truncate table items")
		for i := 0; i < 15; i++ {
			uu, _ := uuid.NewV4()
			for j := 0; j < 20; j++ {
				uu_categories, _ := uuid.NewV4()

				alias:= fake.FirstName()
				title:= fake.Title()
				desc:= fake.Month()
				logo:= fake.Continent()
				models.DB.RawQuery("INSERT INTO categories (id, alias, title, descr, logo, parent_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
					uu_categories,alias, title,desc,logo , uu, time.Now(),time.Unix(1531747237,567))

			}

			for k := 0; k < 150; k++ {
				uu_items, _ := uuid.NewV4()

				models.DB.RawQuery("INSERT INTO items (id, alias, title, descr, pictures, price, count, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
					uu_items, fake.FirstName(), fake.Title(), fake.Model(), fake.Brand(), fake.MonthNum(), fake.WeekdayNum(), uu, time.Now(),time.Unix(1531741217,567))
			}

		}
		return nil
	})

})
