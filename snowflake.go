package snowflake

type ArgBuilder = func(*Node) error

func SetName(name string) ArgBuilder {
	return func(node *Node) error {
		node.name = name
		return nil
	}
}

func SetStartSeqFrom(initialValue int16) ArgBuilder {
	return func(node *Node) error {
		if initialValue < minSequenceNumber {
			return ErrSeqMinValue
		}
		if initialValue > maxSequenceNumber {
			return ErrSeqMaxValue
		}
		node.startSeqFrom = initialValue
		node.currentSeqValue = initialValue
		return nil
	}
}

func SetStartTimestampFrom(initialValue int64) ArgBuilder {
	return func(node *Node) error {
		if initialValue < minTimestamp {
			return ErrTimestampMinValue
		}
		if initialValue > maxTimestamp {
			return ErrTimestampMaxValue
		}
		node.startTimestampFrom = initialValue
		return nil
	}
}

func validateMachineId(machineId int16) error {
	if machineId < minMachineId {
		return ErrMachineIdMinValue
	}
	if machineId > maxMachineId {
		return ErrMachineIdMaxValue
	}
	return nil
}

func New(machineId int16, options ...ArgBuilder) (*Node, error) {
	validationResult := validateMachineId(machineId)
	if validationResult != nil {
		return nil, validationResult
	}
	returnValue := &Node{
		machineId: machineId,
	}

	for _, option := range options {
		err := option(returnValue)
		if err != nil {
			return nil, err
		}
	}

	return returnValue, nil
}
