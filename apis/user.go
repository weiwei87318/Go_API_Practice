package apis

import (
	model "api_practice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "No Info found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func Store(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Failed Adding info",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Success Adding info",
		"data":    id,
	})
}

func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(int(id))
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Update Successful",
	})
}

func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(int(id))
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Delete failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Delete Successful",
	})
}
