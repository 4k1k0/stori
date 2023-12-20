package transaction

type Collection []Transaction

func (c Collection) Validate() error {
	if len(c) == 0 {
		return nil
	}

	var err error

	for i := 0; i < len(c); i++ {
		err = c[i].Validate()
		if err != nil {
			break
		}
	}

	return err
}
