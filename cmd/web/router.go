package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/nilsonmart/we-exchange/internal/models"
)

func router() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("ui/pages/*")

	//Route that emit a JWT Token
	r.POST("/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		//check credentials
		//TODO - REAL AUTHENTICATION VALIDATION
		isValid, err := ValidateAccount(email, password)
		if err != nil {
			//c.AbortWithStatus(http.StatusBadRequest)
			//c.Redirect(http.StatusNotFound, "/")
			c.HTML(http.StatusUnauthorized, "login.html", nil)
		}

		if isValid {
			expToken := time.Now().Add(time.Hour * 24).Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": email,
				"exp":   expToken, //Valid for 24h
			})

			tokenString, err := token.SignedString([]byte(getSecretKey()))

			fmt.Printf("tokenstring: %v - error: %v \t", tokenString, err)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error when creating token"})
				return
			}
			c.SetCookie("userid", "id", int(expToken), "/", "localhost", false, true)
			c.SetCookie("token", tokenString, int(expToken), "/", "localhost", false, true)

			//c.JSON(http.StatusOK, gin.H{"token": tokenString})
			c.Redirect(http.StatusOK, "/home")
		} else {
			//c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	})

	//Route requires JWT Token authentication
	r.GET("/home", authMiddleware(), func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "Protected route. Welcome!"})
		c.HTML(http.StatusOK, "home.html", nil)
	})

	//Route to render a template
	r.GET("/", func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.HTML(http.StatusOK, "login.html", nil)
			return
		}
		if tokenString != "" {
			c.HTML(http.StatusOK, "home.html", nil)
		} else {
			c.HTML(http.StatusOK, "login.html", nil)
		}
	})

	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//SCHEMA
	r.GET("/schema", authMiddleware(), func(ctx *gin.Context) {
		models, err := AllSchema()
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, models)
	})

	r.GET("/schema:userid", authMiddleware(), func(ctx *gin.Context) {
		id := ctx.Param("id")

		model, err := GetSchemaByUserID(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, model)
	})

	r.PUT("/schema:id", authMiddleware(), func(ctx *gin.Context) {
		var model models.Schema
		id := ctx.Param("id")

		if err := ctx.BindJSON(&model); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		isUpdated, err := UpdateSchema(id, model)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusAccepted, isUpdated)
	})

	r.POST("/schema", authMiddleware(), func(ctx *gin.Context) {
		var model models.Schema

		if err := ctx.BindJSON(&model); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		isUpdated, err := CreateSchema(model)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusAccepted, isUpdated)
	})

	//REQUEST CHANGE
	r.GET("/requestchange", authMiddleware(), func(ctx *gin.Context) {
		models, err := AllRequestChange()
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, models)
	})

	r.GET("/requestchange:id", authMiddleware(), func(ctx *gin.Context) {
		id := ctx.Param("id")
		model, err := GetRequestChangeByID(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, model)
	})

	r.GET("/requestchange:userid", authMiddleware(), func(ctx *gin.Context) {
		userId := ctx.Param("userid")
		model, err := GetRequestChangeByUserID(userId)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, model)
	})

	r.PUT("/requestchange:id", authMiddleware(), func(ctx *gin.Context) {
		var model models.RequestChange
		id := ctx.Param("id")
		if err := ctx.BindJSON(&model); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		isUpdated, err := UpdateRequestChange(id, model)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusAccepted, isUpdated)

	})

	r.POST("/requestchange", authMiddleware(), func(ctx *gin.Context) {
		var model models.RequestChange
		if err := ctx.BindJSON(&model); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		isUpdated, err := CreateRequestChange(model)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusAccepted, isUpdated)

	})

	r.DELETE("/requestchange:id", authMiddleware(), func(ctx *gin.Context) {
		id := ctx.Param("id")
		isUpdated, err := DeleteRequestChange(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusAccepted, isUpdated)
	})

	return r
}
