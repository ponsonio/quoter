package mortgage

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type CalculatorService interface {
	Execute(mortgageCalc *Calc) error
}

type calculatorServiceRules struct {
	knowledgeLibrary ast.KnowledgeLibrary
}

func NewCalculatorService() CalculatorService {
	ret := &calculatorServiceRules{}
	ret.buildRuleEngine()
	return ret
}

func (srv *calculatorServiceRules) buildRuleEngine() {
	srv.knowledgeLibrary = *ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(&srv.knowledgeLibrary)

	// Read rule from file and build rules
	ruleFile := pkg.NewFileResource("./mortgage/rules.grl")
	err := ruleBuilder.BuildRuleFromResource("MortgageRules", "0.0.1", ruleFile)
	if err != nil {
		panic(err)
	}

}

func (srv *calculatorServiceRules) Execute(mortgageCalc *Calc) error {
	dataCtx := ast.NewDataContext()
	mortgageCalc.Valid = true
	mortgageCalc.RequiresCMHC = false
	err := dataCtx.Add("Calc", mortgageCalc)
	if err != nil {
		panic(err)
	}

	knowledgeBase := srv.knowledgeLibrary.NewKnowledgeBaseInstance("MortgageRules", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		panic(err)
	}
	return nil
}
