package streamlog

import (
	"bytes"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	var (
		sl StreamLogger
		l  *outErrStreamLogger
		ok bool
	)

	Describe("Internals", func() {
		Context("New Logger", func() {
			JustBeforeEach(func() {
				sl = New()
				l, ok = sl.(*outErrStreamLogger)
			})

			It("should create an outErrStreamLogger", func() {
				Ω(ok).To(Equal(true))
				Ω(l.out).ShouldNot(BeNil())
				Ω(l.err).ShouldNot(BeNil())
				outputLogger, ook := l.out.(*log.Logger)
				Ω(ook).To(Equal(true))
				Ω(outputLogger).ShouldNot(BeNil())
				errorLogger, eok := l.err.(*log.Logger)
				Ω(eok).To(Equal(true))
				Ω(errorLogger).ShouldNot(BeNil())
			})

			It("should propagate flags to internal log.Loggers", func() {
				outputLogger, _ := l.out.(*log.Logger)
				errorLogger, _ := l.err.(*log.Logger)
				Ω(log.LstdFlags).ShouldNot(Equal(log.Ldate))
				Ω(sl.Flags()).Should(Equal(log.LstdFlags))
				Ω(outputLogger.Flags()).Should(Equal(log.LstdFlags))
				Ω(errorLogger.Flags()).Should(Equal(log.LstdFlags))

				sl.SetFlags(log.Ldate)

				Ω(sl.Flags()).Should(Equal(log.Ldate))
				Ω(outputLogger.Flags()).Should(Equal(log.Ldate))
				Ω(errorLogger.Flags()).Should(Equal(log.Ldate))
			})

			It("should propagate prefix to internal log.Loggers", func() {
				outputLogger, _ := l.out.(*log.Logger)
				errorLogger, _ := l.err.(*log.Logger)
				Ω(sl.Prefix()).Should(Equal(""))
				Ω(outputLogger.Prefix()).Should(Equal(""))
				Ω(errorLogger.Prefix()).Should(Equal(""))

				sl.SetPrefix("test")

				Ω(sl.Prefix()).Should(Equal("test"))
				Ω(outputLogger.Prefix()).Should(Equal("test"))
				Ω(errorLogger.Prefix()).Should(Equal("test"))
			})
		})

		Context("NewWithWriters", func() {
			var (
				ow, ew bytes.Buffer
			)

			JustBeforeEach(func() {
				ow.Reset()
				ew.Reset()
				sl = NewWithWriters(&ow, &ew, "", 0)
				l, ok = sl.(*outErrStreamLogger)
			})

			It("should create an outErrStreamLogger", func() {
				Ω(ok).To(Equal(true))
				Ω(l.out).ShouldNot(BeNil())
				Ω(l.err).ShouldNot(BeNil())
				outputLogger, ook := l.out.(*log.Logger)
				Ω(ook).To(Equal(true))
				Ω(outputLogger).ShouldNot(BeNil())
				errorLogger, eok := l.err.(*log.Logger)
				Ω(eok).To(Equal(true))
				Ω(errorLogger).ShouldNot(BeNil())
			})

			It("Println should log to the out Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.Println(Out, "Hello, World!")
				Ω(ow.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).ShouldNot(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
			})

			It("Println should log to the err Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.Println(Err, "Hello, World!")
				Ω(ew.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).ShouldNot(Equal(0))
			})

			It("Print should log to the out Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.Print(Out, "Hello, World!")
				Ω(ow.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).ShouldNot(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
			})

			It("Print should log to the err Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.Print(Err, "Hello, World!")
				Ω(ew.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).ShouldNot(Equal(0))
			})

			It("Printf should log to the out Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.Printf(Out, "%v", "Hello, World!")
				Ω(ow.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).ShouldNot(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
			})

			It("Printf should log to the err Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.Printf(Err, "%v", "Hello, World!")
				Ω(ew.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).ShouldNot(Equal(0))
			})

			It("PrintlnOut should log to the out Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.PrintlnOut("Hello, World!")
				Ω(ow.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).ShouldNot(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
			})

			It("PrintlnErr should log to the err Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.PrintlnErr("Hello, World!")
				Ω(ew.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).ShouldNot(Equal(0))
			})

			It("PrintOut should log to the out Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.PrintOut("Hello, World!")
				Ω(ow.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).ShouldNot(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
			})

			It("PrintErr should log to the err Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.PrintErr("Hello, World!")
				Ω(ew.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).ShouldNot(Equal(0))
			})

			It("PrintfOut should log to the out Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.PrintfOut("%v", "Hello, World!")
				Ω(ow.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).ShouldNot(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
			})

			It("PrintfErr should log to the err Logger when requested", func() {
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).Should(Equal(0))
				sl.PrintfErr("%v", "Hello, World!")
				Ω(ew.String()).Should(BeEquivalentTo("[Hello, World!]\n"))
				Ω(ow.Len()).Should(Equal(0))
				Ω(ew.Len()).ShouldNot(Equal(0))
			})
		})

		Context("NewWithLoggers", func() {
			var (
				ow, ew bytes.Buffer
				ol, el *log.Logger
			)

			JustBeforeEach(func() {
				ow.Reset()
				ew.Reset()
				ol = log.New(&ow, "", 0)
				el = log.New(&ew, "", 0)
				sl = NewWithLoggers(ol, el, "", 0)
				l, ok = sl.(*outErrStreamLogger)
			})

			It("should create an outErrStreamLogger", func() {
				Ω(ok).To(Equal(true))
				Ω(l.out).ShouldNot(BeNil())
				Ω(l.err).ShouldNot(BeNil())
				outputLogger, ook := l.out.(*log.Logger)
				Ω(ook).To(Equal(true))
				Ω(outputLogger).ShouldNot(BeNil())
				errorLogger, eok := l.err.(*log.Logger)
				Ω(eok).To(Equal(true))
				Ω(errorLogger).ShouldNot(BeNil())
			})
		})
	})
})
