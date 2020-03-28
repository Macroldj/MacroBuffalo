package actions

import (
	"strconv"

	"github.com/chenghung/buffalo_api/models"
	"github.com/gobuffalo/buffalo"
)

func UserIndex(c buffalo.Context) error {
	page, _ := strconv.Atoi(c.Param("page"))
	perPage, _ := strconv.Atoi(c.Param("per_page"))

	c.Logger().Debug("page: ", page)
	c.Logger().Debug("per page: ", perPage)

	users := models.UsersWithPagination{}
	query := models.DB.Paginate(page, perPage)
	query.All(&users.Users)

	users.Paginator = query.Paginator

	c.Logger().Debug("paginator", query.Paginator)

	return c.Render(200, r.JSON(users))
}

func UserShow(c buffalo.Context) error {
	id := c.Param("id")
	user := models.User{}

	err := models.DB.Find(&user, id)
	if err != nil {
		return err
	}

	return c.Render(200, r.JSON(user))
}
