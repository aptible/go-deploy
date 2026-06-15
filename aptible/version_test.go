package aptible

import "testing"

func TestVersion(t *testing.T) {
	v := version()
	if v == "unknown" {
		t.Errorf("version() returned %q", v)
	}
}

func TestUserAgent(t *testing.T) {
	ua := userAgent()
	expected := "aptible/go-deploy/" + version()
	if ua != expected {
		t.Errorf("userAgent() = %q, want %q", ua, expected)
	}
}
