package dtos

type AddProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
}
