package integrationtests

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

const (
	binaryName = "configuration-validator"
)

func TestMain(m *testing.M) {
	// TODO I could not find a more robust way of reaching the root of the repo
	err := os.Chdir(os.Getenv("GOREPO"))
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	make := exec.Command("make", "compile-validator-tool")
	err = make.Run()
	if err != nil {
		fmt.Printf("could not make binary for %s: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestConfigurationValidatorTool(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "configuration-validator integration tests suite")
}

var testCases = []TableEntry{
	{
		Description: "a positive test",
		Parameters: []interface{}{
			"./integration-tests/testdata/keyer_configuration_request_type_match_string_fragment.yaml",
			false,
			"",
		},
	},
	{
		Description: "a negative test",
		Parameters: []interface{}{
			"./integration-tests/testdata/keyer_configuration_missing_match_predicate.yaml",
			true,
			"invalid KeyerConfiguration.Fragments[0]: embedded message failed validation | caused by: " +
				"invalid KeyerConfiguration_Fragment.Rules[0]: embedded message failed " +
				"validation | caused by: invalid KeyerConfiguration_Fragment_Rule.Match: value is " +
				"required",
		},
	},
	{
		Description: "a negative test containing an inexistent file",
		Parameters: []interface{}{
			"bogus-file.yaml",
			true,
			"open bogus-file.yaml: no such file or directory",
		},
	},
}

var _ = Describe("Integration tests for the validator tool", func() {
	DescribeTable("table driven integration tests for the validator tool",
		func(ymlFilename string, wantErr bool, errorMessage string) {
			dir, err := os.Getwd()
			Expect(err).To(BeNil())

			// #nosec G204
			cmd := exec.Command(path.Join(dir, "bin", binaryName), "-config", ymlFilename)
			output, err := cmd.CombinedOutput()
			if wantErr {
				Expect(err).NotTo(BeNil())
				// There is an extra newline in the output.
				Expect(string(output)).Should(HaveSuffix(errorMessage + "\n"))
			} else {
				Expect(err).To(BeNil())
			}
		}, testCases...)

	It("should fail if -config flag is missing", func() {
		dir, err := os.Getwd()
		Expect(err).To(BeNil())

		// #nosec G204
		cmd := exec.Command(path.Join(dir, "bin", binaryName))
		output, err := cmd.CombinedOutput()
		Expect(err).NotTo(BeNil())
		// There is an extra newline in the output.
		expectedErrorMessage := "-config string\n    \tpath to configuration file\n"
		Expect(string(output)).Should(HaveSuffix(expectedErrorMessage))
	})
})
