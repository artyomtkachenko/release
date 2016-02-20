package main
import (
  "errors"
  "regexp"
)

type Pattern struct {
  Field string
  Pattern string
  Content string
}

func validateInput(rm ReleaseMeta) error {
//TODO add all possible validations
  patterns := []Pattern{
    {
      "Name",
      "^[a-z0-9_-]+$",
      rm.Project.Name,
    },
  }

  for _, item := range patterns {
    if match, _ := regexp.MatchString(item.Pattern, item.Content); match == false {
      return errors.New("Validation failed: " + item.Field + " should follow the " + item.Pattern + " pattern.")
    }
  }
  return nil
}
