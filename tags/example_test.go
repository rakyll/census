package tags_test

import (
	"fmt"
	"log"

	"github.com/rakyll/census/tags"
)

func ExampleNewTagSet() {
	key, err := tags.NewStringKey("key")
	if err != nil {
		log.Fatal(err)
	}
	tagSet := tags.NewTagSet(nil,
		key.Update("foo"),
		key.Insert("bar"),
	)
	fmt.Println(tagSet)
}
