package attest

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/docker/labs-brown-tape/attest/digest"
	"github.com/docker/labs-brown-tape/attest/types"
)

// MockPathChecker es una implementación simulada de PathChecker para pruebas.
type MockPathChecker struct {
	detectRepoResult bool
	detectRepoError  error
}

func (m *MockPathChecker) DetectRepo() (bool, error) {
	return m.detectRepoResult, m.detectRepoError
}

// TestDetectVCS prueba la función DetectVCS.
func TestDetectVCS(t *testing.T) {
	tests := []struct {
		name             string
		path             string
		mockRepoResult   bool
		mockRepoError    error
		expectedSuccess  bool
		expectedError    error
	}{
		{"valid VCS detection", "/valid/path", true, nil, true, nil},
		{"invalid VCS detection", "/invalid/path", false, nil, false, nil},
		{"empty path", "", false, fmt.Errorf("path cannot be empty"), false, fmt.Errorf("path cannot be empty")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProvider := func(string, digest.SHA256) types.PathChecker {
				return &MockPathChecker{
					detectRepoResult: tt.mockRepoResult,
					detectRepoError:  tt.mockRepoError,
				}
			}

			// Cambiar el proveedor a un mock.
			originalProvider := git.ProviderName // Asegúrate de tener esto correcto.
			git.ProviderName = "mock"             // Cambia esto según tu implementación real.
			defer func() { git.ProviderName = originalProvider }()

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
