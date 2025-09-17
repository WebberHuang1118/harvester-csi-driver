package version

import "fmt"

var (
	Version   = "v0.2.4"
	GitCommit = "HEAD"
)

func FriendlyVersion() string {
	return fmt.Sprintf("%s (%s)", Version, GitCommit)
}
