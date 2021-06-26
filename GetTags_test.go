package githubgo

import "testing"

func TestGetTags(t *testing.T) {
	gitHubUser := "companytone"
	accessToken := "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"
	GetTags(gitHubUser, accessToken)
}
