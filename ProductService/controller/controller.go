package controller

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type controller struct {
	Service service.Service
}

func New(svc service.Service) *controller {
	return &controller{Service: svc}
}

func (ctrl *controller) Create(c *gin.Context) {
	var input dto.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Controller - Create: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Hard coded
	input.UserId = 1
	input.CreatedTime = time.Now()
	input.ExpiredTime = time.Now().Add(time.Hour * 24)

	if err := ctrl.Service.Create(&input); err != nil {
		log.Printf("Controller - Create: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (ctrl *controller) GetByUserID(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
	}
	userID := 1
	products, err := ctrl.Service.GetByUserID(userID, 2, pageNum)
	c.JSON(http.StatusOK, products)
}

func (ctrl *controller) GetList(c *gin.Context) {

}
func (ctrl *controller) Update(c *gin.Context) {
	var input dto.ProductUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Controller - Update: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	product, err := ctrl.Service.Update(&input)
	if err != nil {
		log.Printf("Controller - Update: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ctrl *controller) Delete(c *gin.Context) {

}
