package middleware

import (
	"fmt"
	"net/http"

	helper "github.com/BrainStation-23/dsdoc-backend/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context)  {
		
		clientToken,error := c.Cookie("token")
		
		if error != nil || clientToken=="" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	
}