package pkg

import (
	"bytes"
	"os"
	"strings"
	"text/template"
	"unicode"
)

type ScriptGenerator struct {
}

func (i *ScriptGenerator) GenerateGet(modelName string) error {
	var err error

	rootPath, err := FindProjectRoot()
	if err != nil {
		return err
	}

	snackCaseModelName := i.converToSnackCase(modelName)

	data := map[string]string{
		"StructName":          modelName,
		"StructNames":         i.convertToPrulal(modelName),
		"SnackCaseModelName":  snackCaseModelName,
		"CamelCaseModelName":  i.convertToCamelCase(modelName),
		"KebabCaseModelNames": i.convertToPrulal(i.convertToKebabCase(modelName)),
	}

	templateWriter := templateWriter{
		Data:                data,
		RootPath:            rootPath,
		SnackCaseModelName:  snackCaseModelName,
		SnackCaseModelNames: i.convertToPrulal(snackCaseModelName),
	}

	return templateWriter.write()
}

func (i *ScriptGenerator) convertToPrulal(modelName string) string {
	if modelName[len(modelName)-1] == 'y' {
		return modelName[:len(modelName)-1] + "ies"
	}

	return modelName + "s"
}

func (i *ScriptGenerator) convertToCamelCase(modelName string) string {

	return string(unicode.ToLower(rune(modelName[0]))) + modelName[1:]
}

func (i *ScriptGenerator) converToSnackCase(str string) string {
	// Convert the string to snake case
	snakeCase := ""
	for i, c := range str {
		if i > 0 && c >= 'A' && c <= 'Z' {
			snakeCase += "_"
		}
		snakeCase += string(unicode.ToLower(c))
	}
	return snakeCase
}

func (i *ScriptGenerator) convertToKebabCase(str string) string {
	// Convert the string to kebab case
	kebabCase := ""
	for i, c := range str {
		if i > 0 && c >= 'A' && c <= 'Z' {
			kebabCase += "-"
		}
		kebabCase += string(unicode.ToLower(c))
	}
	return kebabCase
}

type templateWriter struct {
	Data                map[string]string
	RootPath            string
	SnackCaseModelName  string
	SnackCaseModelNames string
}

func (te *templateWriter) write() error {
	err := te.writeTemplateToFile(entityScriptTemplate,
		te.RootPath+"/internal/domain/entities/"+te.SnackCaseModelName+".go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFile(repositoryInterfaceScriptTemplate,
		te.RootPath+"/internal/domain/repositories/"+te.SnackCaseModelName+"_repository.go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFile(modelTemplate,
		te.RootPath+"/internal/data/models/"+te.SnackCaseModelName+".go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFile(repositoryImplScriptTemplate,
		te.RootPath+"/internal/data/repositories/"+te.SnackCaseModelName+"_repository.go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFile(usecaseScriptTemplate,
		te.RootPath+"/internal/domain/usecases/get_"+te.SnackCaseModelNames+"_use_case.go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFile(dtoScriptTemplate,
		te.RootPath+"/internal/view/dto/"+te.SnackCaseModelName+".go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFile(controllerScriptTemplate,
		te.RootPath+"/internal/view/controllers/"+te.SnackCaseModelName+"_controller.go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFile(routerScriptTemplate,
		te.RootPath+"/internal/view/routers/"+te.SnackCaseModelName+"_router.go")
	if err != nil {
		return err
	}

	err = te.writeTemplateToFileIfNotExist(controllerScripteTemplate,
		te.RootPath+"/internal/view/controllers/controller.go")
	if err != nil {
		return err
	}

	err = te.appendTemplateToFile(repositoryDefinitionScriptTemplate,
		te.RootPath+"/internal/view/controllers/controller.go")
	if err != nil {
		return err
	}

	err = te.appendTemplateToFile(usecaseDefinitionScriptTemplate,
		te.RootPath+"/internal/view/controllers/controller.go")
	if err != nil {
		return err
	}

	return nil
}

func (te *templateWriter) writeTemplateToFile(templateName string, filename string) error {
	var err error
	t, err := template.New("template").Parse(templateName)
	if err != nil {
		LogError("error parsing template: %v", err)
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, te.Data)
	if err != nil {
		LogError("error executing template: %v", err)
		return err
	}
	return err
}

