package mortgage

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidations(t *testing.T) {
	service := NewCalculatorService("rules.grl")

	type TestCase struct {
		Name          string
		Calc          *Calc
		ExpectedError string
	}
	tc := []TestCase{
		{
			Name: "Not enough Years",
			Calc: &Calc{
				TotalPropertyValue:      500000.00,
				AnualInterestRate:       0.06,
				PaymentsPerMonth:        1,
				AmortizationPeriodYears: 4,
				DownPayment:             100000,
			},
			ExpectedError: "AmortizationPeriod must be between 5 and 30 years",
		},
		{
			Name: "To many Years",
			Calc: &Calc{
				TotalPropertyValue:      500000.00,
				AnualInterestRate:       0.06,
				PaymentsPerMonth:        1,
				AmortizationPeriodYears: 31,
				DownPayment:             100000,
			},
			ExpectedError: "AmortizationPeriod must be between 5 and 30 years",
		},
		{
			Name: "Invalid interest rate",
			Calc: &Calc{
				TotalPropertyValue:      500000.00,
				AnualInterestRate:       0.0,
				PaymentsPerMonth:        1,
				AmortizationPeriodYears: 5,
				DownPayment:             100000,
			},
			ExpectedError: "Interest Rate must be more than 0.00",
		},
		{
			Name: "Invalid interest rate",
			Calc: &Calc{
				TotalPropertyValue:      500000.00,
				AnualInterestRate:       0.0,
				PaymentsPerMonth:        1,
				AmortizationPeriodYears: 5,
				DownPayment:             100000,
			},
			ExpectedError: "Interest Rate must be more than 0.00",
		},
		{
			Name: "At least 20 if more that 1M",
			Calc: &Calc{
				TotalPropertyValue:      10000001.00,
				AnualInterestRate:       0.06,
				PaymentsPerMonth:        1,
				AmortizationPeriodYears: 5,
				DownPayment:             100000,
			},
			ExpectedError: "DownPayment should be at least 20% for a 1,000,000 property",
		},
		{
			Name: "Min DP is 5%",
			Calc: &Calc{
				TotalPropertyValue:      499000.00,
				AnualInterestRate:       0.06,
				PaymentsPerMonth:        1,
				AmortizationPeriodYears: 5,
				DownPayment:             24000,
			},
			ExpectedError: "Minimum Down Payment is 5%, for properties less than 500000",
		},
	}
	for _, input := range tc {
		err := service.Execute(input.Calc)
		if err != nil {
			t.Fatal(fmt.Sprintf("test fail: %s, unexepected error during tests", input.Name), err)
		}

		if input.Calc.Valid == true {
			t.Fatal(fmt.Sprintf("test fail: %s, no error found", input.Name))
		}

		if !strings.Contains(input.ExpectedError, strings.Join(input.Calc.Errors, ",")) {
			t.Fatal(fmt.Sprintf("test fail: %s, expected error (%s) not found", input.Name, input.ExpectedError))
		}
	}

}
