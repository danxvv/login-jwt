package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"login-user/internal/controller/http/middleware"
	"login-user/internal/entity"
	"login-user/internal/usecase"
	"login-user/pkg/jwtHelper"
)

type userRoutes struct {
	u usecase.User
	v *validator.Validate
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.User, v *validator.Validate) {
	user := userRoutes{u: u, v: v}
	handler.POST("/register", user.registerUser)
	handler.POST("/login", user.loginUser)
	authorized := handler.Group("/")
	authorized.Use(middleware.JWTMiddleware())
	authorized.GET("/user", user.getUser)

}

type userRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" validate:"password"`
	Phone    string `json:"phone" binding:"required" validate:"number,len=10"`
	Email    string `json:"email" binding:"required" validate:"email"`
}

func (u *userRoutes) registerUser(c *gin.Context) {
	var user userRegister
	trans, _ := uni.GetTranslator("en")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := u.v.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			println(e.Translate(trans))
			c.JSON(400, gin.H{"error": e.Translate(trans)})
			return
		}
	}
	userEntity := entity.User{
		Email:    user.Email,
		Phone:    user.Phone,
		Username: user.Username,
		Password: user.Password,
	}

	err = u.u.RegisterUser(userEntity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

type userLoginIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" validate:"password"`
}

type userLoginOut struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *userRoutes) loginUser(c *gin.Context) {
	var user userLoginIn
	trans, _ := uni.GetTranslator("en")

	if err := c.ShouldBindJSON(&user); err != nil {
		//custom error message
		errMsg := err.(validator.ValidationErrors)
		for _, e := range errMsg {
			key := e.Field()
			fmt.Println(key)
			switch key {
			case "Username":
				c.JSON(400, gin.H{"error": "Username is required"})
				return
			case "Password":
				c.JSON(400, gin.H{"error": "Password is required"})
				return
			}
		}
	}

	err := u.v.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			c.JSON(400, gin.H{"error": e.Translate(trans)})
			return
		}
	}

	userEntity, err := u.u.LoginUser(user.Username, user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	userOut := userLoginOut{
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Phone:     userEntity.Phone,
		ID:        userEntity.ID.String(),
		CreatedAt: userEntity.CreatedAt.String(),
		UpdatedAt: userEntity.UpdatedAt.String(),
	}

	userToken, _ := jwtHelper.GenerateToken(userEntity.ID.String())
	c.JSON(200, gin.H{"message": "User logged in successfully", "data": userOut, "token": userToken})
}

func (u *userRoutes) getUser(c *gin.Context) {
	userID := c.GetString("userID")
	userEntity, err := u.u.GetUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	userOut := userLoginOut{
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Phone:     userEntity.Phone,
		ID:        userEntity.ID.String(),
		CreatedAt: userEntity.CreatedAt.String(),
		UpdatedAt: userEntity.UpdatedAt.String(),
	}

	c.JSON(200, gin.H{"message": "User data", "data": userOut})
}
