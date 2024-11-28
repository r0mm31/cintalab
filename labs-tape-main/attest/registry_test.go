package attest_test

import (
    "bytes"
    "context"
    "crypto/sha256"
    "encoding/hex"
    "io"
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
    // "github.com/docker/labs-brown-tape/trex"
)

var (
    craneOptions []crane.Option
    //makeDestination func(string) string
)

const repeat = 3

func TestRegistry(t *testing.T) {
    // trex.RunShared()
    // craneOptions = trex.Shared.CraneOptions()
    // makeDestination = trex.Shared.NewUniqueRepoNamer("bpt-registry-test")

    cases := testdata.BaseYAMLCasesWithDigests(t)
    cases.Run(t, ("../"), makeRegistryTest)
}

func makeRegistryTest(tc testdata.TestCase) func(t *testing.T) {
    return func(t *testing.T) {
        g := NewWithT(t)
        t.Parallel()

        checksums := make([]string, repeat)
        for i := range checksums {

            loader := loader.NewRecursiveManifestDirectoryLoader(tc.Directory)
            g.Expect(loader.Load()).To(Succeed())

            pathChecker, attreg, err := DetectVCS(tc.Directory)
            g.Expect(err).NotTo(HaveOccurred())
            g.Expect(pathChecker).ToNot(BeNil())
            g.Expect(attreg).ToNot(BeNil())

            scanner := imagescanner.NewDefaultImageScanner()
            scanner.WithProvinanceAttestor(attreg)

            expectedNumPaths := len(tc.Manifests)
            g.Expect(loader.Paths()).To(HaveLen(expectedNumPaths))
