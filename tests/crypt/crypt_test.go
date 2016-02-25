package crypt_test

import (
	. "github.com/efrenfuentes/go-authentication/core/crypt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crypt", func() {
	Describe("Crypt function", func() {
		It("encrypt some passwords", func() {
			Expect(Crypt("abcdefg", "aa")).To(Equal("aaTcvO819w3js"))
			Expect(Crypt("1234567", "00")).To(Equal("00eD1Q7PHZ7O."))
			Expect(Crypt("testtest", "es")).To(Equal("esDRYJnY4VaGM"))
		})
	})
})
