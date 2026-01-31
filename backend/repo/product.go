package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Delete(id int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	products []*Product
}

func NewProductRepo() *productRepo {
	repo := &productRepo{}
	GenerateInitialProduct(repo)

	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.products) + 1
	r.products = append(r.products, &p)
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.products, nil

}

func (r *productRepo) Delete(id int) error {
	var temp []*Product
	for _, product := range r.products {
		if product.ID != id {
			temp = append(temp, product)
		}
	}
	r.products = temp
	return nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	for _, product := range r.products {
		if product.ID == p.ID {
			*product = p
			return product, nil
		}
	}
	return &p, nil
}

func GenerateInitialProduct(r *productRepo) {
	prod1 := Product{ID: 1, Title: "orange",
		Description: "my father's fav fruit",
		Price:       129.99,
		ImgURL:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	prod2 := Product{ID: 2,
		Title:       "apple",
		Description: "tasty to eat while fever",
		Price:       199.99,
		ImgURL:      "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}
	prod3 := Product{ID: 3,
		Title:       "Watermelon",
		Description: "Tasty & Juicy",
		Price:       89.50,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}

	r.products = append(r.products, &prod1, &prod2, &prod3)
}
