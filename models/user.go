package models

type User struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type UserResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        User   `json:"data"`
}

type UsersResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        Users  `json:"data"`
}

type Response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type CreateUser struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

type Book struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Count         int    `json:"count"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}

type CreateBook struct {
	Name          string `json:"name"`
	Count         int    `json:"count"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
}

type Books struct {
	Books []Book `json:"books"`
}

type BookResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        Book   `json:"data"`
}

type BooksResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        Books  `json:"data"`
}

type SaleBook struct {
	Id     string `json:"id"`
	BookId string `json:"book_id"`
	UserId string `json:"user_id"`
}

type SaleBookResponse struct {
	Status      string   `json:"status"`
	Description string   `json:"description"`
	Data        SaleBook `json:"data"`
}

type SaleBooksResponse struct {
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Data        SaleBooks `json:"data"`
}

type CreateSaleBook struct {
	BookId string `json:"book_id"`
	UserId string `json:"user_id"`
}

type SaleBooks struct {
	SaleBooks []SaleBook `json:"sale_books"`
}
