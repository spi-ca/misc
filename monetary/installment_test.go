package monetary

import (
	"github.com/ericlagergren/decimal"
	"math/big"
	"testing"
)

func TestInstallment(t *testing.T) {
	var (
		charge                       decimal.Big
		annualInterest               decimal.Big
		annualInterestFrac           big.Rat
		annualInterestFracPercentage big.Rat
		modDigit                     = 2
		division                     = 36
	)
	charge.SetUint64(10_000_000)
	annualInterest.SetMantScale(125, 3)
	annualInterest.Rat(&annualInterestFrac)
	annualInterestFracPercentage.Mul(&annualInterestFrac, big.NewRat(100, 1))
	chargeInt, _ := charge.Int64()

	func() {
		interestFee := MonthlyInstallmentInterestAmount(&charge, &annualInterestFrac)
		interestFeeInt, _ := interestFee.Int64()
		t.Logf("할부원금 %d 년이율 %s%% 납부 이자 %d",
			chargeInt,
			annualInterestFracPercentage.FloatString(2),
			interestFeeInt,
		)
	}()

	func() {
		installment, installmentFractional, installmentWithFractional := MonthlyInstallmentInfo(&charge, modDigit, division)
		installmentInt, _ := installment.Int64()
		installmentFractionalInt, _ := installmentFractional.Int64()
		installmentWithFractionalInt, _ := installmentWithFractional.Int64()
		t.Logf("할부총원금 %d 월 할부 절사자리수 %d, %d 개월동안, 월납입액 %d, 우수리 %d 우수리 포함 월납입액 %d",
			chargeInt,
			modDigit,
			division,
			installmentInt,
			installmentFractionalInt,
			installmentWithFractionalInt,
		)
	}()
	func() {
		payExtraAmountEarlier := true
		_, _, totalInterests,
			principleBalanceBeforePayments,
			principleBalanceAfterPayments,
			installments,
			interests,
			schedules := MonthlyInstallmentSchedule(&charge, nil, modDigit, division, payExtraAmountEarlier)
		totalInterestsInt, _ := totalInterests.Int64()

		t.Logf("할부총원금 %d 년이율 %s%%, 월 할부 절사자리수 %d, %d 개월동안, 우수리 선납 %v 총 납부 이자 %d",
			chargeInt,
			"0",
			modDigit,
			division,
			payExtraAmountEarlier,
			totalInterestsInt,
		)
		for i := 0; i < division; i++ {
			principleBalanceBeforePayments, principleBalanceAfterPayments, installment, interest, schedule := principleBalanceBeforePayments[i], principleBalanceAfterPayments[i], installments[i], interests[i], schedules[i]
			principleBalanceBeforePaymentsInt, _ := principleBalanceBeforePayments.Int64()
			installmentInt, _ := installment.Int64()
			interestInt, _ := interest.Int64()
			scheduleInt, _ := schedule.Int64()
			principleBalanceAfterPaymentsInt, _ := principleBalanceAfterPayments.Int64()

			t.Logf("%d번쨰 할부 납부 도래시, 납부전 할부원금 %d -> 월총금액 (월납입액 %d + 할부수수료 %d)=%d -> 납부 할부원금 %d",
				i+1,
				principleBalanceBeforePaymentsInt,
				installmentInt,
				interestInt,
				scheduleInt,
				principleBalanceAfterPaymentsInt,
			)
		}
	}()

	func() {
		payExtraAmountEarlier := false
		_, _, totalInterests,
			principleBalanceBeforePayments,
			principleBalanceAfterPayments,
			installments,
			interests,
			schedules := MonthlyInstallmentSchedule(&charge, nil, modDigit, division, payExtraAmountEarlier)
		totalInterestsInt, _ := totalInterests.Int64()

		t.Logf("할부총원금 %d 년이율 %s%%, 월 할부 절사자리수 %d, %d 개월동안, 우수리 선납 %v 총 납부 이자 %d",
			chargeInt,
			"0",
			modDigit,
			division,
			payExtraAmountEarlier,
			totalInterestsInt,
		)
		for i := 0; i < division; i++ {
			principleBalanceBeforePayments, principleBalanceAfterPayments, installment, interest, schedule := principleBalanceBeforePayments[i], principleBalanceAfterPayments[i], installments[i], interests[i], schedules[i]
			principleBalanceBeforePaymentsInt, _ := principleBalanceBeforePayments.Int64()
			installmentInt, _ := installment.Int64()
			interestInt, _ := interest.Int64()
			scheduleInt, _ := schedule.Int64()
			principleBalanceAfterPaymentsInt, _ := principleBalanceAfterPayments.Int64()

			t.Logf("%d번쨰 할부 납부 도래시, 납부전 할부원금 %d -> 월총금액 (월납입액 %d + 할부수수료 %d)=%d -> 납부 할부원금 %d",
				i+1,
				principleBalanceBeforePaymentsInt,
				installmentInt,
				interestInt,
				scheduleInt,
				principleBalanceAfterPaymentsInt,
			)
		}
	}()

	func() {
		payExtraAmountEarlier := true
		_, _, totalInterests,
			principleBalanceBeforePayments,
			principleBalanceAfterPayments,
			installments,
			interests,
			schedules := MonthlyInstallmentSchedule(&charge, &annualInterest, modDigit, division, payExtraAmountEarlier)
		totalInterestsInt, _ := totalInterests.Int64()

		t.Logf("할부총원금 %d 년이율 %s%%, 월 할부 절사자리수 %d, %d 개월동안, 우수리 선납 %v 총 납부 이자 %d",
			chargeInt,
			annualInterestFracPercentage.FloatString(2),
			modDigit,
			division,
			payExtraAmountEarlier,
			totalInterestsInt,
		)
		for i := 0; i < division; i++ {
			principleBalanceBeforePayments, principleBalanceAfterPayments, installment, interest, schedule := principleBalanceBeforePayments[i], principleBalanceAfterPayments[i], installments[i], interests[i], schedules[i]
			principleBalanceBeforePaymentsInt, _ := principleBalanceBeforePayments.Int64()
			installmentInt, _ := installment.Int64()
			interestInt, _ := interest.Int64()
			scheduleInt, _ := schedule.Int64()
			principleBalanceAfterPaymentsInt, _ := principleBalanceAfterPayments.Int64()

			t.Logf("%d번쨰 할부 납부 도래시, 납부전 할부원금 %d -> 월총금액 (월납입액 %d + 할부수수료 %d)=%d -> 납부 할부원금 %d",
				i+1,
				principleBalanceBeforePaymentsInt,
				installmentInt,
				interestInt,
				scheduleInt,
				principleBalanceAfterPaymentsInt,
			)
		}
	}()

	func() {
		payExtraAmountEarlier := false
		_, _, totalInterests,
			principleBalanceBeforePayments,
			principleBalanceAfterPayments,
			installments,
			interests,
			schedules := MonthlyInstallmentSchedule(&charge, &annualInterest, modDigit, division, payExtraAmountEarlier)
		totalInterestsInt, _ := totalInterests.Int64()

		t.Logf("할부총원금 %d 년이율 %s%%, 월 할부 절사자리수 %d, %d 개월동안, 우수리 선납 %v 총 납부 이자 %d",
			chargeInt,
			annualInterestFracPercentage.FloatString(2),
			modDigit,
			division,
			payExtraAmountEarlier,
			totalInterestsInt,
		)
		for i := 0; i < division; i++ {
			principleBalanceBeforePayments, principleBalanceAfterPayments, installment, interest, schedule := principleBalanceBeforePayments[i], principleBalanceAfterPayments[i], installments[i], interests[i], schedules[i]
			principleBalanceBeforePaymentsInt, _ := principleBalanceBeforePayments.Int64()
			installmentInt, _ := installment.Int64()
			interestInt, _ := interest.Int64()
			scheduleInt, _ := schedule.Int64()
			principleBalanceAfterPaymentsInt, _ := principleBalanceAfterPayments.Int64()

			t.Logf("%d번쨰 할부 납부 도래시, 납부전 할부원금 %d -> 월총금액 (월납입액 %d + 할부수수료 %d)=%d -> 납부 할부원금 %d",
				i+1,
				principleBalanceBeforePaymentsInt,
				installmentInt,
				interestInt,
				scheduleInt,
				principleBalanceAfterPaymentsInt,
			)
		}
	}()

}
