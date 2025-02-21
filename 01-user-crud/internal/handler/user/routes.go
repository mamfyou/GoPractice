package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mamfyou/internal/service/user"
)

func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/user")

	group.POST("/", insert)
	group.PUT("/", update)
	group.GET("/:id", get)
	group.DELETE("/:id", delete_)
}

func insert(c *gin.Context) {
	req := new(user.User)

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := user.Insert(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func get(c *gin.Context) {
	req := new(user.User)

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_param, ok := extractIDParam(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	user, err := user.Get(id_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func update(c *gin.Context) {
	req := new(user.User)

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Update(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "successful"})

}

func delete_(c *gin.Context) {
	req := new(user.User)

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_param, ok := extractIDParam(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = user.Delete(id_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Successful"})

}

/* ---------- helpers ---------- */

func extractIDParam(c *gin.Context) (id int64, ok bool) {
	value := c.Param("id")

	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return id, false
	}

	return id, true
}
