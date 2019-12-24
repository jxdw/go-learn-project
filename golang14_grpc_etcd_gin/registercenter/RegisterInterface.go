package registercenter

type RegisterInterface interface {
	Register(info ServiceDescInfo) error
	Unregister(info ServiceDescInfo) error
}
