package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context){
		tokenString := c.GetHeader("Auth")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		fmt.Println("auth midddlware");
		fmt.Println("logged in");
		fmt.Println(tokenString);
	c.Next();
}
func validateToken(token string){


}
	/*
	//var token *jwt.Token;
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": cUser.Username,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")));
	if err != nil {
		c.IndentedJSON(404, gin.H{"message":"could not craete token"});
		return;
	}
    expiration := time.Now().Add(365 * 24 * time.Hour)
    cookie := http.Cookie{Name: "Auth",Value:tokenString,Expires:expiration, Path: "/",Domain: "", Secure: true, HttpOnly: true}
    http.SetCookie(c.Writer, &cookie)

	c.JSON(200, gin.H{"token":"token created"});
	*/