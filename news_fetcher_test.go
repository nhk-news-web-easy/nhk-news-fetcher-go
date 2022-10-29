package nhk_fetcher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFetchNews(t *testing.T) {
	startDate := time.Date(2022, 10, 28, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 10, 28, 23, 59, 59, 0, time.UTC)
	newsList, err := FetchNews(startDate, endDate)

	assert.True(t, len(newsList) > 0)
	assert.Nil(t, err)
}
