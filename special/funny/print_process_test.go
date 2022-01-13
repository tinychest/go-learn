package funny

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// https://play.golang.org/
func TestPrintProcess(t *testing.T) {
	const col = 30

	// Clear the screen by printing \x0c.
	bar := fmt.Sprintf("\x0c[%%-%vs]", col)

	for i := 0; i < col; i++ {
		t.Logf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	t.Logf(bar+" Done!", strings.Repeat("=", col))
}
