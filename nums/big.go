package nums

import (
	"github.com/ericlagergren/decimal"
)

// BigSeperateFractionalPotions decompose a decimal number with a sign, integer, and fractional parts.
func BigSeperateFractionalPotions(src *decimal.Big) (sign bool, integer, fractional uint64, fractionalLength uint8) {
	var num, integerPart decimal.Big
	num.Copy(src)
	num.Context.RoundingMode = decimal.ToZero
	sign = num.Signbit()
	integer, _ = num.Abs(&num).Uint64()
	integerPart.SetUint64(integer)
	scale := num.Sub(&num, &integerPart).Quantize(6).Reduce().Scale()
	fractional, _ = num.Mul(&num, decimal.New(1, -scale)).Reduce().Uint64()
	fractionalLength = uint8(scale)
	return
}

// BigMergeFractionalPotions compose a decimal number with a sign, integer, and fractional parts.
func BigMergeFractionalPotions(sign bool, integer, fractional uint64, fractionalLength uint8) (dst *decimal.Big) {
	var num, frac decimal.Big
	num.SetUint64(integer)
	frac.SetUint64(fractional)
	dst = &num

	num.Mul(&num, decimal.New(1, -int(fractionalLength))).Add(&num, &frac).SetScale(int(fractionalLength))

	if sign {
		num.Neg(&num)
	}
	num.Reduce()
	return
}
