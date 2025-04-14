/**
 * Package version - Handle various version information of Terox.
 *
 * The "version" package handles various version information like the Git, Go
 * and the compiler version. Along with those information, the package also
 * handles the Git commit hash to print out when the "terox version" command is
 * invoked.
 */
package version

import (
	"fmt"
	"runtime"
)

// Store the Git version, the commit hash and the build date to print the
// version info.
var (
	version    = "unknown"
	gitVersion = "unknown"
	gitCommit  = "unknown"
	buildDate  = "unknown"
)

type versionInfo struct {
	Version    string
	GitVersion string
	GitCommit  string
	BuildDate  string
	GoVersion  string
	Compiler   string
	Platform   string
}

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
