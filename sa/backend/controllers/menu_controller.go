package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pomea/app/ent"
	"github.com/pomea/app/ent/menu"
	"github.com/pomea/app/ent/menugroup"
	"github.com/pomea/app/ent/menutype"
	"github.com/pomea/app/ent/user"
)

// MenuController is ...
type MenuController struct {
	client *ent.Client
	router gin.IRouter
}

// Menu defines the struct for the menu ...
type Menu struct {
	Addedtime  string
	Calories   int
	GroupID    int
	Ingredient string
	Menuname   string
	TypeID     int
	UserID     int
}

// CreateMenu handles POST requests for adding menu entities
// @Summary Create menu
// @Description Create menu
// @ID create-menu
// @Accept   json
// @Produce  json
// @Param menu body Menu true "Menu entity"
// @Success 200 {object} ent.Menu
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menus [post]
func (ctl *MenuController) CreateMenu(c *gin.Context) {
	obj := Menu{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "menu binding failed",
		})
		return

	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "User not found",
		})
		return
	}

	t, err := ctl.client.Menutype.
		Query().
		Where(menutype.IDEQ(int(obj.TypeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Type not found",
		})
		return
	}

	g, err := ctl.client.Menugroup.
		Query().
		Where(menugroup.IDEQ(int(obj.GroupID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "group not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Addedtime)

	fm, err := ctl.client.Menu.
		Create().
		SetMenuname(obj.Menuname).
		SetType(t).
		SetGroup(g).
		SetCalories(obj.Calories).
		SetOwner(u).
		SetIngredient(obj.Ingredient).
		SetAddedTime(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, fm)
}

// GetMenu handles GET requests to retrieve a menu entity
// @Summary Get a menu entity by ID
// @Description get menu by ID
// @ID get-menu
// @Produce  json
// @Param id path int true "Menu ID"
// @Success 200 {object} ent.Menu
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menus/{id} [get]
func (ctl *MenuController) GetMenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := ctl.client.Menu.
		Query().
		Where(menu.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListMenu handles request to get a list of menu entities
// @Summary List menu entities
// @Description list menu entities
// @ID list-menu
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Menu
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menus [get]
func (ctl *MenuController) ListMenu(c *gin.Context) {
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

	menus, err := ctl.client.Menu.
		Query().
		WithOwner().
		WithType().
		WithGroup().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, menus)
}

// DeleteMenu handles DELETE requests to delete a menu entity
// @Summary Delete a menu entity by ID
// @Description get menu by ID
// @ID delete-menu
// @Produce  json
// @Param id path int true "Menu ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menus/{id} [delete]
func (ctl *MenuController) DeleteMenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Menu.
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

// NewMenuController creates and registers handles for the menu controller
func NewMenuController(router gin.IRouter, client *ent.Client) *MenuController {
	fmc := &MenuController{
		client: client,
		router: router,
	}

	fmc.register()

	return fmc

}

// InitMenuController registers routes to the main engine
func (ctl *MenuController) register() {
	menus := ctl.router.Group("/menus")

	menus.GET("", ctl.ListMenu)

	//CRUD
	menus.GET(":id", ctl.GetMenu)
	menus.POST("", ctl.CreateMenu)
	menus.DELETE(":id", ctl.DeleteMenu)

}
