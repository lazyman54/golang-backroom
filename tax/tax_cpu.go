package tax

var baseSalary int32 = 500000

type Context struct {
	/*金额单位统一用分*/
	SalaryBase                int32         //每月的月薪
	SalaryDetailMap           map[int]int32 //如果某个月多发了钱，可以这里填写，比如月薪30000元，但是3月发了33000元，则这里写3：33000元
	SocialSecurityAmount1     int32         //五险一金累计扣除金额（上半年）
	SocialSecurityChangeMonth int32         //五险一金金额更新的月份，一般是7月，如果未更新，这里是0
	SocialSecurityAmount2     int32         //五险一金累计扣除金额（下半年），如果7月未更新，这里不填
	SpecialDeductionBase      int32         //专项扣除金额
	SpecialDeductionDetailMap map[int]int32 //专项扣除里，如果某个月扣除数不对，可以额外指定，key=月份，value=当月的值
}

func CalcYearsTax(ctx Context) map[int]int64 {

	var result = make(map[int]int64)
	var totalSalary, totalSecurity, totalDeduction, totalTax int64 = 0, 0, 0, 0

	for i := 1; i <= 12; i++ {
		monthSalary := ctx.SalaryBase
		if salary, ok := ctx.SalaryDetailMap[i]; ok {
			monthSalary = salary
		}
		monthSecurity := ctx.SocialSecurityAmount1
		if ctx.SocialSecurityChangeMonth != 0 && int32(i) >= ctx.SocialSecurityChangeMonth {
			monthSecurity = ctx.SocialSecurityAmount2
		}

		monthDeduction := ctx.SpecialDeductionBase
		if deduction, ok := ctx.SpecialDeductionDetailMap[i]; ok {
			monthDeduction = deduction
		}

		totalSalary = totalSalary + int64(monthSalary)
		totalSecurity = totalSecurity + int64(monthSecurity)
		totalDeduction = totalDeduction + int64(monthDeduction)
		totalBase := int64(i) * int64(baseSalary)

		tax := calcTax(totalSalary, totalBase, totalSecurity, totalDeduction)

		result[i] = tax - totalTax
		totalTax += result[i]
	}
	result[13] = totalTax

	return result

}

func calcTax(totalSalary int64, totalBase int64, totalSecurity int64, totalDeduction int64) int64 {

	taxBaseAmount := totalSalary - totalBase - totalSecurity - totalDeduction

	if taxBaseAmount < 500000 {
		return 0
	} else if taxBaseAmount <= 3600000 {
		return taxBaseAmount * 3 / 100
	} else if taxBaseAmount <= 14400000 {
		return taxBaseAmount*10/100 - 252000
	} else if taxBaseAmount <= 30000000 {
		return taxBaseAmount*20/100 - 1692000
	} else if taxBaseAmount <= 42000000 {
		return taxBaseAmount*25/100 - 3192000
	} else if taxBaseAmount <= 66000000 {
		return taxBaseAmount*30/100 - 5292000
	} else if taxBaseAmount <= 96000000 {
		return taxBaseAmount*35/100 - 8592000
	} else {
		return taxBaseAmount*45/100 - 18192000
	}

}
