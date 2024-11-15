package bank_test

import (
    . "task-bank/bank"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
    var account Account

    BeforeEach(func() {
        account = Account{}
    })

    Describe("Deposit", func() {
        It("should increase balance with a valid deposit", func() {
            err := account.Deposit(100)
            Expect(err).To(BeNil())
            Expect(account.GetBalance()).To(Equal(100.0))
        })

        It("should return an error for zero or negative deposit", func() {
            err := account.Deposit(0)
            Expect(err).To(MatchError("invalid deposit amount"))

            err = account.Deposit(-50)
            Expect(err).To(MatchError("invalid deposit amount"))
        })
    })

    Describe("Withdraw", func() {
        BeforeEach(func() {
            _ = account.Deposit(100)
        })

        It("should decrease balance with a valid withdrawal", func() {
            err := account.Withdraw(50)
            Expect(err).To(BeNil())
            Expect(account.GetBalance()).To(Equal(50.0))
        })

        It("should return an error for insufficient funds", func() {
            err := account.Withdraw(150)
            Expect(err).To(MatchError("insufficient funds"))
        })

        It("should return an error for zero or negative withdrawal", func() {
            err := account.Withdraw(0)
            Expect(err).To(MatchError("invalid withdrawal amount"))

            err = account.Withdraw(-30)
            Expect(err).To(MatchError("invalid withdrawal amount"))
        })
    })

    Describe("GetBalance", func() {
        It("should return the correct balance", func() {
            _ = account.Deposit(100)
            Expect(account.GetBalance()).To(Equal(100.0))
        })
    })
})