package api

import (
	"encoding/json"
	"github.com/CuCTeMeH/gopher_translate/translator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Translator API methods", func() {
	BeforeSuite(func() {
		translator.InitHistoryStorage()
	})
	BeforeEach(func() {
		translator.ClearStorage()
	})

	It("Post Words", func() {
		ew := []string{"square", "xray", "chair", "jumps", "dog", "brown", "lazy", "fox", "apple"}
		gw := []string{"aresquogo", "gexray", "airchogo", "umpsjogo", "ogdogo", "ownbrogo", "azylogo", "oxfogo", "gapple"}

		for i := 0; i < len(ew); i++ {
			enWord := map[string]string{
				"english-word": ew[i],
			}
			enWordString, err := json.Marshal(enWord)
			Expect(err).To(BeNil())

			expectedResult := map[string]string{
				"gopher-word": gw[i],
			}

			res := ExpectRequest(handler(), "POST", "/word", string(enWordString))
			res.run()

			respBody := res.Recorder.Body.Bytes()

			respMap := make(map[string]string)
			err = json.Unmarshal(respBody, &respMap)

			Expect(err).To(BeNil())
			Expect(respMap).To(BeEquivalentTo(expectedResult))
		}
	})

	It("Post Sentence", func() {
		enSentence := map[string]string{
			"english-sentence": "The quick brown fox jumps, over the lazy dog!",
		}
		enSentenceString, err := json.Marshal(enSentence)
		Expect(err).To(BeNil())

		expectedResult := map[string]string{
			"gopher-sentence": "ethogo ickquogo ownbrogo oxfogo umpsjogo, gover ethogo azylogo ogdogo!",
		}

		res := ExpectRequest(handler(), "POST", "/sentence", string(enSentenceString))
		res.run()

		respBody := res.Recorder.Body.Bytes()

		respMap := make(map[string]string)
		err = json.Unmarshal(respBody, &respMap)

		Expect(respMap).To(BeEquivalentTo(expectedResult))
	})
})
