package database

var Products []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func init() {
	prod1 := Product{ID: 1, Title: "orange", Description: "my father's fav fruit", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"}
	prod2 := Product{ID: 2, Title: "apple", Description: "tasty to eat while fever", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"}
	prod3 := Product{ID: 3, Title: "Watermelon", Description: "Tasty & Juicy", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"}

	Products = append(Products, prod1, prod2, prod3)
}
