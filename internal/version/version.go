// Package version provides build-time and runtime version information for Crisp.
//
// It includes details such as the semantic version, Git version, commit hash,
// build date, Go runtime version, compiler, and target platform. These values
// are intended to be set at build time using -ldflags and are typically
// displayed when the `crisp version` command is executed.
package version

import (
	"fmt"
	"runtime"
)

// These variables store version metadata and are meant to be overridden at build time
// using -ldflags. For example:
//
// go build -ldflags="-X 'github.com/Weburz/crisp/internal/version.version=1.2.3' ..."
var (
	version    = "unknown" // Semantic version of the application
	gitVersion = "unknown" // Git tag or descriptive version (e.g., v1.2.3-4-gabcdef)
	gitCommit  = "unknown" // Git commit hash
	buildDate  = "unknown" // Build date in YYYY-MM-DD format
)

// versionInfo holds metadata about the build and runtime environment.
type versionInfo struct {
	Version    string // Semantic version
	GitVersion string // Git tag or descriptive version
	GitCommit  string // Git commit hash
	BuildDate  string // Build timestamp
	GoVersion  string // Version of the Go runtime
	Compiler   string // Compiler used (e.g., "gc")
	Platform   string // OS/Architecture (e.g., linux/amd64)
}

// GetVersionInfo returns a populated versionInfo struct containing both
// build-time and runtime metadata about Crisp.
func GetVersionInfo() *versionInfo {
	return &versionInfo{
		Version:    version,
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
