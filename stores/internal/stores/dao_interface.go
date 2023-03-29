package stores

type IChashierPoolService interface {
	CashierOrderCompletedDAO(orderRequestChan <-chan map[string]interface{})
}
