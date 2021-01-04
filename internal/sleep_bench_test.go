package internal_test

import (
	"path/filepath"
	"testing"

	"github.com/willabides/bench-test/internal"
)

func BenchmarkSleep_Sleep(b *testing.B) {
	dp, err := internal.NewFileProvider(filepath.FromSlash("../sleeptime.txt"))
	if err != nil {
		b.Fatal(err)
	}
	sleeper := &internal.Sleeper{
		DurationProvider: dp,
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sleeper.Sleep()
	}
}
