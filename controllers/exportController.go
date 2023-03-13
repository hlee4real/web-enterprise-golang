package controllers

import (
	"archive/zip"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"web-enterprise-backend/models"
)

func fileSize(filePath string) int64 {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return info.Size()
}

func ExportIntoZip() gin.HandlerFunc {
	return func(c *gin.Context) {
		filePaths := strings.Split(c.Query("files"), ",")
		tempDir, err := os.MkdirTemp("", "file-download")
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		defer os.RemoveAll(tempDir)
		zipPath := filepath.Join(tempDir, "download.zip")
		zipFile, err := os.Create(zipPath)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		defer zipFile.Close()
		zipWriter := zip.NewWriter(zipFile)
		defer zipWriter.Close()
		for _, filePath := range filePaths {
			if err := addFileToZip(zipWriter, filePath); err != nil {
				c.JSON(500, gin.H{"message": err.Error()})
				return
			}
		}
		c.Header("Content-Type", "application/zip")
		c.Header("Content-Disposition", "attachment; filename=download.zip")
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Length", strconv.FormatInt(fileSize(zipPath), 10))
		c.File(zipPath)
	}
}

// addFileToZip adds a file to a zip archive
func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	zipFile, err := zipWriter.Create(filepath.Base(filePath))
	if err != nil {
		return err
	}

	_, err = io.Copy(zipFile, file)
	return err
}

func writeCSVRows(cursor *mongo.Cursor, csvWriter *csv.Writer) error {
	for cursor.Next(context.Background()) {
		var idea models.IdeasModel
		if err := cursor.Decode(&idea); err != nil {
			return err
		}
		row := []string{
			idea.ID.Hex(),
			idea.Image,
			idea.Title,
			idea.Slug,
			idea.Filename,
			idea.Department,
			idea.Content,
			idea.Category,
			strconv.Itoa(idea.Views),
			strconv.Itoa(idea.UpVote),
			strconv.Itoa(idea.DownVote),
			idea.CreatedAt.String(),
			idea.UpdatedAt.String(),
			idea.Username,
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	return cursor.Err()
}

func ExportIntoCSV() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		csvPath := c.Query("csvPath")
		// create CSV file
		csvFile, err := os.Create(csvPath)
		if err != nil {
			fmt.Print("hihi", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer csvFile.Close()

		// set up CSV writer
		csvWriter := csv.NewWriter(csvFile)
		defer csvWriter.Flush()
		cursor, err := ideaCollection.Find(ctx, bson.D{})
		if err != nil {
			fmt.Print("haha", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Fatal(err)
		}
		defer cursor.Close(ctx)
		if err := csvWriter.Write([]string{"_id", "image", "title", "slug", "filename", "department", "content", "category", "views", "up_vote", "down_vote", "created_at", "updated_at", "username"}); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// write data rows to CSV file
		if err := writeCSVRows(cursor, csvWriter); err != nil {
			fmt.Print("hehe", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// set response headers for file download
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+csvPath)
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Length", strconv.FormatInt(fileSize(csvPath), 10))

		// stream file to response body
		c.File(csvPath)
	}
}
