package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultXrayGithubRepo = "decoder-dev/Xray-core"

func xrayGithubRepo() string {
	if v := strings.TrimSpace(os.Getenv("XRAY_GITHUB_REPO")); v != "" {
		return v
	}
	return defaultXrayGithubRepo
}

func xrayGithubAPIReleasesURL() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/releases", xrayGithubRepo())
}

func xrayReleaseDownloadURL(version, fileName string) string {
	return fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", xrayGithubRepo(), version, fileName)
}

// xrayReleaseTagAllowed decides whether a GitHub release tag is offered for install.
func xrayReleaseTagAllowed(tagName string) bool {
	if strings.Contains(tagName, "-decoder") {
		return true
	}
	if xrayGithubRepo() != "XTLS/Xray-core" {
		return false
	}
	tagVersion := strings.TrimPrefix(tagName, "v")
	tagParts := strings.Split(tagVersion, ".")
	if len(tagParts) != 3 {
		return false
	}
	major, err1 := strconv.Atoi(tagParts[0])
	minor, err2 := strconv.Atoi(tagParts[1])
	patch, err3 := strconv.Atoi(tagParts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}
	return major > 26 || (major == 26 && minor > 6) || (major == 26 && minor == 6 && patch >= 27)
}

func xrayDigestRequired() bool {
	return xrayGithubRepo() == "XTLS/Xray-core"
}
