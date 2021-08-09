package clients

type Clients interface {
	GetNumberOfPods(namespace string) int
}
