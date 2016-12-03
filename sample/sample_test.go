package sample_test

import (
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {

	var resp *http.Response
	var requestErr, decodeErr error
	var respData map[string]interface{}

	BeforeSuite(func() {
		resp, requestErr = http.Get(`https://api.github.com/users/aubm`)
		if requestErr == nil {
			decodeErr = json.NewDecoder(resp.Body).Decode(&respData)
		}
	})

	It("should not generate a request error", func() {
		Expect(requestErr).To(BeNil())
	})

	It("should not generate a decode error", func() {
		Expect(requestErr).To(BeNil())
	})

	It("should have a created_at field", func() {
		Expect(respData["created_at"]).To(Equal("2013-11-28T18:18:17Z"))
	})

})
