package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "adm64" {
		t.Skip("Não funciona nessa aquitetura")
	}
	// ...and
	t.Fail()
}
