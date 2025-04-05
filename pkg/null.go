package pkg

import "gopkg.in/guregu/null.v4"

func NullStringValueOrDefault(value null.String, defaultValue string) string {
	if !value.Valid {
		return defaultValue
	}
	return value.String
}

func NullIntValueOrDefault(value null.Int, defaultValue int64) int64 {
	if !value.Valid {
		return defaultValue
	}
	return value.Int64
}
