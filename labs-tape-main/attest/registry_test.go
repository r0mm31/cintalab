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
		{
			name:            "valid VCS detection",
			path:            "/valid/path",
			mockRepoResult:  true,
			expectedSuccess: true,
			expectedError:   nil,
		},
		{
			name:            "invalid VCS detection",
			path:            "/invalid/path",
			mockRepoResult:  false,
			expectedSuccess: false,
			expectedError:   nil,
		},
		{
			name:          "empty path",
			path:          "",
			expectedSuccess: false,
			expectedError: fmt.Errorf("path cannot be empty"),
        },
        // Más casos de prueba...
    }

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockProvider := func(string, digest.SHA256) types.PathChecker {
                return &MockPathChecker{
                    detectRepoResult: tt.mockRepoResult,
                    detectRepoError:  tt.mockRepoError,
                }
            }

            // Aquí debes asegurarte de que estás usando el mock correctamente.
            originalProvider := git.ProviderName // Cambia esto si es necesario.
            git.ProviderName = "mock" // Cambia esto según tu implementación real.
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
    // Crear un directorio temporal para la prueba.
    tempDir := t.TempDir()
    
    // Inicializar un repositorio Git en el directorio temporal.
    cmd := exec.Command("git", "init")
    cmd.Dir = tempDir
    if err := cmd.Run(); err != nil {
        t.Fatalf("failed to initialize git repo: %v", err)
    }

    // Probar la detección del VCS en el directorio temporal.
    success, _, err := DetectVCS(tempDir)
    if !success || err != nil {
        t.Errorf("expected success but got failure with error: %v", err)
    }

    // Agregar un archivo y hacer commit.
    filePath := fmt.Sprintf("%s/testfile.txt", tempDir)
    exec.Command("touch", filePath).Run()
    exec.Command("git", "add", ".").Run()
    exec.Command("git", "commit", "-m", "test commit").Run()

    // Volver a probar la detección del VCS.
    success, _, err = DetectVCS(tempDir)
    if !success || err != nil {
        t.Errorf("expected success after commit but got failure with error: %v", err)
    }
}
