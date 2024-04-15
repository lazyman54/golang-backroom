package tax

import (
	"fmt"
)

type RewardTax struct{}

var rewardTaxIns = &RewardTax{}

func NewRewardTaxIns() *RewardTax {
	return rewardTaxIns
}

type RateTiers struct {
	levelVal     int64
	taxRate      int64
	deductionVal int64
}

var taxRateTiersList []*RateTiers

func init() {
	firstRate := &RateTiers{
		levelVal:     3000 * 100,
		taxRate:      3,
		deductionVal: 0,
	}
	secondRate := &RateTiers{
		levelVal:     12000 * 100,
		taxRate:      10,
		deductionVal: 210 * 100,
	}

	thirdRate := &RateTiers{
		levelVal:     25000 * 100,
		taxRate:      20,
		deductionVal: 1410 * 100,
	}
	fourRate := &RateTiers{
		levelVal:     35000 * 100,
		taxRate:      25,
		deductionVal: 2660 * 100,
	}
	fiveRate := &RateTiers{
		levelVal:     55000 * 100,
		taxRate:      30,
		deductionVal: 4410 * 100,
	}
	sixRate := &RateTiers{
		levelVal:     80000 * 100,
		taxRate:      35,
		deductionVal: 7160 * 100,
	}
	sevenRate := &RateTiers{
		levelVal:     1000000000 * 100,
		taxRate:      45,
		deductionVal: 15160 * 100,
	}
	taxRateTiersList = append(taxRateTiersList, firstRate, secondRate, thirdRate, fourRate, fiveRate, sixRate, sevenRate)
}

func (cls *RewardTax) GetMyRewardPlan(rewardVal int64, salaryRate int64) {

	tax := cls.allCashRewardTax(rewardVal)
	fmt.Printf("all cash you will pay tax amount:%d元, you can get cash:%d元\n", tax/100, (rewardVal-tax)/100)

	for _, rateTiers := range taxRateTiersList {
		yearAmount := rateTiers.levelVal * 12
		if rewardVal < yearAmount {
			return
		}
		rewardTax := cls.calcTax(yearAmount, rateTiers)
		restAmount := rewardVal - yearAmount
		restTax := restAmount * salaryRate / 100
		totalTax := rewardTax + restTax
		fmt.Printf("take %d 元 cash, and rest:%d元 buy stock, your tax amount is :%d+%d=%d 元, you can get cash:%d 元, help you save tax:%d 元\n",
			yearAmount/100, restAmount/100, rewardTax/100, restTax/100, totalTax/100, (rewardVal-totalTax)/100, (tax-totalTax)/100)
	}

}

func (cls *RewardTax) allCashRewardTax(cashRewardVal int64) int64 {

	return cls.calcTax(cashRewardVal, cls.getRateTiers(cashRewardVal))
}

func (cls *RewardTax) calcTax(amount int64, rateTiers *RateTiers) int64 {

	taxAmount := amount*rateTiers.taxRate/100 - rateTiers.deductionVal

	return taxAmount
}

func (cls *RewardTax) getRateTiers(amount int64) *RateTiers {

	monthRewardVal := amount / 12
	for _, rateTiers := range taxRateTiersList {
		if monthRewardVal <= rateTiers.levelVal {
			return rateTiers
		}
	}
	return taxRateTiersList[len(taxRateTiersList)-1]

}
