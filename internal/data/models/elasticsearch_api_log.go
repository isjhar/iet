package models

type ElasticsearchApiLog struct {
	Type       string
	Name       string
	Timestamp  string
	Duration   int64
	Path       string
	Method     string
	UserAgent  string
	StatusCode int
}
