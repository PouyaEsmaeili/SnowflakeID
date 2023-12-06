package snowflake

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrSeqMinValue = errors.New(
		fmt.Sprintf("Minimum value for intial sequence number is %d", minSequenceNumber))
	ErrSeqMaxValue = errors.New(
		fmt.Sprintf("Maximum value for initial sequence number is %d", maxSequenceNumber))
	ErrTimestampMinValue = errors.New(
		fmt.Sprintf("Minimum value for intial timestamp is %d", minTimestamp))
	ErrTimestampMaxValue = errors.New(
		fmt.Sprintf("Maximum value for initial timestamp is %d", maxTimestamp))
	ErrMachineIdMinValue = errors.New(
		fmt.Sprintf("Minimum value for intial machine id is %d", minMachineId))
	ErrMachineIdMaxValue = errors.New(
		fmt.Sprintf("Maximum value for initial machine id is %d", maxMachineId))
)

var (
	ErrInvalidSequence = errors.New(
		fmt.Sprintf("Sequence number is invalid. The minimum value is: %d", minSequenceNumber))
	ErrInvalidMachineId = errors.New("MachineId is invalid.")
	ErrInvalidTimestamp = errors.New(
		fmt.Sprintf("Timestamp is invalid. The minimum value is: %d", minTimestamp))
)
