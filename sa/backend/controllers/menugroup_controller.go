package controllers

import (
		"context"
		"strconv"
		"fmt"
		"github.com/pomea/app/ent/menugroup"
		"github.com/gin-gonic/gin"
		"github.com/pomea/app/ent"
)

// MenugroupController is ...
type MenugroupController struct { 
	client *ent.Client 
	router gin.IRouter
}

// CreateMenugroup handles POST requests for adding menugroup entities
// @Summary Create menugroup
// @Description Create menugroup
// @ID create-menugroup
// @Accept   json
// @Produce  json
// @Param menugroup body ent.Menugroup true "Menugroup entity"
// @Success 200 {object} ent.Menugroup
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menugroups [post]
func (ctl *MenugroupController) CreateMenugroup(c *gin.Context) {
	obj := ent.Menugroup{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "menugroup binding failed",
		})
		return
	}

	g, err := ctl.client.Menugroup.
		Create().
		SetGroup(obj.Group).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, g)
}

// GetMenugroup handles GET requests to retrieve a menugroup entity
// @Summary Get a menugroup entity by ID
// @Description get menugroup by ID
// @ID get-menugroup
// @Produce  json
// @Param id path int true "Menugroup ID"
// @Success 200 {object} ent.Menugroup
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menugroups/{id} [get]
func (ctl *MenugroupController) GetMenugroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	g, err := ctl.client.Menugroup.
		Query().
		Where(menugroup.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, g)
}

// ListMenugroup handles request to get a list of menugroup entities
// @Summary List menugroup entities
// @Description list menugroup entities
// @ID list-menugroup
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Menugroup
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menugroups [get]
func (ctl *MenugroupController) ListMenugroup(c *gin.Context) {
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

	menugroup, err := ctl.client.Menugroup.
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

	c.JSON(200, menugroup)
}

// DeleteMenugroup handles DELETE requests to delete a menugroup entity
// @Summary Delete a menugroup entity by ID
// @Description get menugroup by ID
// @ID delete-menugroup
// @Produce  json
// @Param id path int true "Menugroup ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menugroups/{id} [delete]
func (ctl *MenugroupController) DeleteMenugroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Menugroup.
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

// UpdateMenugroup handles PUT requests to update a menugroup entity
// @Summary Update a menugroup entity by ID
// @Description update menugroup by ID
// @ID update-menugroup
// @Accept   json
// @Produce  json
// @Param id path int true "Menugroup ID"
// @Param menugroup body ent.Menugroup true "Menugroup entity"
// @Success 200 {object} ent.Menugroup
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /menugroups/{id} [put]
func (ctl *MenugroupController) UpdateMenugroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Menugroup{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "menugroup binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	r, err := ctl.client.Menugroup.
		UpdateOneID(int(id)).
		SetGroup(obj.Group).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, r)
}

// NewMenugroupController creates and registers handles for the menugroup controller
func NewMenugroupController(router gin.IRouter, client *ent.Client) *MenugroupController {
	rc := &MenugroupController{
		client: client,
		router: router,
	}

	rc.register()
	return rc

}

func (ctl *MenugroupController) register() {
	menugroups := ctl.router.Group("/menugroups")
	
	menugroups.GET("", ctl.ListMenugroup)

	//CRUD
	menugroups.POST("", ctl.CreateMenugroup)
	menugroups.GET(":id", ctl.GetMenugroup)
	menugroups.PUT(":id", ctl.UpdateMenugroup)
	menugroups.DELETE(":id", ctl.DeleteMenugroup)
}