package model

// Products -
type Products struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Price string `json:"price" gorm:"type:varchar(100)"`
	Point string `json:"point" gorm:"type:varchar(100)"`
	Pic   string `json:"pic" gorm:"longtext"`
	Link  string `json:"link" gorm:"type:varchar(255)"`
}

// Carts -
type Carts struct {
	ID          int    `json:"id" gorm:"column:id"`
	Username    string `json:"username" gorm:"type:varchar(50)"`
	ProductName string `json:"product_name" gorm:"type:varchar(255)"`
	Qty         int    `json:"qty"`
	Price       string `json:"price" gorm:"type:varchar(100)"`
}

// Orders -
type Orders struct {
	ID         int    `json:"id" gorm:"column:id"`
	Username   string `json:"username" gorm:"type:varchar(50)"`
	Detail     string `json:"detail" gorm:"text;column:detail"`
	TotalPrice int    `json:"total_price" gorm:"int(11);column:total_price"`
}