func (te *templateWriter) writeTemplateToFileIfNotExist(templateName string, filename string) error {
	var err error
	t, err := template.New("template").Parse(templateName)
	if err != nil {
		LogError("error parsing template: %v", err)
		return err
	}

	_, err = os.Stat(filename)
	if err == nil {
		return nil
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, te.Data)
	if err != nil {
		LogError("error executing template: %v", err)
		return err
	}
	return err
}

func (te *templateWriter) appendTemplateToFile(templateName string, filename string) error {
	var err error
	t, err := template.New("template").Parse(templateName)
	if err != nil {
		LogError("error parsing template: %v", err)
		return err
	}

	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	contents := string(contentBytes)

	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufferTemplate := new(bytes.Buffer)
	err = t.Execute(bufferTemplate, te.Data)
	if err != nil {
		panic(err)
	}

	if strings.Contains(contents, bufferTemplate.String()) {
		return nil
	}

	_, err = file.Write(bufferTemplate.Bytes())
	if err != nil {
		panic(err)
	}
	return nil
}

const entityScriptTemplate = `
package entities

type {{.StructName}} struct {
	ID       int64
}
	
type {{.StructNames}} []{{.StructName}}`

const repositoryInterfaceScriptTemplate = `
package repositories

import (
	"context"

	"github.com/isjhar/iet/internal/domain/entities"
)

type {{.StructName}}Repository interface {
	Count(ctx context.Context, arg Count{{.StructNames}}Params) (int64, error)
	Get(ctx context.Context, arg Get{{.StructNames}}Params) (entities.{{.StructNames}}, error)
}

type Get{{.StructNames}}Params struct {
	GetParams
	Filter{{.StructName}}Params
}

type Count{{.StructNames}}Params struct {
	FilterParams
	Filter{{.StructName}}Params
}

type Filter{{.StructName}}Params struct {
}
`

const modelTemplate = `
package models

type {{.StructName}} struct {
	ID int64
}

func ({{.StructName}}) TableName() string {
	return "{{.SnackCaseModelName}}_detail"
}

type {{.StructNames}} []{{.StructName}}
`

