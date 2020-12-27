package testrepo

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"
)

func getSleepTime(filename string) (time.Duration, error) {
	b, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return 10 * time.Millisecond, nil
	}
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}
	return time.Duration(i) * time.Millisecond, nil
}

type sleeper struct {
	sleepTime time.Duration
}

func (s *sleeper) sleep() {
	time.Sleep(s.sleepTime)
}

func Benchmark_sleeper_sleep(b *testing.B) {
	sleepTime, err := getSleepTime("sleeptime.txt")
	if err != nil {
		b.Fatal(err)
	}
	sl := &sleeper{
		sleepTime: sleepTime,
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sl.sleep()
	}
}

func Benchmark_sleeper_sleep_doubled(b *testing.B) {
	sleepTime, err := getSleepTime("sleeptime.txt")
	if err != nil {
		b.Fatal(err)
	}
	sl := &sleeper{
		sleepTime: sleepTime * 2,
	}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sl.sleep()
	}
}
