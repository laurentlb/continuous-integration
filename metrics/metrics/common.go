package metrics

import (
	"strings"

	"github.com/fweikert/go-buildkite/buildkite"
)

func getPlatfrom(job *buildkite.Job) string {
	if job.Name == nil {
		return ""
	} else if strings.Contains(*job.Name, "ubuntu") {
		return "linux"
	} else if strings.Contains(*job.Name, "windows") {
		return "windows"
	} else if strings.Contains(*job.Name, "darwin") {
		return "macos"
	} else if strings.Contains(*job.Name, "gcloud") {
		return "rbe"
	} else {
		return ""
	}
}

func getDifferenceSeconds(start *buildkite.Timestamp, end *buildkite.Timestamp) float64 {
	if start == nil || end == nil {
		return -1
	}
	return end.Time.Sub(start.Time).Seconds()
}