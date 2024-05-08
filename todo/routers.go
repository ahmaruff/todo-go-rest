package todo

import (
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type apiResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

func writeResponse(c echo.Context, status string, statusCode int, message *string, data interface{}) error {
	response := apiResponse{
		Status:  status,
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	return c.JSON(statusCode, response)
}

func InitTodoRoutes(e *echo.Echo) {
	e.GET("/", listItemHandler)
	e.POST("/", createItemHandler)
	e.GET("/:id", getItemHandler)
	e.PUT("/:id", editItemHandler)
	e.POST("/done", makeItemDoneHandler)
	e.DELETE("/:id", deleteItemHandler)
}

func listItemHandler(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := listItems(ctx)

	if err != nil {
		errMsg := "Internal Server Error"
		return writeResponse(c, "error", echo.ErrInternalServerError.Code, &errMsg, err)
	}

	return writeResponse(c, "success", 200, nil, resp)
}

func getItemHandler(c echo.Context) error {
	ctx := c.Request().Context()
	item_id := c.Param("id")
	Id, err := ulid.Parse(item_id)

	if err != nil {
		msg := "Invalid ID"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &msg, err)
	}

	var resp TodoItem
	item, err := findItem(ctx, Id)

	if err != nil {
		if err == ErrTodoNotFound {
			notfound := "item not found"
			return writeResponse(c, "fail", echo.ErrNotFound.Code, &notfound, nil)
		}
		errMsg := "Internal Server Error"
		return writeResponse(c, "error", echo.ErrInternalServerError.Code, &errMsg, err)
	}

	resp = item
	return writeResponse(c, "success", 200, nil, resp)
}

func createItemHandler(c echo.Context) error {
	if err := c.Request().ParseForm(); err != nil {
		errMsg := "Bad Request"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	ctx := c.Request().Context()

	type titleReq struct {
		Title string `json:"title"`
	}

	title := new(titleReq)

	if err := c.Bind(title); err != nil {
		errMsg := "Bad Request"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	id, err := createItem(ctx, title.Title)

	if err != nil {
		errMsg := "Bad Request"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	var resp struct {
		Id ulid.ULID `json:"id"`
	}

	resp.Id = id

	return writeResponse(c, "success", 200, nil, resp)
}

func makeItemDoneHandler(c echo.Context) error {
	ctx := c.Request().Context()

	type doneReq struct {
		Id string `json:"id"`
	}

	done_req := new(doneReq)
	if err := c.Bind(done_req); err != nil {
		errMsg := "Bad Request"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	Id, err := ulid.Parse(done_req.Id)
	if err != nil {
		errMsg := "Bad Request"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	err = makeItemDone(ctx, Id)
	if err != nil {
		errMsg := "Internal Server Error"
		return writeResponse(c, "fail", echo.ErrInternalServerError.Code, &errMsg, err)
	}

	item, err := findItem(ctx, Id)
	if err != nil {
		if err == ErrTodoNotFound {
			notfound := "item not found"

			return writeResponse(c, "fail", echo.ErrNotFound.Code, &notfound, nil)
		}
		errMsg := "Internal Server Error"
		return writeResponse(c, "error", echo.ErrInternalServerError.Code, &errMsg, err)
	}

	resp := item
	return writeResponse(c, "success", 200, nil, resp)
}

func editItemHandler(c echo.Context) error {
	ctx := c.Request().Context()
	item_id := c.Param("id")

	Id, err := ulid.Parse(item_id)
	if err != nil {
		errMsg := "bad Request 1"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	type editItem struct {
		Title string `json:"title"`
	}

	edit_item := new(editItem)
	if err := c.Bind(edit_item); err != nil {
		errMsg := "bad Request 2"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, edit_item)
	}

	item, err := updateItem(ctx, Id, edit_item.Title)
	if err != nil {
		errMsg := "bad Request 3"
		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	return writeResponse(c, "success", 200, nil, item)
}

func deleteItemHandler(c echo.Context) error {
	ctx := c.Request().Context()

	item_id := c.Param("id")

	Id, err := ulid.Parse(item_id)
	if err != nil {
		errMsg := "bad Request"

		return writeResponse(c, "fail", echo.ErrBadRequest.Code, &errMsg, err)
	}

	_, err = findItem(ctx, Id)
	if err != nil {
		if err == ErrTodoNotFound {
			errMsg := "Not Found"
			return writeResponse(c, "fail", echo.ErrNotFound.Code, &errMsg, err)
		}
		errMsg := "Internal Server Error"

		return writeResponse(c, "error", echo.ErrInternalServerError.Code, &errMsg, err)
	}

	err = deleteItem(ctx, Id)
	if err != nil {
		errMsg := "Internal Server Error"

		return writeResponse(c, "error", echo.ErrInternalServerError.Code, &errMsg, err)
	}

	msg := "Item Deleted"
	return writeResponse(c, "success", 200, &msg, nil)
}
