package repo

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
	"fmt"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repo {
	return repo{
		DB: db,
	}
}

func (r *repo) Create(product *dto.Product) error {
	if err := r.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) GetByUserID(userID int, limit, offset int) ([]*entity.Product, error) {
	var products []*entity.Product
	tx := r.DB.Where("user_id=?", userID).Limit(limit).Offset(offset).Find(&products)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repo) Update(userID int, product *entity.Product) (*entity.Product, error) {
	var output entity.Product
	//if userID != product.UserId {
	//	return nil, errors.New("invalid user")
	//}
	fmt.Println(product.Id)
	if err := r.DB.Where("id = ?", product.Id).First(&output).Error; err != nil {
		return nil, err
	}
	r.DB.Model(&output).Updates(&product)
	fmt.Println(output)
	return &output, nil
}
