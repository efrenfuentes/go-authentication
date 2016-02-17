package user_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/efrenfuentes/go-authentication/models"
)

var _ = Describe("User Model", func() {
	It("Valid mail", func() {
		var user *models.User
		user = new(models.User)

		valid_emails := []string{"user@foo.COM",
			                     "A_US-ER@f.b.org",
			                     "frst.lst@foo.jp"}

		for _, email := range valid_emails {
			err := user.SetEmail(email)

			Expect(err).To(BeNil())

			Expect(email).To(Equal(user.GetEmail()))
		}
	})

	It("Invalid mail", func() {
		var user *models.User
		user = new(models.User)

		invalid_emails := []string{"user@foo,com",
		                           "user_at_foo.org",
			                       "example.user@foo.",
			                       "foo@bar+baz.com",
			                       "foo@bar..com"}

		for _, email := range invalid_emails {
			before_email := user.GetEmail()

			err := user.SetEmail(email)

			Expect(err).NotTo(BeNil())

			Expect(user.GetEmail()).To(Equal(before_email))
		}
	})
})
