package aptible

import "testing"

func TestVersion(t *testing.T) {
	v := version()
	if v == "aptible/go-deploy/unknown" {
		t.Errorf("version() returned %q", v)
	}
}
