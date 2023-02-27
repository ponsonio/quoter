package mortgage

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"math"
)

// CalculatorService
//using an interface here, will allow to change the implementation (if you don't like the rule stuff)
type CalculatorService interface {
	Execute(mortgageCalc *Calc) error
}

type calculatorServiceRules struct {
	knowledgeLibrary ast.KnowledgeLibrary
}

func NewCalculatorService(rulesDir string) CalculatorService {
	ret := &calculatorServiceRules{}
	ret.buildRuleEngine(rulesDir)
	return ret
}

func (srv *calculatorServiceRules) buildRuleEngine(rulesDir string) {
	srv.knowledgeLibrary = *ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(&srv.knowledgeLibrary)

	// Read rule from file and build rules
	ruleFile := pkg.NewFileResource(rulesDir)
	err := ruleBuilder.BuildRuleFromResource("MortgageRules", "0.0.1", ruleFile)
	if err != nil {
		panic(err) // seems ok to panic here, as service won't  work, so we need to fix it
	}

}

func (srv *calculatorServiceRules) Execute(mortgageCalc *Calc) error {
	dataCtx := ast.NewDataContext()
	mortgageCalc.Valid = true
	mortgageCalc.RequiresCMHC = false
	mortgageCalc.CMHCRate = 0.0
	fmt.Printf("down payment percentaje:%v\n", mortgageCalc.DownPayment/mortgageCalc.TotalPropertyValue)
	err := dataCtx.Add("Calc", mortgageCalc)
	if err != nil {
		return fmt.Errorf("error executing mortgage validation :%s", err)
	}

	knowledgeBase := srv.knowledgeLibrary.NewKnowledgeBaseInstance("MortgageRules", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		return fmt.Errorf("error executing mortgage validation :%s", err)
	}

	if mortgageCalc.Valid {
		mortgageCalc.PaymentPerSchedule = calculatePaymentPerSchedule(mortgageCalc)
	} else {
		mortgageCalc.PaymentPerSchedule = -1.00
	}

	return nil
}

// Even the engine is capable of executing the calculations, code was preferred just because it's easy to debug and
// understand, console log are left intentionally just to trace easy
func calculatePaymentPerSchedule(c *Calc) float64 {
	p := c.TotalPropertyValue - c.DownPayment + c.CMHCAmount
	fmt.Printf("principal:%v\n", p)
	r := c.AnualInterestRate / (12 * float64(c.PaymentsPerMonth))
	fmt.Printf("r:%v\n", r)
	n := float64((12 * c.AmortizationPeriodYears) * int32(c.PaymentsPerMonth))
	fmt.Printf("n:%v\n", n)
	res := p * ((r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1))
	fmt.Printf("res:%v\n", res)
	res = math.Floor(res*100) / 100
	fmt.Printf("res rouded :%v\n", res)
	return res
}
