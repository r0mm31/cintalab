package attest

import (
	"fmt"
	"log"

	"github.com/docker/labs-brown-tape/attest/digest"
	"github.com/docker/labs-brown-tape/attest/types"
	"github.com/docker/labs-brown-tape/attest/vcs/git"
)

// PathCheckerRegistry es una estructura que almacena información sobre el PathChecker.
type PathCheckerRegistry struct {
	Path     string
	Provider string
	Checker  types.PathChecker
}

// NewPathCheckerRegistry crea una nueva instancia de PathCheckerRegistry.
func NewPathCheckerRegistry(path string, provider string) *PathCheckerRegistry {
	return &PathCheckerRegistry{
		Path:     path,
		Provider: provider,
	}
}

// DetectVCS detecta el sistema de control de versiones en la ruta especificada.
func DetectVCS(path string) (bool, *PathCheckerRegistry, error) {
	if path == "" {
		return false, nil, fmt.Errorf("path cannot be empty")
	}

	for _, provider := range map[string]func(string, digest.SHA256) types.PathChecker{
		git.ProviderName: git.NewPathChecker,
	} {
		checker := provider(path, "")
		ok, err := checker.DetectRepo()
		if err != nil {
			log.Printf("Error detecting VCS: %v", err)
			return false, nil, fmt.Errorf("unable to detect VCS: %w", err)
		}
		if ok {
			registry := NewPathCheckerRegistry(path, provider)
			if err := registry.init(checker); err != nil {
				return false, nil, err
			}
			return true, registry, nil
		}
	}
	return false, nil, nil
}

// init inicializa el PathCheckerRegistry.
func (r *PathCheckerRegistry) init(checker types.PathChecker) error {
	if checker == nil {
		return fmt.Errorf("checker cannot be nil")
	}
	r.Checker = checker
	return nil
}
