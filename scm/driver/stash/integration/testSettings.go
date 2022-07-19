package integration

import (
	"context"
	//"os"

	"github.com/drone/go-scm/scm"
)

var (
	client *scm.Client
	token  = "OTI1NDU1Mzc4NjMyOgLLspuvcOvQBlETmo58Ap6nSi3z" //os.Getenv("BITBUCKET_SERVER_TOKEN")

	endpoint = "https://bitbucket.dev.harness.io/"
	repoID   = "har/deepakgitsynctest"
	username = "harnessadmin"
)

func GetCurrentCommitOfBranch(client *scm.Client, branch string) (string, error) {
	commits, _, err := client.Git.ListCommits(context.Background(), repoID, scm.CommitListOptions{Ref: branch})
	if err != nil {
		return "", err
	}
	return commits[0].Sha, nil
}
