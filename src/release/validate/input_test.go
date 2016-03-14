package validate

import (
	"testing"

	"release/meta"
)

func TestInput(t *testing.T) {
	shouldFail := []meta.ReleaseMeta{
		{
			Project: meta.Project{
				Name:        "@foo",
				BuildRoot:   "/foo",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "fooar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "foo",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "https://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "npm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: true,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob@",
				Group:   "bob",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
				User:    "bob",
				Group:   "bob@",
				RootDir: "/myapp/foo",
			},
		},
		{
			Project: meta.Project{
				Name:        "foo",
				BuildRoot:   "/root",
				Email:       "foo@bar.com",
				Description: "Some text written here",
				ScmUrl:      "htt://foo.com",
				Version:     "1.0.2",
			},
			Publish: meta.Publish{
				Type:        "webdav",
				Destination: "http://foo.com",
			},
			Package: meta.Package{
				Type: "rpm",
				Sign: false,
			},
			Deploy: meta.Deploy{
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
	success := meta.ReleaseMeta{
		Project: meta.Project{
			Name:        "foo",
			BuildRoot:   "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Publish: meta.Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: meta.Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: meta.Deploy{
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
	success := meta.ReleaseMeta{
		Project: meta.Project{
			Name:        "foo",
			BuildRoot:   "/root",
			Email:       "foo@bar.com",
			Description: "Some text written here",
			ScmUrl:      "https://foo.com",
			Version:     "1.0.2",
		},
		Publish: meta.Publish{
			Type:        "webdav",
			Destination: "http://foo.com",
		},
		Package: meta.Package{
			Type: "rpm",
			Sign: false,
		},
		Deploy: meta.Deploy{
			User:    "bob",
			Group:   "bob",
			RootDir: "/root",
		},
	}
	for n := 0; n < b.N; n++ {
		Input(success)
	}
}
