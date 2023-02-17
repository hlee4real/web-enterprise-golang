package helper

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

var extWhiteList = map[string]bool{".docx": true, ".pdf": true, ".jpg": true, ".jpeg": true, ".png": true, ".xlsx": true}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := ValidateFileExt(filepath.Ext(file.Filename)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", file.Filename)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"filename": file.Filename, "message": "file uploaded successfully"})
}

func ValidateFileExt(ext string) error {
	if _, exist := extWhiteList[strings.ToLower(ext)]; !exist {
		return errors.New("file extension not allowed")
	}
	return nil
}
