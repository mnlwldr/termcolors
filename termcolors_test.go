package colors

import (
	"testing"
)

func Test(t *testing.T) {

	want := "\x1b[1;36m"
	got := BOLD_CYAN

	if want != got {
		t.Errorf("Test() = %q, want %q", got, want)
	}
}
