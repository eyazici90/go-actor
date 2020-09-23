package rabbitmq

func returnOnErr(actions ...func() error) error {
	for _, act := range actions {
		if e := act(); e != nil {
			return e
		}
	}
	return nil
}
