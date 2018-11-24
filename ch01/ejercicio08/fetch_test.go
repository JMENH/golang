
import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	if exit != 0 {
		//os.Exit(exit)
		fmt.Printf("result:%v", exit)
	}
}

func TestFetch(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"fetch", "http://github.com/JMENH"}, "resp.status 200 OK\n"},
		{[]string{"fetch", "https://github.com/JMENH"}, "resp.status 200 OK\n"},
		{[]string{"fetch", "github.com/JMENH"}, "resp.status 200 OK\n"},
		//{[]string{"fetch", "https://"}, "\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer)
		stderr = stdout
		main()
		got := stdout.(*bytes.Buffer).String()
		if !strings.Contains(got, test.expected) {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
