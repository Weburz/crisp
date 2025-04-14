package version

import (
	"runtime"
	"strings"
	"testing"
)

func TestGetVersionInfo_NotNil(t *testing.T) {
	info := GetVersionInfo()
	if info == nil {
		t.Fatal("expected non-nil versionInfo, got nil")
	}
}

func TestGetVersionInfo_RuntimeFields(t *testing.T) {
	info := GetVersionInfo()

	if info.GoVersion != runtime.Version() {
		t.Errorf(
			"GoVersion mismatch: got %s, want %s",
			info.GoVersion,
			runtime.Version(),
		)
	}

	if info.Compiler != runtime.Compiler {
		t.Errorf("Compiler mismatch: got %s, want %s", info.Compiler, runtime.Compiler)
	}

	expectedPlatform := runtime.GOOS + "/" + runtime.GOARCH
	if info.Platform != expectedPlatform {
		t.Errorf("Platform mismatch: got %s, want %s", info.Platform, expectedPlatform)
	}
}

func TestGetVersionInfo_DefaultBuildVars(t *testing.T) {
	info := GetVersionInfo()

	if info.Version != "unknown" {
		t.Errorf("expected default Version to be 'unknown', got %s", info.Version)
	}
	if info.GitVersion != "unknown" {
		t.Errorf("expected default GitVersion to be 'unknown', got %s", info.GitVersion)
	}
	if info.GitCommit != "unknown" {
		t.Errorf("expected default GitCommit to be 'unknown', got %s", info.GitCommit)
	}
	if info.BuildDate != "unknown" {
		t.Errorf("expected default BuildDate to be 'unknown', got %s", info.BuildDate)
	}
}

func TestPlatformFormat(t *testing.T) {
	info := GetVersionInfo()

	if !strings.Contains(info.Platform, "/") {
		t.Errorf("expected Platform format 'os/arch', got %s", info.Platform)
	}
}
