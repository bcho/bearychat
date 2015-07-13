package bearychat

import (
	"testing"
)

func TestBuildIncomingMessage(t *testing.T) {
	var (
		m   Incoming
		err error
	)

	m = Incoming{}
	_, err = m.Build()
	if err == nil {
		t.Errorf("expected error when missing text")
	}

	m = Incoming{
		Text: "foobar",

		Attachments: []IncomingAttachment{
			IncomingAttachment{},
			IncomingAttachment{},
		},
	}
	_, err = m.Build()
	if err == nil {
		t.Errorf("expected error when missing text and title")
	}

	m = Incoming{
		Text:     "foobar",
		Markdown: true,
	}
	_, err = m.Build()
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	m = Incoming{
		Text:     "foobar",
		Markdown: true,

		Attachments: []IncomingAttachment{
			IncomingAttachment{
				Title: "foo",
			},
			IncomingAttachment{
				Text: "bar",
				Images: []IncomingImage{
					IncomingImage{Url: "http://foo.bar"},
				},
			},
		},
	}
	_, err = m.Build()
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}
