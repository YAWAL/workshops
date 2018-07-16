package actions

import "github.com/gobuffalo/buffalo"

// CategoriesList default implementation.
func CategoriesList(c buffalo.Context) error {
	return c.Render(200, r.JSON("categories/list.html"))
}

// CategoriesIndex default implementation.
func CategoriesIndex(c buffalo.Context) error {
	return c.Render(200, r.JSON("categories/index.html"))
}
