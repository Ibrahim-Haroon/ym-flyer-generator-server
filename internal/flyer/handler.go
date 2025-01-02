package flyer

import (
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/flyer/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func GetBackground(c *gin.Context) {
	var createRequest model.CreateRequest

	err := c.BindJSON(&createRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("failed to bind json: %v", err))
	}

	var backgrounds []model.Background
	backgroundPaths, err := CreateBackground(
		createRequest.ColorPalette,
		createRequest.TextModelProviderMeta,
		createRequest.ImageModelProviderMeta,
	)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	for _, backgroundPath := range backgroundPaths {
		image, err := os.Open(backgroundPath)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("failed to open image: %v", err))
		}
		defer image.Close()

		imageData, err := ioutil.ReadAll(image)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("failed to read image: %v", err))
		}

		backgrounds = append(backgrounds, model.Background{
			SavePath: backgroundPath,
			Image:    imageData,
		})
	}

	c.IndentedJSON(http.StatusOK, model.CreateResponse{Backgrounds: backgrounds})
}
