package responses

import "time"

type Currencies struct {
	CurrencyId   int64
	Symbol       string
	Currency     string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
}

type Countries struct {
	CountryId       int64
	Country         string
	Description     string
	CountryCode     string
	DefaultCurrency *Currencies
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       int
	ModifiedBy      int
}

// type Users struct {
// 	UserId        int64
// 	UserDetails   int64
// 	ImagePath     string
// 	UserType      int
// 	FullName      string
// 	Username      string
// 	Password      string
// 	Email         string
// 	PhoneNumber   string
// 	Gender        string
// 	Dob           time.Time
// 	Address       string
// 	IdType        string
// 	IdNumber      string
// 	MaritalStatus string
// 	Active        int
// 	Role          int64
// 	IsVerified    bool
// 	DateCreated   time.Time
// 	DateModified  time.Time
// 	CreatedBy     int
// 	ModifiedBy    int
// }

type Branches struct {
	BranchId      int64
	Branch        string
	Country       *Countries
	Location      string
	PhoneNumber   string
	Active        int
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
	BranchManager *Users
}

type BranchesOriResponseDTO struct {
	StatusCode int
	Branches   *[]Branches
	StatusDesc string
}

type BranchOriResponseDTO struct {
	StatusCode int
	Branch     *Branches
	StatusDesc string
}

type CountriesOriResponseDTO struct {
	StatusCode int
	Countries  *[]Countries
	StatusDesc string
}

type CountryResponseDTO struct {
	StatusCode int
	Country    *Countries
	StatusDesc string
}

type CurrenciesResponseDTO struct {
	StatusCode int
	Currencies *[]Currencies
	StatusDesc string
}

type CurrencyResponseDTO struct {
	StatusCode int
	Currency   *Currencies
	StatusDesc string
}

type Users struct {
	UserId        int64
	UserDetails   *UserExtraDetails
	ImagePath     string
	UserType      int
	FullName      string
	Username      string
	Password      string
	Email         string
	PhoneNumber   string
	Gender        string
	Dob           time.Time
	Address       string
	IdType        string
	IdNumber      string
	MaritalStatus string
	Active        int
	Role          *Roles
	IsVerified    bool
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
}

type Roles struct {
	RoleId       int64
	Role         string
	Description  string
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
	Active       int
}

type Shops struct {
	ShopId              int64
	ShopName            string
	ShopDescription     string
	ShopAssistantName   string
	ShopAssistantNumber string
	PhoneNumber         string
	Email               string
	Image               string
	DateCreated         time.Time
	DateModified        time.Time
	CreatedBy           int
	ModifiedBy          int
	Active              int
}

type UserExtraDetails struct {
	UserDetailsId int64
	Branch        *Branches
	Shop          *Shops
	Nickname      string
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
	Active        int
}

type UserResponseDTO struct {
	StatusCode int
	User       *Users
	StatusDesc string
}
