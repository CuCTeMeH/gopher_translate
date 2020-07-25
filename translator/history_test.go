package translator

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
	"sync"
)

func initHistory() *History {
	s := &History{
		Data:  nil,
		Mutex: sync.RWMutex{},
	}

	return s
}

var _ = Describe("History methods", func() {
	It("Storing Data", func() {
		Storage := initHistory()
		Storage.Store("apple", "gapple")
		res := Storage.Load("apple")
		Expect(res).To(BeEquivalentTo("gapple"))
	})

	It("Sorting Data", func() {
		Storage := initHistory()
		ew := []string{"jumps", "dog", "brown", "lazy", "fox", "apple"}
		gw := []string{"umpsjogo", "ogdogo", "ownbrogo", "azylogo", "oxfogo", "gapple"}

		for i := 0; i < len(ew); i++ {
			Storage.Store(ew[i], gw[i])
		}

		res := Storage.GetOrderedMap()
		ewRes := make([]string, 0, len(res))
		for val, _ := range res {
			ewRes = append(ewRes, val)
		}
		sort.Strings(ewRes)
		sort.Strings(ew)

		Expect(ewRes).To(BeEquivalentTo(ew))
	})
})
