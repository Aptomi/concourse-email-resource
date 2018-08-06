package check_test

import (
	"github.com/Aptomi/concourse-email-resource/check"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Check", func() {
	AfterEach(func() {
		CleanupBuildArtifacts()
	})

	It("should compile", func() {
		_, err := Build("github.com/Aptomi/concourse-email-resource/check/cmd")
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("should output an empty JSON list", func() {
		output, err := check.Execute()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(output).Should(MatchJSON("[]"))
	})
})
