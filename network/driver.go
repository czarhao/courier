package network

type Driver interface {
	Name() string
	Create(sub, name string) (*Net, error)
	Delete(new *Net) error

	Connect(net *Net, endpoint *Endpoint) error
	Disconnect(net *Net, endpoint *Endpoint) error
}
