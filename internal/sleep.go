package internal

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"time"
)

type Sleeper struct {
	DurationProvider DurationProvider
}

func (s *Sleeper) Sleep() {
	time.Sleep(s.DurationProvider.Duration())
}

type DurationProvider interface {
	Duration() time.Duration
}

type DurationProviderFn func() time.Duration

func (d DurationProviderFn) Duration() time.Duration {
	return d()
}

func NewFileProvider(filename string) (DurationProvider, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	b = bytes.TrimSpace(b)
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return nil, err
	}
	return DurationProviderFn(func() time.Duration {
		return time.Duration(i) * time.Millisecond
	}), nil
}

func NewProvider(duration time.Duration) DurationProvider {
	return DurationProviderFn(func() time.Duration {
		return duration
	})
}
