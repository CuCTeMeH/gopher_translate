package api

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("History API methods", func() {
	It("Get Sorted History", func() {
		enSentence := map[string]string{
			"english-sentence": "The quick brown fox jumps, over the lazy dog!",
		}
		enSentenceString, err := json.Marshal(enSentence)
		Expect(err).To(BeNil())

		expectedResult := map[string]string{
			"brown": "ownbrogo",
			"dog":   "ogdogo",
			"fox":   "oxfogo",
			"jumps": "umpsjogo",
			"lazy":  "azylogo",
			"over":  "gover",
			"quick": "ickquogo",
			"the":   "ethogo",
		}

		res := ExpectRequest(handler(), "POST", "/sentence", string(enSentenceString))
		res.run()

		res = ExpectRequest(handler(), "GET", "/history", "")
		res.run()

		respBody := res.Recorder.Body.Bytes()

		respMap := make(map[string]string)
		err = json.Unmarshal(respBody, &respMap)

		Expect(err).To(BeNil())
		Expect(respMap).To(BeEquivalentTo(expectedResult))
	})
})
