package validate

import (
	"errors"
	"github.com/artyomtkachenko/release/meta"
	"regexp"
)

type pattern struct {
	field   string
	pattern string
	content string
}

func Input(rm ReleaseMeta) error {
	//TODO add all possible validations
	patterns := []pattern{
		{
			"Name",
			"^[a-z0-9_-]+$",
			rm.Project.Name,
		},
	}

	for _, item := range patterns {
		if match, _ := regexp.MatchString(item.pattern, item.content); match == false {
			return errors.New("Validation failed: " + item.field + " should follow the " + item.pattern + " pattern.")
		}
	}
	return nil
}
