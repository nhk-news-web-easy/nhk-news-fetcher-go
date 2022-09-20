package nhk

import "time"

type News struct {
	NewsId          string
	Title           string
	TitleWithRuby   string
	OutlineWithRuby string
	Body            string
	BodyWithoutRuby string
	Url             string
	M3u8Url         string
	ImageUrl        string
	PublishedAtUtc  time.Time
}
