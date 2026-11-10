package models

type AdvertisementsList struct {
	HasNext    bool            `json:"has_next"`
	Items      []Advertisement `json:"items"`
	NextCursor string          `json:"next_cursor"`
	Total      int             `json:"total"`
}

type Advertisement struct {
	CreatedAt   string  `json:"created_at"`
	Description string  `json:"description"`
	ID          string  `json:"id"`
	Photos      []Photo `json:"photos"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Title       string  `json:"title"`
	UpdatedAt   string  `json:"updated_at"`
	User        User    `json:"user"`
	UserID      string  `json:"user_id"`
}

type Photo struct {
	AdvertisementID string `json:"advertisement_id"`
	CreatedAt       string `json:"created_at"`
	ID              string `json:"id"`
	SortOrder       int    `json:"sort_order"`
	URL             string `json:"url"`
}

type User struct {
	BirthDate string `json:"birth_date"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	ID        string `json:"id"`
	LastName  string `json:"last_name"`
	UpdatedAt string `json:"updated_at"`
}
