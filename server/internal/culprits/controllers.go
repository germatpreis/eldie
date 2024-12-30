package culprits

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCulprits(c *gin.Context) {
	culprits := []Culprit{
		{
			Food: Food{
				Name:  "Coffee",
				Alias: nil,
			},
			Likelihood: 10,
		},
		{
			Food: Food{
				Name:  "Tomatoes",
				Alias: nil,
			},
			Likelihood: 8,
		},
	}

	c.IndentedJSON(http.StatusOK, culprits)
}
