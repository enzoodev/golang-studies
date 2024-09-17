package architecture

import (
	"runtime"
	"testing"
)

func TestDependency(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não é possível executar este teste em arquitetura amd64")
	}

	t.Fail()
}