const repositoryImplScriptTemplate = `
package repositories

import (
	"context"

	"github.com/isjhar/iet/internal/data/models"

	"github.com/isjhar/iet/pkg"

	"github.com/isjhar/iet/internal/domain/entities"
	"github.com/isjhar/iet/internal/domain/repositories"

	"gorm.io/gorm"
)

type {{.StructName}}Repository struct {
}

func (r {{.StructName}}Repository) Count(ctx context.Context, arg repositories.Count{{.StructNames}}Params) (int64, error) {
	var results int64
	var query *gorm.DB

	if arg.ID.Valid {
		query = r.createFindFilterQuery(ctx, arg)
	} else {
		query = r.createGetFilterQuery(ctx, arg)
	}

	err := query.Count(&results).Error
	if err != nil {
		pkg.LogError("error count {{.SnackCaseModelName}} %v", err)
		return results, entities.InternalServerError
	}

	return results, nil
}

func (r {{.StructName}}Repository) Get(ctx context.Context, arg repositories.Get{{.StructNames}}Params) (entities.{{.StructNames}}, error) {
	results := entities.{{.StructNames}}{}
	var err error

	if arg.ID.Valid {
		item, err := r.find(ctx, arg)
		if err != nil {
			return results, err
		}
		results = append(results, item)
	} else {
		results, err = r.get(ctx, arg)
		if err != nil {
			return results, err
		}
	}
	return results, nil
}

func (r {{.StructName}}Repository) get(ctx context.Context, arg repositories.Get{{.StructNames}}Params) (entities.{{.StructNames}}, error) {
	var results entities.{{.StructNames}}
	var models models.{{.StructNames}}

	query := r.createGetFilterQuery(ctx, repositories.Count{{.StructNames}}Params{
		Filter{{.StructName}}Params: arg.Filter{{.StructName}}Params,
		FilterParams:                       arg.FilterParams,
	})

	limit := int(-1)
	if arg.Limit.Valid {
		limit = int(arg.Limit.Int64)
	}
	offset := int(0)
	if arg.Offset.Valid {
		offset = int(arg.Offset.Int64)
	}
	orderBy := ""
	switch arg.Sort.String {
	default:
		orderBy += "{{.SnackCaseModelName}}.tanggal"
	}
	orderBy += " " + GetOrderQuery(arg.Order)

	query = query.Limit(limit).Offset(offset).Order(orderBy + ` + "`, {{.SnackCaseModelName}}.\"OID\"`)" + `
	err := query.Find(&models).Error
	if err != nil {
		pkg.LogError("error get {{.SnackCaseModelName}} %v", err)
		return results, entities.InternalServerError
	}

	for _, model := range models {
		result := entities.{{.StructName}}{
			ID: model.ID,
		}
		results = append(results, result)
	}

	return results, nil
}

func (r {{.StructName}}Repository) find(ctx context.Context, arg repositories.Get{{.StructNames}}Params) (entities.{{.StructName}}, error) {
	var model models.{{.StructName}}
	var result entities.{{.StructName}}

	query := r.createFindFilterQuery(ctx, repositories.Count{{.StructNames}}Params{
		FilterParams:                       arg.FilterParams,
		Filter{{.StructName}}Params: arg.Filter{{.StructName}}Params,
	})

	err := query.First(&model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		pkg.LogError("error find {{.SnackCaseModelName}} %v", err)
		return result, entities.InternalServerError
	}

	result.ID = model.ID
	return result, nil
}

func (r {{.StructName}}Repository) createGetFilterQuery(ctx context.Context, arg repositories.Count{{.StructNames}}Params) *gorm.DB {
	query := ORM.WithContext(ctx).
		Model(&models.{{.StructName}}{})

	query = r.applyDefaultFilterQuery(query)

	query = r.applySearchFilterQuery(query, arg)

	return query
}

func (r {{.StructName}}Repository) createFindFilterQuery(ctx context.Context, arg repositories.Count{{.StructNames}}Params) *gorm.DB {
	query := ORM.WithContext(ctx).
		Model(&models.{{.StructName}}{})

	query = r.applyDefaultFilterQuery(query)

	query = query.Where(` + "`{{.SnackCaseModelName}}.\"OID\" = ?`" + `, arg.ID)

	return query
}

func (r {{.StructName}}Repository) applySearchFilterQuery(query *gorm.DB, arg repositories.Count{{.StructNames}}Params) *gorm.DB {
	if arg.Search.Valid && arg.Search.String != "" {
		query = query.Where(
			ORM.Where("{{.SnackCaseModelName}}.nomor ilike ?", "%"+arg.Search.String+"%"),
		)
	}

	return query
}

func (r {{.StructName}}Repository) applyDefaultFilterQuery(query *gorm.DB) *gorm.DB {
	query = query.Where(` + "`{{.SnackCaseModelName}}.\"GCRecord\" is null`" + `)

	return query
}
`

