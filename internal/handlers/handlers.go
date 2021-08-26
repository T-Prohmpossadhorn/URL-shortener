package handlers

import (
	"fmt"
	"net/http"
	"url-shortener/config"
	"url-shortener/internal/store"

	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hey Go URL Shortener !",
	})
}

var storage *store.Service
var configuration config.Config

func Setupstore(store *store.Service, cfg config.Config) {
	storage = store
	configuration = cfg
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest store.PostItem
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl, err := (*storage).Save(creationRequest.URL, creationRequest.Expires)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "short url created successfully",
		"short_url": configuration.Options.Schema + configuration.Options.Prefix + shortUrl,
	})
}

func GetShortUrlInfo(c *gin.Context) {
	shorturl := c.Param("shortUrl")

	info, err := (*storage).LoadInfo(shorturl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, info)
}

func DeleteShortUrl(c *gin.Context) {
	shorturl := c.Param("shortUrl")

	err := (*storage).Delete(shorturl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "short url deleted",
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shorturl := c.Param("shortUrl")

	url, err := (*storage).Load(shorturl)
	if err != nil {
		switch err {
		case fmt.Errorf("url not found"):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		case fmt.Errorf("url expired"):
			c.JSON(http.StatusGone, gin.H{
				"message": err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
	}

	c.Redirect(http.StatusFound, url)
}
