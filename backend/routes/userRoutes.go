package routes

import (
	"context"
	"fmt"
	"katsuebackend/backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"github.com/golang-jwt/jwt/v5"
	"katsuebackend/backend/database"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"katsuetwo/backend/middleware"
)

var users[] models.User = []models.User{
	{Name: "joe",Username: "joe123",Email: "joe@gmail.com", Password: "asdf"},
	{Name: "sal",Username: "sal567",Email: "sal@gmail.com", Password: "lakers"},
	{Name: "bob",Username: "bobbob", Email: "bob@gmail.com", Password: "1234"},
}
func test(c *gin.Context){
	c.IndentedJSON(200, gin.H{"message":"docker wokring"});
}
func getAllUsers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, users);
}
type CreateHandler struct {
	Name string `json:"name"`;
	Username string `json:"username"`;
	Email string `json:"email"`;
	Password string `json:"password"`;
}
func createUser(c *gin.Context){
	var cUser CreateHandler;
	var err error = c.BindJSON(&cUser);
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message":"error...",
		})
		return;
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(cUser.Password), bcrypt.DefaultCost);
	if err != nil {
		c.IndentedJSON(404, gin.H{"message":"could not create account"});
	}
	var nwUser models.User = models.User{
		Name: cUser.Name,
		Username: cUser.Username,
		Email: cUser.Email,
		Password: string(hash),
	}
	userCollection := database.MongoKatuseDb.DB.Collection("users");
	userResult, err := userCollection.InsertOne(context.TODO(), nwUser);  //inserting documents
	if err != nil {
		c.IndentedJSON(404, gin.H{"message":"could not create account"});
		return;
	}
	fmt.Println(userResult);
	cUser.hidePas();
	c.IndentedJSON(http.StatusAccepted, cUser);
}
type Testing struct{
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`  //bson
	Name string `json:"name"`
	Username string `json:"username"`;
	Email string `json:"email"`
	Password string `json:"password"`
}

func (d *Testing)hidePass(){
	d.Password = "";
}
func (u *CreateHandler)hidePas(){
	u.Password = "";
}

func getUserFromDb(typ string, name string) (*Testing){
	var res Testing = Testing{}
	var result *Testing = &res;
	userCollection := database.MongoKatuseDb.DB.Collection("users");
	
	filter := bson.D{{Key: typ, Value: name}};

	var err error = userCollection.FindOne(context.TODO(), filter).Decode(result);
	if err != nil {
		return result;
	}
	fmt.Println(result);
	return result
}

func getOneUser(c *gin.Context){
	var name string = c.Param("name");
	var currUser *Testing = getUserFromDb("name", name);
	if currUser.Email == "" {
		c.IndentedJSON(404, gin.H{"message":"user not found"});
	}
	c.IndentedJSON(200, currUser);
}


func logIn(c *gin.Context){
	var adduser CreateHandler;
	var cUser *CreateHandler = &adduser;

	var err error = c.BindJSON(&cUser); // binds user data struct to json data
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message":"error...",
		})
		return;
	}  //cUser is from request.body
	var currUser *Testing = getUserFromDb("email", cUser.Email);
	var hasErr error = bcrypt.CompareHashAndPassword([]byte(currUser.Password),[]byte(cUser.Password));
	if hasErr != nil {
		c.IndentedJSON(404, gin.H{"message":"wrong password"});
		return;
	}
	currUser.hidePass();
	c.IndentedJSON(200, currUser);

}
/*
func validate(c *gin.Context){
	c.IndentedJSON(200, gin.H{"message":"validate"})
}
*/
func UserRoutes(g *gin.RouterGroup){
	g.GET("/getAllUsers", getAllUsers);
	g.POST("/createUser", createUser);
	g.GET("/getOneUser/:name", getOneUser);
	g.POST("/loginUser",logIn);
	g.GET("/test", test);
	//g.GET("/validate", middleware.RequireAuth ,validate);

}