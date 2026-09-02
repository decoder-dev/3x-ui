package sub

import "strings"

// Happ fingerprint presets for change-user-agent / user-agent-geo-files headers.
// Values mirror /agent/core_config.py (bot + sidecar subscription stack).
var happUAByFingerprint = map[string]string{
	"ios": `Mozilla/5.0 (iPhone; CPU iPhone OS 18_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Mobile/15E148 Safari/604.1`,
	"safari": `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Safari/605.1.15`,
	"chrome": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36`,
	"firefox": `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:128.0) Gecko/20100101 Firefox/128.0`,
}

var happGeoUAByFingerprint = map[string]string{
	"ios":     "safari-ios",
	"safari":  "safari-mac",
	"chrome":  "chrome-win",
	"firefox": "firefox-win",
}

type happHeaderConfig struct {
	ChangeUserAgent bool
	NoLimitXhttp    bool
	Fingerprint     string
	UserAgent       string
	UserAgentGeo    string
}

func (c happHeaderConfig) httpHeaders() map[string]string {
	if !c.ChangeUserAgent {
		return nil
	}
	out := map[string]string{}
	ua := strings.TrimSpace(c.UserAgent)
	if ua == "" {
		fp := strings.ToLower(strings.TrimSpace(c.Fingerprint))
		if fp == "" {
			fp = "chrome"
		}
		ua = happUAByFingerprint[fp]
	}
	if ua != "" {
		out["change-user-agent"] = ua
	}
	geo := strings.TrimSpace(c.UserAgentGeo)
	if geo == "" {
		fp := strings.ToLower(strings.TrimSpace(c.Fingerprint))
		if fp == "" {
			fp = "chrome"
		}
		geo = happGeoUAByFingerprint[fp]
	}
	if geo != "" {
		out["user-agent-geo-files"] = geo
	}
	return out
}

func (c happHeaderConfig) bodyCommentLines(headers map[string]string) []string {
	if len(headers) == 0 {
		return nil
	}
	lines := make([]string, 0, len(headers))
	for k, v := range headers {
		lines = append(lines, "#"+k+": "+v)
	}
	return lines
}
