package controller

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/settings"
	"github.com/scafol/KP-Backend/utils"
)

type UploaderController struct{}

// Uploader - do upload a file to specific target
// POST /upload
func (uploaderController UploaderController) Uploader(c echo.Context) error {
	file, err := c.FormFile("file")

	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	copyingError := CopyingFile(file)
	if copyingError != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusCreated, res)
	}

	res := utils.MakeResponse("Success", "Files successfully uploaded", nil)
	return c.JSON(http.StatusCreated, res)
}

// Uploaders - do upload a file to specific target
// POST /uploads
func (uploaderController UploaderController) Uploaders(c echo.Context) error {
	form, err := c.MultipartForm()

	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	files := form.File["files"]

	for _, file := range files {
		err := CopyingFile(file)
		if err != nil {
			log.Printf("Upload file failed with %s", err.Error())
		}
	}

	res := utils.MakeResponse("Success", "Files successfully uploaded", nil)
	return c.JSON(http.StatusCreated, res)
}

// CopyingFile - copy file to specific location
func CopyingFile(file *multipart.FileHeader) error {
	// File source
	src, err := file.Open()
	if err != nil {
		return err
	}
	// close connection to file
	defer src.Close()

	currentDirectory, _ := os.Getwd()

	// Copying file to destination location
	target := currentDirectory + "/" + settings.GoDotEnvVariable("DIR_PATH") + "/" + file.Filename
	dst, err := os.Create(target)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
