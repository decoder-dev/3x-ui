package sub

import "testing"

func TestHappHeaderConfig(t *testing.T) {
	cfg := happHeaderConfig{ChangeUserAgent: true, Fingerprint: "ios"}
	h := cfg.httpHeaders()
	if h["change-user-agent"] == "" || h["user-agent-geo-files"] != "safari-ios" {
		t.Fatalf("unexpected ios headers: %#v", h)
	}
	comments := cfg.bodyCommentLines(h)
	if len(comments) != 2 {
		t.Fatalf("expected 2 comment lines, got %d", len(comments))
	}

	cfg = happHeaderConfig{ChangeUserAgent: false, Fingerprint: "chrome"}
	if len(cfg.httpHeaders()) != 0 {
		t.Fatalf("disabled change-user-agent should emit no headers")
	}

	cfg = happHeaderConfig{
		ChangeUserAgent: true,
		UserAgent:       "CustomUA/1.0",
		UserAgentGeo:    "custom-geo",
	}
	h = cfg.httpHeaders()
	if h["change-user-agent"] != "CustomUA/1.0" || h["user-agent-geo-files"] != "custom-geo" {
		t.Fatalf("overrides ignored: %#v", h)
	}
}
