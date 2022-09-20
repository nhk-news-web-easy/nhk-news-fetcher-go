package nhk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const requestUrl = "https://nhk.dekiru.app/news"
const isoFormat = "2006-01-02T15:04:05.000Z07:00"

func FetchNews(startDate time.Time, endDate time.Time) ([]News, error) {
	url := fmt.Sprintf(requestUrl+"?startDate=%s&endDate=%s", startDate.Format(isoFormat), endDate.Format(isoFormat))
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Expected 200, but got %d", response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var newsList []News

	err = json.Unmarshal(body, &newsList)

	if err != nil {
		return nil, err
	}

	return newsList, nil
}
