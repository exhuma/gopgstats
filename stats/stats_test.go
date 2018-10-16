package stats

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    if got != "foo" {
        t.Errorf("Hello(%q) == %q, want %q", "", got, "foo")
    }
}
