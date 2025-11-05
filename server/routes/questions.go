package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// validate will ensure that the data we receive is correct
var validate = validator.New()
var questionCollection *mongo.Collection = OpenCollection(Client, "questions")

// AddQuestion Add a question post to the database
func AddQuestion(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var post models.Post

	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(post)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	// Create and add a question
	post.ID = primitive.NewObjectID()
	result, insertErr := questionCollection.InsertOne(ctx, post)

	if insertErr != nil {
		msg := fmt.Sprintf("Error inserting new question: %s", insertErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, result)
}