const usecaseScriptTemplate = `
package usecases

import (
	"context"

	"github.com/isjhar/iet/internal/domain/entities"
	"github.com/isjhar/iet/internal/domain/repositories"

	"gopkg.in/guregu/null.v4"
)

type Get{{.StructNames}}UseCase struct {
	{{.StructName}}Repository repositories.{{.StructName}}Repository
}

type Get{{.StructNames}}UseCaseParams struct {
	GetUseCaseParams
	Filter{{.StructName}}UseCaseParams
}

type Filter{{.StructName}}UseCaseParams struct {
}

type Get{{.StructNames}}UseCaseResult struct {
	Items entities.{{.StructNames}}
	Total int64
}

func (i *Get{{.StructNames}}UseCase) Execute(ctx context.Context, arg Get{{.StructNames}}UseCaseParams) (Get{{.StructNames}}UseCaseResult, error) {
	var result Get{{.StructNames}}UseCaseResult

	filterParams := repositories.Filter{{.StructName}}Params{}

	count, err := i.{{.StructName}}Repository.Count(ctx, repositories.Count{{.StructNames}}Params{
		FilterParams: repositories.FilterParams{
			Search: arg.Search,
		},
		Filter{{.StructName}}Params: filterParams,
	})
	if err != nil {
		return result, err
	}

	limit := count
	if arg.Limit.Valid {
		limit = arg.Limit.Int64
	}

	items, err := i.{{.StructName}}Repository.Get(ctx, repositories.Get{{.StructNames}}Params{
		GetParams: repositories.GetParams{
			Limit:  null.IntFrom(limit),
			Offset: arg.Offset,
			Sort:   arg.Sort,
			Order:  arg.Order,
			FilterParams: repositories.FilterParams{
				Search: arg.Search,
				ID:     arg.ID,
			},
		},
		Filter{{.StructName}}Params: filterParams,
	})
	if err != nil {
		return result, err
	}

	result.Items = items
	result.Total = count
	return result, nil
}
`
const dtoScriptTemplate = `
package dto

import (
	"github.com/isjhar/iet/internal/domain/entities"
)

type Count{{.StructNames}}Params struct {
	FilterParams
	Filter{{.StructName}}Params
}

type Get{{.StructNames}}Params struct {
	GetParams
	Filter{{.StructName}}Params
}

type Filter{{.StructName}}Params struct {
}

type Get{{.StructNames}}Response struct {
	Response
	Data Get{{.StructNames}}ResponseData ` + "`json:\"data\"`" + `
}

type Get{{.StructNames}}ResponseData struct {
	Items {{.StructNames}} ` + "`json:\"items\"`" + `
	Total int64            ` + "`json:\"total\"`" + `
}

type {{.StructName}} struct {
	ID int64 ` + "`json:\"id\"`" + `
}

func New{{.StructName}}(entity entities.{{.StructName}}) {{.StructName}} {
	item := {{.StructName}}{
		ID: entity.ID,
	}
	return item
}

type {{.StructNames}} []{{.StructName}}

func New{{.StructNames}}(entities entities.{{.StructNames}}) {{.StructNames}} {
	result := make({{.StructNames}}, 0)
	for _, entity := range entities {
		item := New{{.StructName}}(entity)
		result = append(result, item)
	}
	return result
}
`

const controllerScriptTemplate = `
package controllers

import (
	"net/http"

	"github.com/isjhar/iet/internal/view/dto"

	"github.com/isjhar/iet/internal/domain/usecases"

	"github.com/labstack/echo/v4"
)

func Get{{.StructNames}}() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := &dto.Get{{.StructNames}}Params{}
		if err := c.Bind(data); err != nil {
			return err
		}

		ctx := c.Request().Context()

		result, err := get{{.StructNames}}UseCase.Execute(ctx, usecases.Get{{.StructNames}}UseCaseParams{
			GetUseCaseParams: usecases.GetUseCaseParams{
				Limit:  data.Limit,
				Offset: data.Offset,
				Sort:   data.Sort,
				Order:  data.Order,
				FilterUseCaseParams: usecases.FilterUseCaseParams{
					Search: data.Search,
					ID:     data.ID,
				},
			},
			Filter{{.StructName}}UseCaseParams: usecases.Filter{{.StructName}}UseCaseParams{},
		})

		if err != nil {
			return err
		}

		items := dto.New{{.StructNames}}(result.Items)
		return c.JSON(http.StatusOK, dto.Get{{.StructNames}}Response{
			Data: dto.Get{{.StructNames}}ResponseData{
				Items: items,
				Total: result.Total,
			},
		})
	}
}
`

const routerScriptTemplate = `
package routers

import (
	"github.com/isjhar/iet/internal/view/controllers"

	"github.com/labstack/echo/v4"
)

func {{.StructName}}Route(api *echo.Group) {
	api.OPTIONS("{{.KebabCaseModelNames}}", controllers.Get{{.StructNames}}())
	api.GET("{{.KebabCaseModelNames}}", controllers.Get{{.StructNames}}())
}
`

const controllerScripteTemplate = `
package controllers

import (
	"github.com/isjhar/iet/internal/data/repositories"
	"github.com/isjhar/iet/internal/domain/usecases"
)
`

const repositoryDefinitionScriptTemplate = `
var {{.CamelCaseModelName}}Repository = repositories.{{.StructName}}Repository{}`

const usecaseDefinitionScriptTemplate = `
var get{{.StructNames}}UseCase = usecases.Get{{.StructNames}}UseCase{
	{{.StructName}}Repository: {{.CamelCaseModelName}}Repository,
}`
