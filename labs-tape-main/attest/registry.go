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

// DetectRepo simula la detección de un repositorio.
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

	if err := initializeGitRepo(tempDir); err != nil {
		t.Fatalf("failed to initialize git repo: %v", err)
	}

	if success, _, err := DetectVCS(tempDir); !success || err != nil {
		t.Errorf("expected success but got failure with error: %v", err)
	}

	if err := createAndCommitFile(tempDir, "testfile.txt", "test commit"); err != nil {
		t.Fatalf("failed to create and commit file: %v", err)
	}

	if success, _, err := DetectVCS(tempDir); !success || err != nil {
		t.Errorf("expected success after commit but got failure with error: %v", err)
	}
}

// initializeGitRepo inicializa un repositorio Git en el directorio especificado.
func initializeGitRepo(dir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	return cmd.Run()
}

// createAndCommitFile crea un archivo y lo añade al repositorio Git.
func createAndCommitFile(dir, filename, commitMessage string) error {
	filePath := fmt.Sprintf("%s/%s", dir, filename)

	if err := exec.Command("touch", filePath).Run(); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	if err := exec.Command("git", "add", ".").Run(); err != nil {
		return fmt.Errorf("failed to add file to git: %w", err)
	}
	if err := exec.Command("git", "commit", "-m", commitMessage).Run(); err != nil {
		return fmt.Errorf("failed to commit file: %w", err)
	}
	return nil
}

// Función para subir imágenes (ejemplo).
func UploadImage(imagePath string) error {
	cmd := exec.Command("docker", "push", imagePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}
	return nil
}
