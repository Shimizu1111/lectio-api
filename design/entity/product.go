package entity

type User struct {
	UserID           uint
	UserName         string
	Email            string
	EmailVerifiedAt  string
	UserPassword     string
	HasBackstagePass bool
	IsSuperuser      bool
	IsDeleted        bool
	IsBanned         bool
	CreatedAt        string
	UpdatedAt        string
}

type Book struct {
	BookID           uint
	UserID           uint
	BookName         string
	Author           string
	Publisher        string
	Price            int
	RegistrationDate string
	IsDeleted        bool
	CreatedAt        string
	UpdatedAt        string
}

type Books struct {
	Books []*Book
}
