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

/**
 * VersionInfo - A struct to contain the version information of the software.
 *
 * Fields:
 *   GitVersion: (string) The version of Git used to build Terox.
 *   GitCommit: (string) The commit hash used to build Terox.
 *   BuildDate: (string) The datetime string when the Terox binary was built.
 *   GoVersion: (string) The Go version used to compile the Terox binary.
 *   Compiler: (string) The compiler used to build and compile the Terox binary.
 *   Platform: (string) The platform (OS & architecture) of the Terox binary.
 */
type VersionInfo struct {
	Version    string
	GitVersion string
	GitCommit  string
	BuildDate  string
	GoVersion  string
	Compiler   string
	Platform   string
}

/**
 * Get: Fetch and return the various version information.
 *
 * Parameters:
 *   None
 *
 * Returns:
 *   Returns an instance of the VersionInfo struct.
 */
func Get() *VersionInfo {
	return &VersionInfo{
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Version:    version,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
