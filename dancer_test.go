package dancer_test

import (
	. "github.com/vertis/dancer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dancer", func() {
  var dancer *Dancer
  
  BeforeEach(func() {
      dancer = NewDancer()
  })
  
  It("should have an empty set of hosts", func() {
      Expect(dancer.Hosts).To(Equal([]string{}))
  })    
})
