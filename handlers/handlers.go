package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/davidchandra95/go-gin/models"
	"strconv"
	"log"
)

func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, models.Jokes)
}

func LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	jokeidtes, _ := strconv.Atoi(c.Param("jokeID"))
	log.Println(jokeidtes)
	jokes := &models.Jokes
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(*jokes); i++ {
			if (*jokes)[i].ID == jokeid {
				(*jokes)[i].Likes++
			}
		}

		c.JSON(http.StatusOK, &jokes)
	}else{
		c.AbortWithStatus(http.StatusNotFound)
	}
}