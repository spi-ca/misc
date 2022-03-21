package monetary

import (
	"github.com/ericlagergren/decimal"
	"math/big"
)

// MonthlyInstallmentInterestAmount returns monthly interest amount from annual interest ratio.
func MonthlyInstallmentInterestAmount(balance *decimal.Big, annualInterest *big.Rat) (fee *decimal.Big) {

	var (
		monthly big.Rat
		feeFrac big.Rat
	)
	monthly.SetFrac64(1, 12)
	balance.Rat(&feeFrac)
	feeFrac.Mul(&feeFrac, annualInterest)
	feeFrac.Mul(&feeFrac, &monthly)
	fee = decimal.New(0, 0)
	fee.Context.RoundingMode = decimal.ToZero
	fee.SetRat(&feeFrac)
	fee.Quantize(0)
	fee.Reduce()
	return fee
}

// MonthlyInstallmentInfo returns monthly repayment plan.
func MonthlyInstallmentInfo(totalBalance *decimal.Big, modDigit, division int) (installment, installmentFractional, installmentExtra *decimal.Big) {
	var (
		monthlyFee   big.Rat
		frac         big.Rat
		divisionDeci decimal.Big
	)

	divisionDeci.SetMantScale(int64(division), 0)
	totalBalance.Rat(&monthlyFee)
	frac.SetFrac64(1, int64(division))
	monthlyFee.Mul(&monthlyFee, &frac)

	installment = decimal.New(0, 0)
	installment.Context.RoundingMode = decimal.ToZero
	installment.SetRat(&monthlyFee)
	installment.Quantize(-modDigit)
	installment.Reduce()

	installmentFractional = decimal.New(0, 0)
	installmentFractional.Mul(installment, &divisionDeci)
	installmentFractional.Sub(totalBalance, installmentFractional)
	installmentFractional.Reduce()

	installmentExtra = decimal.New(0, 0)
	installmentExtra.Add(installment, installmentFractional)
	installmentExtra.Reduce()

	installmentTest := decimal.New(0, 0)
	installmentTest.Mul(installment, &divisionDeci)
	installmentTest.Add(installmentTest, installmentFractional)
	return
}

// MonthlyInstallmentSchedule returns detailed monthly repayment plan
func MonthlyInstallmentSchedule(
	totalBalance *decimal.Big,
	annualInterest *decimal.Big,
	modDigit, division int,
	payExtraAmountEarlier bool) (
	installment, installmentExtra, totalInterests *decimal.Big,
	principleBalanceBeforePayments,
	principleBalanceAfterPayments,
	installments,
	interests,
	schedules []*decimal.Big,
) {
	if division < 1 {
		return
	} else if totalBalance == nil {
		return
	}
	var (
		leftBalance        decimal.Big
		annualInterestFrac big.Rat
	)

	if annualInterest != nil {
		annualInterest.Rat(&annualInterestFrac)
	}

	totalInterests = decimal.New(0, 0)
	leftBalance.Copy(totalBalance)
	installmentAmount, _, installmentAmountExtra := MonthlyInstallmentInfo(totalBalance, modDigit, division)
	principleBalanceBeforePayments, principleBalanceAfterPayments, installments, interests, schedules = make([]*decimal.Big, 0, division), make([]*decimal.Big, 0, division), make([]*decimal.Big, 0, division), make([]*decimal.Big, 0, division), make([]*decimal.Big, 0, division)
	for i := 0; i < division; i++ {
		var (
			interestAmount                 = MonthlyInstallmentInterestAmount(&leftBalance, &annualInterestFrac)
			monthlyInstallmentWithInterest decimal.Big
			principleBalanceBeforePayment  decimal.Big
			principleBalanceAfterPayment   decimal.Big
			paymentAmount                  *decimal.Big
		)
		if payExtraAmountEarlier {
			if i == 0 {
				paymentAmount = installmentAmountExtra
			} else {
				paymentAmount = installmentAmount
			}
		} else {
			if i+1 == division {
				paymentAmount = installmentAmountExtra
			} else {
				paymentAmount = installmentAmount
			}
		}
		installments = append(installments, paymentAmount)
		interests = append(interests, interestAmount)
		totalInterests.Add(totalInterests, interestAmount)
		principleBalanceBeforePayment.Copy(&leftBalance)
		principleBalanceBeforePayments = append(principleBalanceBeforePayments, &principleBalanceBeforePayment)
		monthlyInstallmentWithInterest.Add(paymentAmount, interestAmount)
		leftBalance.Sub(&leftBalance, paymentAmount)
		principleBalanceAfterPayment.Copy(&leftBalance)
		principleBalanceAfterPayments = append(principleBalanceAfterPayments, &principleBalanceAfterPayment)
		schedules = append(schedules, &monthlyInstallmentWithInterest)
	}
	return
}
