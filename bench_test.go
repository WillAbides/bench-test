package testrepo

import (
	"testing"
	"time"

	"github.com/willabides/bench-test/internal"
)

type sleeper struct {
	sleepTime time.Duration
}

func (s *sleeper) sleep() {
	time.Sleep(s.sleepTime)
}

func Benchmark_sleeper_sleep(b *testing.B) {
	dp, err := internal.NewFileProvider("sleeptime.txt")
	if err != nil {
		b.Fatal(err)
	}
	sl := &sleeper{
		sleepTime: dp.Duration(),
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sl.sleep()
	}
}

func Benchmark_sleeper_sleep_doubled(b *testing.B) {
	dp, err := internal.NewFileProvider("sleeptime.txt")
	if err != nil {
		b.Fatal(err)
	}
	sl := &sleeper{
		sleepTime: dp.Duration() * 2,
	}

	for i := 0; i < b.N; i++ {
		sl.sleep()
	}
}
