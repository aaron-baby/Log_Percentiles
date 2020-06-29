package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogParser(t *testing.T) {
	format := `$remote_addr [$time_local] "$request" $status $response_time`
	l := NewLogParser(format)
	go l.AddValues("../log/sample.log")
	got := []float64{}
	for x := range l.GetValues() {
		got = append(got, x)
	}
	want := []float64{1230, 4630}
	assert := assert.New(t)
	if !assert.ElementsMatch(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
