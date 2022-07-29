package main

import (
	"fmt"
	"github.com/lazymank54/golang-backroom/tax"
)

func main() {

	//ctx := myCtx()
	ctx := demoCtx()

	result := tax.CalcYearsTax(ctx)

	for i := 1; i <= 13; i++ {
		if i == 13 {
			fmt.Printf("全年总纳税额是:%.2f元\n", float64(result[i])/100)
		} else {
			fmt.Printf("%d月纳税额是:%.2f元\n", i, float64(result[i])/100)
		}
	}

}
func demoCtx() tax.Context {

	return tax.Context{
		SalaryBase: 3000000,
		SalaryDetailMap: map[int]int32{
			1: 3050000,
			5: 3090000,
		},
		SocialSecurityAmount1:     400000,
		SocialSecurityAmount2:     450000,
		SocialSecurityChangeMonth: 7,
		SpecialDeductionBase:      250000,
		SpecialDeductionDetailMap: map[int]int32{
			1: 150000,
			3: 150000,
		},
	}
}
