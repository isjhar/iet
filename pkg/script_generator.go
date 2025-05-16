package pkg

import (
	"os"
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
		"StructName":         modelName,
		"StructNames":        i.convertToPrulal(modelName),
		"SnackCaseModelName": snackCaseModelName,
	}

	templateWriter := templateWriter{
		Data:               data,
		RootPath:           rootPath,
		SnackCaseModelName: snackCaseModelName,
	}

	return templateWriter.write()
}

func (i *ScriptGenerator) convertToPrulal(modelName string) string {
	if modelName[len(modelName)-1] == 'y' {
		return modelName[:len(modelName)-1] + "ies"
	}

	return modelName + "s"
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

type templateWriter struct {
	Data               map[string]string
	RootPath           string
	SnackCaseModelName string
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
