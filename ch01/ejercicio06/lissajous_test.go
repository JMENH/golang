import (
	"bytes"
	"testing"
)

func TestLissajous(t *testing.T) {
	out = new(bytes.Buffer)
	main()
	got := out.(*bytes.Buffer).String()
	if size := len(got); size == 0 {
		t.Errorf("Result data size = %d", size)
	}
}
