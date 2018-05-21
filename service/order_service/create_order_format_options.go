package order_service

type options struct {
	OrderSn string
}

func newOptions() *options {
	return new(options)
}
