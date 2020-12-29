package models

import (
	"billing/config"
	"time"
)

type EnsaasTime struct {
	Epoch  int64  `json:"-"`
	Format string `json:"-"`
}

func (p *EnsaasTime) UnmarshalJSON(data []byte) (err error) {
	timeFormat := "2006-01-02 15:04:05 -0700 UTC"
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	p.Epoch = now.Unix() * 1000
	return
}

func (p *EnsaasTime) MarshalJSON() ([]byte, error) {
	timeFormat := "2006-01-02 15:04:05 -0700 UTC"

	if p.Format == "marketplace" {
		timeFormat = "2006-01-02 15:04:05"
	}

	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Unix(p.Epoch/1000, 0).In(config.Location).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}
