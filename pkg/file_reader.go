package pkg

import (
	"github.com/satyrius/gonx"
	"io"
	"log"
	"os"
	"strconv"
)

type (
	LogParser struct {
		format string
		values chan float64
	}
)

func NewLogParser(format string) *LogParser {
	return &LogParser{
		format: format,
		values: make(chan float64),
	}
}

func (l *LogParser) GetValues() chan float64 {
	return l.values
}

func (l *LogParser) AddValues(logFile string) {

	file, err := os.Open(logFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer close(l.values)
	reader := gonx.NewReader(file, l.format)
	cnt := 0
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		// Process the record... e.g.
		val, err := rec.Field("response_time")
		if err != nil {
			log.Fatal(err)
		}
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			cnt++
			l.values <- f
		}
	}
}
