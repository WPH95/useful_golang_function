func retry(attempts int, callback func() error, check func(error) bool) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return nil
		}
		if !check(err) {
			return err
		}

		if i >= (attempts - 1) {
			log.Error("Retry too many time")
			break
		}

		time.Sleep(1 * time.Second)

		log.Warn("retrying...")
	}
	return err
}


err = c.createNode(nodeID, newWorkflow)
  if err != nil {
    return retry(5,
      func() error { return c.createNode(nodeID, newWorkflow) },
      waitingDeleteFinished)
