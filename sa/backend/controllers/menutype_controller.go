package controllers

import (
	"context"
	"fmt"
	"strconv"
	"github.com/pomea/app/ent"
	"github.com/pomea/app/ent/menutype"
	"github.com/gin-gonic/gin"
 )

 // MenutypeController is ...
type MenutypeController struct { 
	client *ent.Client 
	router gin.IRouter
}

// CreateMenutype handles POST requests for adding menutype entities
// @Summary Create menutype
// @Description Create menutype
// @ID create-menutype
// @Accept   json
// @Produce  json
// @Param menutype body ent.Menutype true "Menutype entity"
// @Success 200 {object} ent.Menutype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menutypes [post]
func (ctl *MenutypeController) CreateMenutype(c *gin.Context) {
	obj := ent.Menutype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "menutype binding failed",
		})
		return
	}

	t, err := ctl.client.Menutype.
		Create().
		SetType(obj.Type).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetMenutype handles GET requests to retrieve a menutype entity
// @Summary Get a menutype entity by ID
// @Description get menutype by ID
// @ID get-menutype
// @Produce  json
// @Param id path int true "Menutype ID"
// @Success 200 {object} ent.Menutype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menutypes/{id} [get]
func (ctl *MenutypeController) GetMenutype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	t, err := ctl.client.Menutype.
		Query().
		Where(menutype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListMenutype handles request to get a list of menutype entities
// @Summary List menutype entities
// @Description list menutype entities
// @ID list-menutype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Menutype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menutypes [get]
func (ctl *MenutypeController) ListMenutype(c *gin.Context) {
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

	menutypes, err := ctl.client.Menutype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, menutypes)
}

// DeleteMenutype handles DELETE requests to delete a menutype entity
// @Summary Delete a menutype entity by ID
// @Description get menutype by ID
// @ID delete-menutype
// @Produce  json
// @Param id path int true "Menutype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menutypes/{id} [delete]
func (ctl *MenutypeController) DeleteMenutype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Menutype.
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

// UpdateMenutype handles PUT requests to update a menutype entity
// @Summary Update a menutype entity by ID
// @Description update menutype by ID
// @ID update-menutype
// @Accept   json
// @Produce  json
// @Param id path int true "Menutype ID"
// @Param menutype body ent.Menutype true "Menutype entity"
// @Success 200 {object} ent.Menutype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menutypes/{id} [put]
func (ctl *MenutypeController) UpdateMenutype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Menutype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "menutype binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	r, err := ctl.client.Menutype.
		UpdateOneID(int(id)).
		SetType(obj.Type).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, r)
}

// NewMenutypeController creates and registers handles for the menutype controller
func NewMenutypeController(router gin.IRouter, client *ent.Client) *MenutypeController {
	rc := &MenutypeController{
		client: client,
		router: router,
	}

	rc.register()
	return rc

}

func (ctl *MenutypeController) register() {
	menutypes := ctl.router.Group("/menutypes")
	
	menutypes.GET("", ctl.ListMenutype)

	//CRUD
	menutypes.POST("", ctl.CreateMenutype)
	menutypes.GET(":id", ctl.GetMenutype)
	menutypes.PUT(":id", ctl.UpdateMenutype)
	menutypes.DELETE(":id", ctl.DeleteMenutype)
}