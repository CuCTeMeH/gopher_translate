package translator

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Translator methods", func() {
	It("Translate Word", func() {
		ew := []string{"square", "xray", "chair", "jumps", "dog", "brown", "lazy", "fox", "apple"}
		gw := []string{"aresquogo", "gexray", "airchogo", "umpsjogo", "ogdogo", "ownbrogo", "azylogo", "oxfogo", "gapple"}

		for i := 0; i < len(ew); i++ {
			res, err := TranslateWord(ew[i])
			Expect(err).To(BeNil())
			Expect(res).To(BeEquivalentTo(gw[i]))
		}
	})

	It("Translate Empty Word", func() {
		res, err := TranslateWord("")
		Expect(err.Error()).To(BeEquivalentTo("empty word"))
		Expect(res).To(BeEquivalentTo(""))
	})

	It("Translate Shortened Word", func() {
		res, err := TranslateWord("shouldn’t")
		Expect(err.Error()).To(BeEquivalentTo("gophers can not understand shortened versions of words or apostrophes"))
		Expect(res).To(BeEquivalentTo(""))
	})

	It("Translate Sentence", func() {
		se := "The quick brown fox jumps, over the lazy dog!"
		sg := "ethogo uickqogo ownbrogo oxfogo umpsjogo, gover ethogo azylogo ogdogo!"

		res, err := TranslateSentence(se)
		Expect(err).To(BeNil())
		Expect(res).To(BeEquivalentTo(sg))
	})
})
