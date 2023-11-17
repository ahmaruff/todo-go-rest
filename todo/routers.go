package todo

import (
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func InitTodoRoutes(e *echo.Echo) {
	e.GET("/", listItemHandler)
	e.POST("/", createItemHandler)
	e.GET("/:id", getItemHandler)
	e.POST("/done", makeItemDoneHandler)
	e.DELETE("/:id", deleteItemHandler)
}

func listItemHandler(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := listItems(ctx)

	if err != nil {
		return echo.NewHTTPError(echo.ErrInternalServerError.Code, err)
	}

	return c.JSON(200, resp)
}

func getItemHandler(c echo.Context) error {
	ctx := c.Request().Context()

	item_id := c.Param("id")

	Id, err := ulid.Parse(item_id)

	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	var resp TodoItem

	item, err := findItem(ctx, Id)

	if err != nil {
		if err == ErrTodoNotFound {
			notfound := map[string]string{
				"message": "item not found",
			}

			return c.JSON(echo.ErrNotFound.Code, notfound)
		}

		return echo.NewHTTPError(echo.ErrInternalServerError.Code, err)
	}

	resp = item

	return c.JSON(200, resp)
}

func createItemHandler(c echo.Context) error {
	if err := c.Request().ParseForm(); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	ctx := c.Request().Context()

	type titleReq struct {
		Title string `json:"title"`
	}
	title := new(titleReq)

	if err := c.Bind(title); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	id, err := createItem(ctx, title.Title)

	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	var resp struct {
		Id ulid.ULID `json:"id"`
	}

	resp.Id = id

	return c.JSON(201, resp)
}

func makeItemDoneHandler(c echo.Context) error {
	ctx := c.Request().Context()

	type doneReq struct {
		Id string `json:"id"`
	}

	done_req := new(doneReq)

	if err := c.Bind(done_req); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	Id, err := ulid.Parse(done_req.Id)

	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	err = makeItemDone(ctx, Id)

	if err != nil {
		return echo.NewHTTPError(echo.ErrInternalServerError.Code, err)
	}

	resp := map[string]string{
		"message": "item updated",
	}

	return c.JSON(200, resp)
}

func deleteItemHandler(c echo.Context) error {
	ctx := c.Request().Context()

	item_id := c.Param("id")

	Id, err := ulid.Parse(item_id)

	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	_, err = findItem(ctx, Id)

	if err != nil {
		if err == ErrTodoNotFound {
			notfound := map[string]string{
				"message": "item not found",
			}

			return c.JSON(echo.ErrNotFound.Code, notfound)
		}

		return echo.NewHTTPError(echo.ErrInternalServerError.Code, err)
	}

	err = deleteItem(ctx, Id)

	if err != nil {
		return echo.NewHTTPError(echo.ErrInternalServerError.Code, err)
	}

	resp := map[string]string{
		"message": "item deleted",
	}

	return c.JSON(200, resp)
}
