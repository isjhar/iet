package models

type ElasticsearchQueryLog struct {
	Type      string
	Name      string
	Timestamp string
	Duration  int64
}
