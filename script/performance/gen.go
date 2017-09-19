package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/icrowley/fake"
)

type Entry struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	for i := 1; i <= 10000; i++ {
		func() {
			dummyEntries := &Entry{
				Name:  fake.FullName(),
				Email: fake.EmailAddress(),
			}
			f, err := os.Create("./test-" + strconv.Itoa(i) + ".json")
			if err != nil {
				panic(err)
			}

			bytes, err := json.Marshal(dummyEntries)
			if err != nil {
				panic(err)
			}

			if _, err := f.Write(bytes); err != nil {
				panic(err)
			}
			defer f.Close()
		}()
	}
}
