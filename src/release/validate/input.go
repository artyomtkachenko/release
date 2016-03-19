package validate

import (
	"errors"
	"reflect"
	"regexp"
	"release/config"
)

type pattern struct {
	field   string
	pattern string
	content string
}

func Input(rm config.ReleaseConfig) error {
	patterns := []pattern{
		{
			"Name",
			"^[a-z0-9_-]+$",
			rm.Project.Name,
		},
		{
			"ContentRoot",
			"\\w",
			rm.Project.ContentRoot,
		},
		{
			"Email",
			"^[a-zA-Z0-9_-]+@[a-z\\.]+$",
			rm.Project.Email,
		},
		{
			"Description",
			"\\w",
			rm.Project.Description,
		},
		{
			"ScmUrl",
			"^https",
			rm.Project.ScmUrl,
		},
		{
			"Version",
			"^[a-z0-9\\.]+$",
			rm.Project.Version,
		},
		{
			"Type",
			"^(webdav|nexus)$",
			rm.Publish.Type,
		},
		{
			"Destination",
			"^http",
			rm.Publish.Destination,
		},
		{
			"Type",
			"^(rpm|msi)$",
			rm.Package.Type,
		},
		{
			"User",
			"^[a-z0-9]+$",
			rm.Deploy.User,
		},
		{
			"Group",
			"^[a-z0-9]+$",
			rm.Deploy.Group,
		},
		{
			"RootDir",
			"\\w+",
			rm.Deploy.RootDir,
		},
	}

	for _, item := range patterns {
		match, _ := regexp.MatchString(item.pattern, item.content)
		if match == false {
			return errors.New("Validation failed: " + item.field + " should follow the " + item.pattern + " pattern, got: " + item.content)
		}
	}

	if reflect.TypeOf(rm.Package.Sign).Kind() != reflect.Bool {
		return errors.New("Validation failed: Sign field is not type of bool")
	}
	return nil
}
