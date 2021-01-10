package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jirayuSai/app/ent/gender"
	"github.com/jirayuSai/app/ent/personal"
	"github.com/jirayuSai/app/ent/title"

	"github.com/gin-gonic/gin"
	"github.com/jirayuSai/app/ent"
	"github.com/jirayuSai/app/ent/customer"
)

// CustomerController defines the struct for the customer controller
type CustomerController struct {
	client *ent.Client
	router gin.IRouter
}
type Customer struct {
	Gender       int
	Personal     int
	Title        int
	Address      string
	CustomerName string
	PhoneNumber  string
}

// CreateCustomer handles POST requests for adding customer entities
// @Summary Create customer
// @Description Create customer
// @ID create-customer
// @Accept   json
// @Produce  json
// @Param customer body Customer true "Customer entity"
// @Success 200 {object} Customer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers [post]
func (ctl *CustomerController) CreateCustomer(c *gin.Context) {
	obj := Customer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "customer binding failed",
		})
		return
	}
	p, err := ctl.client.Personal.
		Query().
		Where(personal.IDEQ(obj.Personal)).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "personal not found",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(obj.Gender)).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "gender not found",
		})
		return
	}

	t, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(obj.Title)).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "title not found",
		})
		return
	}

	cm, err := ctl.client.Customer.
		Create().
		SetPersonals(p).
		SetGenders(g).
		SetTitles(t).
		SetAddress(obj.Address).
		SetCustomerName(obj.CustomerName).
		SetPhoneNumber(obj.PhoneNumber).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cm)
}

// GetCustomer handles GET requests to retrieve a customer entity
// @Summary Get a customer entity by ID
// @Description get customer by ID
// @ID get-customer
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers/{id} [get]
func (ctl *CustomerController) GetCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cm, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cm)
}

// ListCustomer handles request to get a list of customer entities
// @Summary List customer entities
// @Description list customer entities
// @ID list-customer
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers [get]
func (ctl *CustomerController) ListCustomer(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	customers, err := ctl.client.Customer.
		Query().
		WithGenders().
		WithPersonals().
		WithTitles().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, customers)
}

// DeleteCustomer handles DELETE requests to delete a customer entity
// @Summary Delete a customer entity by ID
// @Description get customer by ID
// @ID delete-customer
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers/{id} [delete]
func (ctl *CustomerController) DeleteCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Customer.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateCustomer handles PUT requests to update a customer entity
// @Summary Update a customer entity by ID
// @Description update customer by ID
// @ID update-customer
// @Accept   json
// @Produce  json
// @Param id path int true "Customer ID"
// @Param customer body ent.Customer true "Customer entity"
// @Success 200 {object} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers/{id} [put]
func (ctl *CustomerController) UpdateCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Customer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "customer binding failed",
		})
		return
	}
	obj.ID = int(id)
	cm, err := ctl.client.Customer.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cm)
}

// NewCustomerController creates and registers handles for the customer controller
func NewCustomerController(router gin.IRouter, client *ent.Client) *CustomerController {
	cmc := &CustomerController{
		client: client,
		router: router,
	}
	cmc.register()
	return cmc
}

// InitCustomerController registers routes to the main engine
func (ctl *CustomerController) register() {
	customers := ctl.router.Group("/customers")

	customers.GET("", ctl.ListCustomer)

	// CRUD
	customers.POST("", ctl.CreateCustomer)
	customers.GET(":id", ctl.GetCustomer)
	customers.PUT(":id", ctl.UpdateCustomer)
	customers.DELETE(":id", ctl.DeleteCustomer)
}
