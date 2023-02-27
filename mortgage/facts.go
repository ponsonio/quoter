package mortgage

type PaymentsPerMonth int64

const (
	Biweekly PaymentsPerMonth = 2
	Monthly  PaymentsPerMonth = 1
)

type Calc struct {
	TotalPropertyValue      float64
	AnualInterestRate       float64
	PaymentsPerMonth        PaymentsPerMonth
	AmortizationPeriodYears int32
	DownPayment             float64
	Valid                   bool
	RequiresCMHC            bool
	CMHCRate                float64
	CMHCAmount              float64
	PaymentPerSchedule      float64
	Errors                  []string
}

func (mc *Calc) AddError(err string) {
	mc.Errors = append(mc.Errors, err)
}
