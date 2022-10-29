package nhk_fetcher

import "time"

type News struct {
	NewsId          string
	Title           string
	TitleWithRuby   string
	Outline         string
	OutlineWithRuby string
	Body            string
	BodyWithoutHtml string
	Url             string
	M3u8Url         string
	ImageUrl        string
	PublishedAtUtc  time.Time
}
