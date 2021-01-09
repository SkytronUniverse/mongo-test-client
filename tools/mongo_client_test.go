package tools_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	config "dev/interview-craft/configs"

	. "dev/interview-craft/tools"
)

var _ = Describe("MongoClient", func() {
	var (
		cfg config.Config
	)

	Describe("NewMongoClient", func() {
		BeforeEach(func() {
			cfg = config.Config{
				DB: config.Database{
					Server: "http://foo",
					Port:   "1234",
				},
				Details: config.Details{
					Name:       "fooDB",
					Collection: "barCollection",
				},
			}
		})

		It("Returns an new MongoClient", func() {
			mgoClient := NewMongoClient(cfg)
			Expect(mgoClient).NotTo(BeNil())
		})
	})
})
