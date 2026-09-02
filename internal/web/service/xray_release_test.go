package service

import "testing"

func TestXrayReleaseTagAllowed(t *testing.T) {
	t.Setenv("XRAY_GITHUB_REPO", "decoder-dev/Xray-core")
	if !xrayReleaseTagAllowed("v26.7.28-decoder") {
		t.Fatal("decoder tag should be allowed")
	}
	if xrayReleaseTagAllowed("v26.7.28") {
		t.Fatal("plain upstream tag should be rejected on decoder repo default")
	}

	t.Setenv("XRAY_GITHUB_REPO", "XTLS/Xray-core")
	if !xrayReleaseTagAllowed("v26.7.28") {
		t.Fatal("upstream semver tag should be allowed")
	}
	if xrayReleaseTagAllowed("v25.1.0") {
		t.Fatal("old upstream tag should be rejected")
	}
}

func TestXrayGithubRepoDefault(t *testing.T) {
	t.Setenv("XRAY_GITHUB_REPO", "")
	if got := xrayGithubRepo(); got != defaultXrayGithubRepo {
		t.Fatalf("repo = %q, want %q", got, defaultXrayGithubRepo)
	}
}
