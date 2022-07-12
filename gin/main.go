package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int
	Name string
}

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

func main() {
	g := gin.Default()
	g.GET("/users", listUsers)
	g.GET("/users/:id", getUser)
	g.POST("/users", createUser)
	g.DELETE("/users/:id", deleteUser)
	g.PATCH("/users/:id", updateUserName)
	g.Run(":8080")
}

func listUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user User
	found := false

	for _, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			user = u
			found = true
			break
		}
	}
	if found {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}

func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name != "" {
		u := User{
			ID:   len(users) + 1,
			Name: name,
		}
		users = append(users, u)
		c.JSON(http.StatusCreated, u)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入用户名称",
		})
	}
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	i := -1
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			i = index
			break
		}
	}
	if i >= 0 {
		users = append(users[:i], users[i+1:]...)
		c.JSON(http.StatusNoContent, "")
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}

func updateUserName(c *gin.Context) {
	id := c.Param("id")
	i := -1
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			i = index
			break
		}
	}
	if i >= 0 {
		users[i].Name = c.DefaultPostForm("name", users[i].Name)
		c.JSON(http.StatusOK, users[i])
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}
