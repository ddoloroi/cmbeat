package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/yannbrrd/cmbeat/beater"
)

func main() {
	err := beat.Run("cmbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
