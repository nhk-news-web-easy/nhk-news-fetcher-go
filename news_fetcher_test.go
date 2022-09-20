package nhk

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateNews(t *testing.T) {
	startDate := time.Date(2022, 8, 8, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 8, 8, 23, 59, 59, 0, time.UTC)
	newsList, err := FetchNews(startDate, endDate)

	assert.True(t, len(newsList) > 0)
	assert.Nil(t, err)
}
