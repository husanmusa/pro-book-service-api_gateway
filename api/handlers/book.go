package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/husanmusa/pro-book-service-api_gateway/models"
	"io/ioutil"
	"net/http"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "CreateUserRequestBody"
// @Success 201 {object} models.User "User data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	responseBody := bytes.NewBuffer(jsonData)
	resp, err := http.Post(h.cfg.BookServiceURL+c.Request.URL.Path, "application/json", responseBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Post error Create User", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Create User", err)
	}
	var user models.UserResponse
	fmt.Println(responseBody)
	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @ID get_user
// @Router /user/{id} [GET]
// @Summary Get User
// @Description Get User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.UserResponse "User data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) GetUser(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error", err)
	}
	var user models.UserResponse
	json.Unmarshal(body, &user)
	c.JSON(http.StatusCreated, user)
}

// ListUser	godoc
// @ID list_user
// @Router /user [GET]
// @Summary List User
// @Description List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} models.UsersResponse "User data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) ListUser(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error", err)
	}
	var user models.UsersResponse

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// FindBookFromUsers godoc
// @ID find_book
// @Router /user/find_book/{id} [GET]
// @Summary List User has a book
// @Description List User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Users "ListUserResponseBody"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) FindBookFromUsers(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error", err)
	}
	var user models.UsersResponse
	json.Unmarshal(body, &user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "UpdateUserRequestBody"
// @Success 200 {object} models.User "User data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	client := &http.Client{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	reqBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest(http.MethodPut, h.cfg.BookServiceURL+c.Request.URL.Path, reqBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Request Update User", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Update User", err)
	}
	var user models.UserResponse

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// DeleteUser godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.SuccessModel{data=string} "User data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {
	client := &http.Client{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	reqBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest(http.MethodDelete, h.cfg.BookServiceURL+c.Request.URL.Path, reqBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Request Delete User", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Delete User", err)
	}
	var user models.Response

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// CreateBook godoc
// @ID create_book
// @Router /book [POST]
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param book body models.CreateBook true "CreateBookRequestBody"
// @Success 201 {object} models.BookResponse "Book data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) CreateBook(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	responseBody := bytes.NewBuffer(jsonData)
	resp, err := http.Post(h.cfg.BookServiceURL+c.Request.URL.Path, "application/json", responseBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Post error Create Book", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Create Book", err)
	}
	var user models.BookResponse
	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// GetBook godoc
// @ID get_book
// @Router /book/{id} [GET]
// @Summary Get Book
// @Description Get Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book "Book data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) GetBook(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get Book error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse Book error", err)
	}
	var user models.BookResponse
	json.Unmarshal(body, &user)
	c.JSON(http.StatusCreated, user)
}

// ListBook godoc
// @ID list_book
// @Router /book [GET]
// @Summary List Book
// @Description List Book
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} models.Books "ListBookResponseBody"
// @Response 400 {object} response.SuccessModel{data=string} "Invalid Argument"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) ListBook(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get List Book error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse list book error", err)
	}
	var user models.BooksResponse

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// ListBooksByUserId godoc
// @ID list_book_by_user_id
// @Router /book/user/{id} [GET]
// @Summary List Book User
// @Description List Book User
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string false "id"
// @Success 200 {object} models.Books "ListBookResponseBody"
// @Response 400 {object} response.SuccessModel{data=string} "Invalid Argument"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) ListBooksByUserId(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get List book by user id error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse List book by user id error", err)
	}
	var user models.BooksResponse
	json.Unmarshal(body, &user)
	c.JSON(http.StatusCreated, user)
}

// UpdateBook godoc
// @ID update_book
// @Router /book/{id} [PUT]
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.CreateBook true "UpdateBookRequestBody"
// @Success 200 {object} models.Book "Book data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) UpdateBook(c *gin.Context) {
	client := &http.Client{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	reqBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest(http.MethodPut, h.cfg.BookServiceURL+c.Request.URL.Path, reqBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Request Update Book", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Book User", err)
	}
	var user models.BookResponse

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// DeleteBook godoc
// @ID delete_book
// @Router /book/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} response.SuccessModel{data=string} "Book data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) DeleteBook(c *gin.Context) {
	client := &http.Client{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	reqBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest(http.MethodDelete, h.cfg.BookServiceURL+c.Request.URL.Path, reqBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Request Delete Book", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Delete Book", err)
	}
	var user models.Response

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// CreateSaleBook godoc
// @ID create_sale_book
// @Router /sale [POST]
// @Summary Create SaleBook
// @Description Create SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param sale body models.CreateSaleBook true "CreateSaleBookRequestBody"
// @Success 201 {object} models.SaleBook "SaleBook data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) CreateSaleBook(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	responseBody := bytes.NewBuffer(jsonData)
	resp, err := http.Post(h.cfg.BookServiceURL+c.Request.URL.Path, "application/json", responseBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Post error Create Sale Book", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Create Sale Book", err)
	}
	var user models.SaleBookResponse
	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// GetSaleBook godoc
// @ID get_sale_book
// @Router /sale/{id} [GET]
// @Summary Get SaleBook
// @Description Get SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param id path string true "SaleBook ID"
// @Success 200 {object} models.SaleBook "SaleBook data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) GetSaleBook(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get Sale Book error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse Sale Book error", err)
	}
	var user models.SaleBookResponse
	json.Unmarshal(body, &user)
	c.JSON(http.StatusCreated, user)
}

// ListSaleBook godoc
// @ID list_sale_book
// @Router /sale [GET]
// @Summary List SaleBook
// @Description List SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} models.SaleBooks "ListSaleBookResponseBody"
// @Response 400 {object} response.SuccessModel{data=string} "Invalid Argument"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) ListSaleBook(c *gin.Context) {
	resp, err := http.Get(h.cfg.BookServiceURL + c.Request.URL.Path)
	if err != nil {
		h.handleErrorResponse(c, 500, "Get List sale Book error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse list sale book error", err)
	}
	var user models.SaleBooksResponse

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// UpdateSaleBook godoc
// @ID update_sale_book
// @Router /sale/{id} [PUT]
// @Summary Update SaleBook
// @Description Update SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param id path string true "SaleBook ID"
// @Param sale body models.CreateSaleBook true "UpdateSaleBookRequestBody"
// @Success 200 {object} models.SaleBook "SaleBook data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) UpdateSaleBook(c *gin.Context) {
	client := &http.Client{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	reqBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest(http.MethodPut, h.cfg.BookServiceURL+c.Request.URL.Path, reqBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Request Update Book", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Book User", err)
	}
	var user models.SaleBookResponse

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}

// DeleteSaleBook godoc
// @ID delete_sale_book
// @Router /sale/{id} [DELETE]
// @Summary Delete SaleBook
// @Description Delete SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param id path string true "SaleBook ID"
// @Success 200 {object} response.SuccessModel{data=string} "SaleBook data"
// @Response 400 {object} response.SuccessModel{data=string} "Bad Request"
// @Failure 500 {object} response.SuccessModel{data=string} "Server Error"
func (h *Handler) DeleteSaleBook(c *gin.Context) {
	client := &http.Client{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	//postBody, err := json.Marshal(jsonData)
	reqBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest(http.MethodDelete, h.cfg.BookServiceURL+c.Request.URL.Path, reqBody)
	if err != nil {
		h.handleErrorResponse(c, 500, "Request Delete Sale Book", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.handleErrorResponse(c, 500, "parse error Delete Sale Book", err)
	}
	var user models.Response

	json.Unmarshal(body, &user)

	c.JSON(http.StatusCreated, user)
}
