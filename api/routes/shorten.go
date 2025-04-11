package routes

import "time"

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expire      time.Duration `json:"expire"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expire         time.Duration `json:"expire"`
	XRateRemainig  int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}
