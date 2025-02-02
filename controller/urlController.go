package controller

import (
	"fmt"
	"net/http"
	"time"
	"url-shortner/constant"
	"url-shortner/database"
	"url-shortner/helper"
	"url-shortner/types"

	"github.com/gin-gonic/gin"
)

func ShortTheUrl(c *gin.Context) {
	var shortUrlBody types.ShortUrlBody
	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Bind error: " + constant.BindError})
		return
	}
	code := helper.GenRandomString(6)
	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode != "" {
		// url present:
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "This code is already in use"})
		return
	}

	var url types.UrlDb
	url.CreatedAt = time.Now().Unix()
	url.ExpiredAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LongUrl
	url.ShortUrl = constant.BaseUrl + code

	res, err := database.Mgr.Insert(url, constant.UrlCollection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":     false,
		"data":      res,
		"short_url": url.ShortUrl,
	})
}

func RedirectUrl(c *gin.Context) {
	code := c.Param("code")
	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "no url found to this code",
		})
	}

	fmt.Println(record.LongUrl)
	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)

}
