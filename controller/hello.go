package controller

import (
	"database/sql"
	"my-gin-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{ // 数据部分
			"code":    200,       // 状态码
			"message": "success", // 消息
			"data": gin.H{
				"greeting": "Hello, World!",
			},
		},
	})
}

func GetUsers(c *gin.Context) {
	rows, err := model.DB.Query("SELECT username, id, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	err := model.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "查询不到该用户"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func SearchArea(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	err := model.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "查询不到该用户"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
