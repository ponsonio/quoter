package mortgage

type PaymentsPerMonth int64

const (
	Biweekly PaymentsPerMonth = 0
	Monthly  PaymentsPerMonth = 1
)

type Calc struct {
	TotalPropertyValue float64
	PaymentsPerMonth   PaymentsPerMonth
	AmortizationPeriod int32
	DownPayment        float64
	Period             int32
	Valid              bool
	RequiresCMHC       bool
	CMHC               float32
	PaymentAmount      float64
	Errors             []string
}

func (mc *Calc) AddError(err string) {
	mc.Errors = append(mc.Errors, err)
}
