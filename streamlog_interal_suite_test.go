package streamlog

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInternalStreamlog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Streamlog Suite")
}
