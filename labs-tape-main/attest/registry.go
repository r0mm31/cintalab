package attest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"slices"

	"github.com/docker/labs-brown-tape/attest/digest"
	"github.com/docker/labs-brown-tape/attest/manifest"
	"github.com/docker/labs-brown-tape/attest/types"
	"github.com/fxamacker/cbor/v2"
)

type PathCheckerRegistry struct {
	newPathChecker func(string, digest.SHA256) types.PathChecker

	registry     map[types.PathCheckerRegistryKey]types.PathChecker
	mutatedPaths types.Mutations
	statements   types.Statements

	baseDir baseDir // Declarar el tipo explícitamente
}

type baseDir struct {
	pathChecker   types.PathChecker
	cachedSummary types.PathCheckSummary

	fromWorkDir, fromRepoRoot string
}

func NewPathCheckerRegistry(dir string, newPathChecker func(string, digest.SHA256) types.PathChecker) *PathCheckerRegistry {
	return &PathCheckerRegistry{
		baseDir:        baseDir{fromWorkDir: dir},
		newPathChecker: newPathChecker,
		registry:       map[types.PathCheckerRegistryKey]types.PathChecker{},
		statements:     types.Statements{},
	}
}

// Resto de las funciones...
