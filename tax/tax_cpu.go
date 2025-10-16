package tax

import "fmt"

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
type taxInfo struct {
	month     int32
	taxAmount int64
	taxRate   int32
}

func CalcAndPrint() {
	//ctx := myCtx()
	ctx := demoCtx()

	result := CalcYearsTax(ctx)

	for i := 1; i <= 14; i++ {
		if i == 14 {
			fmt.Printf("全年总纳税额是:%.2f元\n", float64(result[i].taxAmount)/100)
		} else {
			fmt.Printf("%d月纳税额是:%.2f元, 纳税费率是：%d%%\n", i, float64(result[i].taxAmount)/100, result[i].taxRate)
		}
	}
}

func CalcYearsTax(ctx Context) map[int]*taxInfo {

	var result = make(map[int]*taxInfo)
	var totalSalary, totalSecurity, totalDeduction, totalTax int64 = 0, 0, 0, 0

	for i := 1; i <= 13; i++ {
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
		if i == 13 {
			monthSecurity = 0
			monthDeduction = 0
		}
		totalSalary = totalSalary + int64(monthSalary)
		totalSecurity = totalSecurity + int64(monthSecurity)
		totalDeduction = totalDeduction + int64(monthDeduction)
		num := i
		if i == 13 {
			num = i - 1
		}
		totalBase := int64(num) * int64(baseSalary)

		rate, tax := calcTax(totalSalary, totalBase, totalSecurity, totalDeduction)
		taxInfo := &taxInfo{
			month:     int32(i),
			taxAmount: tax - totalTax,
			taxRate:   rate,
		}

		result[i] = taxInfo
		totalTax += result[i].taxAmount
	}

	result[14] = &taxInfo{
		month:     14,
		taxAmount: totalTax,
		taxRate:   0,
	}

	return result

}

func calcTax(totalSalary int64, totalBase int64, totalSecurity int64, totalDeduction int64) (int32, int64) {

	taxBaseAmount := totalSalary - totalBase - totalSecurity - totalDeduction

	if taxBaseAmount < 500000 {
		return 0, 0
	} else if taxBaseAmount <= 3600000 {
		return 3, taxBaseAmount * 3 / 100
	} else if taxBaseAmount <= 14400000 {
		return 10, taxBaseAmount*10/100 - 252000
	} else if taxBaseAmount <= 30000000 {
		return 20, taxBaseAmount*20/100 - 1692000
	} else if taxBaseAmount <= 42000000 {
		return 25, taxBaseAmount*25/100 - 3192000
	} else if taxBaseAmount <= 66000000 {
		return 30, taxBaseAmount*30/100 - 5292000
	} else if taxBaseAmount <= 96000000 {
		return 35, taxBaseAmount*35/100 - 8592000
	} else {
		return 45, taxBaseAmount*45/100 - 18192000
	}
}
func demoCtx() Context {

	return Context{
		SalaryBase: ,
		SalaryDetailMap: map[int]int32{
			1:  ,
			3: ,
			13: ,
		},
		SocialSecurityAmount1:     ,
		SocialSecurityAmount2:     ,
		SocialSecurityChangeMonth: 7,
		SpecialDeductionBase:      450000,
		SpecialDeductionDetailMap: map[int]int32{},
	}
}
