package solution

import "testing"

func TestGetMessage(t *testing.T) {
    want := "Hello 🗺️ !"
    if got := GetMessage(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}