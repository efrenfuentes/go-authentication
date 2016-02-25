package crypt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCrypt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crypt Suite")
}
