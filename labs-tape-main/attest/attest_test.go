package attest

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/google/go-containerregistry/pkg/crane"
	. "github.com/onsi/gomega"

	. "github.com/docker/labs-brown-tape/attest"
	"github.com/docker/labs-brown-tape/attest/manifest"
	"github.com/docker/labs-brown-tape/manifest/imageresolver"
	"github.com/docker/labs-brown-tape/manifest/imagescanner"
	"github.com/docker/labs-brown-tape/manifest/loader"
	"github.com/docker/labs-brown-tape/manifest/testdata"
	"github.com/docker/labs-brown-tape/oci"
)

// TestDetectVCS prueba la función DetectVCS.
func TestDetectVCS(t *testing.T) {
	tests := []struct {
		name             string
		path             string
		expectedSuccess  bool
		expectedError    error
	}{
		{"valid VCS detection", "/valid/path", true, nil},
		{"invalid VCS detection", "/invalid/path", false, nil},
		{"empty path", "", false, fmt.Errorf("path cannot be empty")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			success, _, err := DetectVCS(tt.path)

			if success != tt.expectedSuccess || (err != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("expected success %v and error %v; got success %v and error %v",
					tt.expectedSuccess, tt.expectedError, success, err)
			}
		})
	}
}

// TestIntegrationDetectVCS prueba la detección del VCS en un entorno real.
func TestIntegrationDetectVCS(t *testing.T) {
	tempDir := t.TempDir()

	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to initialize git repo: %v", err)
	}

	success, _, err := DetectVCS(tempDir)
	if !success || err != nil {
		t.Errorf("expected success but got failure with error: %v", err)
	}

	filePath := fmt.Sprintf("%s/testfile.txt", tempDir)
	exec.Command("touch", filePath).Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "test commit").Run()

	success, _, err = DetectVCS(tempDir)
	if !success || err != nil {
		t.Errorf("expected success after commit but got failure with error: %v", err)
	}
}
