
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"outline", "https://github.com/JMENH/"}
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()

	if len(got) == 0 {
		t.Errorf("no output")
	}

}

func TestOutline(t *testing.T) {
	var tests = []struct {
		url string
	}{
		{"https://github.com/JMENH/"},
	}

	for _, test := range tests {
		stdout = new(bytes.Buffer) // captured output
		outline2(test.url)         //original implementation
		expected := stdout.(*bytes.Buffer).String()

		stdout = new(bytes.Buffer) // captured output
		err := outline(test.url)
		actual := stdout.(*bytes.Buffer).String()

		if err != nil || actual != expected {
			t.Errorf("Expected:\n%v\nActual:\n%v", expected, actual)
		}
	}
}
