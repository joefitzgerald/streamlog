package streamlog_test

import (
	"bytes"

	"github.com/joefitzgerald/streamlog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	var (
		sl streamlog.Logger
	)

	Describe("API", func() {
		Context("New Logger", func() {
			JustBeforeEach(func() {
				sl = streamlog.New()
			})

			It("should create a Logger", func() {
				Ω(sl).ShouldNot(BeNil())
			})
		})

		Context("New Logger", func() {
			var (
				ow, ew bytes.Buffer
			)

			JustBeforeEach(func() {
				ow.Reset()
				ew.Reset()
				sl = streamlog.NewWithWriters(&ow, &ew, "", 0)
			})

			It("should create a Logger", func() {
				Ω(sl).ShouldNot(BeNil())
			})

			It("should print to the output stream when requested", func() {
				Ω(sl).ShouldNot(BeNil())
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.Println(streamlog.Out, "Hello, World!")
				Ω(ow.String()).Should(BeEquivalentTo("Hello, World!\n"))
				Ω(ow.Len()).ShouldNot(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
			})
		})
	})
})
