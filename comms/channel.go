package comms

type Channel interface {
	Get(url string) ([]byte, error)
}
