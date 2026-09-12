package compile

import (
	"runtime"
	"testing"
)

func TestAssertOutputPath(t *testing.T) {
	for _, ok := range []string{
		"out/main.luau",
		"out/sub/init.luau",
		"out/weird..name.luau",
	} {
		if err := assertOutputPath(".", "out", ok); err != nil {
			t.Errorf("assertOutputPath(%q) = %v, want nil", ok, err)
		}
	}
	for _, bad := range []string{
		"../escape.luau",
		"out/../../escape.luau",
		"..",
	} {
		if err := assertOutputPath(".", "out", bad); err == nil {
			t.Errorf("assertOutputPath(%q) = nil, want error", bad)
		}
	}
	if err := assertOutputPath(".", "../build", "../build/main.luau"); err != nil {
		t.Errorf("configured external output was rejected: %v", err)
	}
	if runtime.GOOS == "windows" {
		if err := assertOutputPath(`C:\project`, `C:\build`, `..\build\main.luau`); err != nil {
			t.Errorf("external absolute outDir rejected: %v", err)
		}
	}
}
