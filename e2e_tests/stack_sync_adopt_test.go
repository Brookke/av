package e2e_tests

import (
	"testing"

	"github.com/aviator-co/av/internal/git/gittest"
	"github.com/aviator-co/av/internal/meta"
	"github.com/stretchr/testify/assert"
)

func TestStackSyncAdopt(t *testing.T) {
	server := RunMockGitHubServer(t)
	defer server.Close()
	repo := gittest.NewTempRepoWithGitHubServer(t, server.URL)
	Chdir(t, repo.RepoDir)

	repo.Git(t, "checkout", "-b", "stack-1")
	repo.CommitFile(t, "my-file", "1a\n", gittest.WithMessage("Commit 1a"))

	RequireAv(t, "stack", "sync", "--parent", "main")

	assert.Equal(t,
		meta.BranchState{
			Name:  "main",
			Trunk: true,
		},
		GetStoredParentBranchState(t, repo, "stack-1"),
		"stack-1 should be re-rooted onto main",
	)
}