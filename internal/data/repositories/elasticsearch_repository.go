package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/isjhar/iet/internal/config"
	"github.com/isjhar/iet/internal/data/models"
	"github.com/isjhar/iet/internal/domain/entities"
)

type ElasticsearchRepository struct {
}

func (r *ElasticsearchRepository) Write(p []byte) (int, error) {
	requestBody := models.ElasticsearchLog{
		Type:      "log",
		Message:   string(p),
		Timestamp: time.Now().Format(time.RFC3339),
	}
	jsonValue, err := json.Marshal(requestBody)
	if err != nil {
		return 0, err
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		category := config.Elasticsearch.Category.String
		url := "/" + category + "/_doc"
		req, err := r.NewRequest(ctx, http.MethodPost, url, bytes.NewBuffer(jsonValue))
		if err != nil {
			return
		}
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusCreated {
			return
		}
	}()
	return len(p), nil
}

func (r *ElasticsearchRepository) LogQuery(name string, duration int64) {
	requestBody := models.ElasticsearchQueryLog{
		Type:      "queryDb",
		Name:      name,
		Duration:  duration,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	jsonValue, err := json.Marshal(requestBody)
	if err != nil {
		return
	}
	r.SendLog(jsonValue)
}

func (r *ElasticsearchRepository) LogFieldChange(name string, message string) {
	requestBody := models.ElasticsearchFieldChangeLog{
		Type:      "fieldChange",
		Name:      name,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	jsonValue, err := json.Marshal(requestBody)
	if err != nil {
		return
	}
	r.SendLog(jsonValue)
}

func (r *ElasticsearchRepository) LogApi(arg LogApiParams) {
	requestBody := models.ElasticsearchApiLog{
		Type:       "api",
		Path:       arg.Path,
		Method:     arg.Method,
		Duration:   arg.Duration,
		Timestamp:  time.Now().Format(time.RFC3339),
		StatusCode: arg.StatusCode,
		UserAgent:  arg.UserAgent,
	}
	jsonValue, err := json.Marshal(requestBody)
	if err != nil {
		return
	}
	r.SendLog(jsonValue)
}

type LogApiParams struct {
	Path       string
	Method     string
	StatusCode int
	Duration   int64
	UserAgent  string
}

func (r ElasticsearchRepository) SendLog(body []byte) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		category := config.Elasticsearch.Category.String
		url := "/" + category + "/_doc"
		req, err := r.NewRequest(ctx, http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			return
		}
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusCreated {
			return
		}
	}()
}

func (r ElasticsearchRepository) NewRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	url := config.Elasticsearch.Url.String + path
	key := config.Elasticsearch.Key.String

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, entities.InternalServerError
	}
	req.Header.Add("Authorization", "Basic "+key)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
