package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/isjhar/iet/internal/domain/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ValidateExt(file *multipart.FileHeader, exts entities.FileExtensions) error {
	fileExt := filepath.Ext(file.Filename)
	for _, ext := range exts {
		if "."+ext != fileExt {
			return entities.InvalidParams
		}
	}

	return entities.InvalidParams
}

// size int byte
func ValidateSize(file *multipart.FileHeader, size int64) error {
	if file.Size > size {
		return entities.InvalidParams
	}
	return nil
}

type FormFileProcessor struct {
	Context echo.Context
}

type ValidateThenExtractParams struct {
	Name           string
	MaxSize        int64
	FileExtensions entities.FileExtensions
}

func (f *FormFileProcessor) ValidateThenExtract(params ValidateThenExtractParams) ([]byte, error) {
	file, err := f.Context.FormFile(params.Name)
	if err != nil {
		return nil, entities.InvalidParams
	}

	if err := ValidateSize(file, params.MaxSize); err != nil {
		return nil, entities.FileSizeReachLimit(params.MaxSize)
	}

	if err := ValidateExt(file, params.FileExtensions); err != nil {
		return nil, entities.InvalidParams
	}

	fileReader, err := file.Open()
	if err != nil {
		return nil, entities.InternalServerError
	}
	defer fileReader.Close()

	fileBuffer := new(bytes.Buffer)
	_, err = io.Copy(fileBuffer, fileReader)
	if err != nil {
		return nil, entities.InternalServerError
	}
	return fileBuffer.Bytes(), nil
}
