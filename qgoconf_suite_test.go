package qgoconf_test

import (
	"github.com/fluent-qa/qgoconf"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestQgoconf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Qgoconf Suite")
}

var _ = Describe("Configuration test", func() {
	var appConfig *qgoconf.AppConfig
	BeforeEach(func() {
		appConfig, _ = qgoconf.NewYamlConfig(qgoconf.DefaultConfigFile)
	})
	It("should read config file", func() {
		Expect(appConfig).NotTo(BeNil())
	})
	It("should read config file", func() {
		result := appConfig.Viper.Get("name")
		Expect(result).To(Equal("FLUENT"))
	})
	It("should convert to Struct ", func() {
		appConfig.ToStruct(qgoconf.AnotherConfigInstance)
		named := &qgoconf.NamedMan{}
		appConfig.ToStructByKey("nested", named)
		Expect(named.Smith).To(Equal("smith"))
		Expect(named.Kevin).To(Equal("string"))
	})
})
