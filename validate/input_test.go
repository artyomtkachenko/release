package validate

import (
	"testing"

	"github.com/artyomtkachenko/release/config"
)

func TestInput(t *testing.T) {
	shouldFail := []config.ReleaseConfig{
		{
			Project: config.Project{
				Name:        "@foo",
				ContentRoot: "/foo",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "fooar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "foo",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "npm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: true,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob@",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob@",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: config.Project{
				Name:        "foo",
				ContentRoot: "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: config.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: config.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: config.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "",
			},
		},
	}

	for _, test := range shouldFail {
		err := Input(test)
		if err == nil {
			t.Fatalf("err: %v", err)
		}
	}
	success := config.ReleaseConfig{
		Project: config.Project{
			Name:        "foo",
			ContentRoot: "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Publish: config.Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: config.Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: config.Deploy{
			User:    "bob",
			Group:   "bob",
			RootDir: "/root",
		},
	}
	err := Input(success)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

}
func BenchmarkInput(b *testing.B) {
	success := config.ReleaseConfig{
		Project: config.Project{
			Name:        "foo",
			ContentRoot: "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Publish: config.Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: config.Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: config.Deploy{
			User:    "bob",
			Group:   "bob",
			RootDir: "/root",
		},
	}
	for n := 0; n < b.N; n++ {
		Input(success)
	}
}
