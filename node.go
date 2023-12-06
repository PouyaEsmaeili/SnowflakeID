package snowflake

import (
	"math"
	"time"
)

type NodeInterface interface {
	Generate() ID
	ConvertToID(id int64) (*ID, error)
}

type Node struct {
	machineId          int16
	name               string
	startSeqFrom       int16
	startTimestampFrom int64
	currentSeqValue    int16
}

func (n *Node) Generate() ID {
	currentTime := time.Now().Unix() - n.startTimestampFrom
	seqNum := n.currentSeqValue + 1
	newId := ID{
		sequence:  seqNum,
		machineId: n.machineId,
		timestamp: currentTime,
	}
	n.currentSeqValue = seqNum
	return newId
}

func (n *Node) ConvertToID(id int64) (*ID, error) {
	sequence := id % int64(math.Pow(2, 12))
	if int16(sequence) < minSequenceNumber {
		return nil, ErrInvalidSequence
	}
	id = (id - sequence) / int64(math.Pow(2, 12))
	machineId := id % int64(math.Pow(2, 10))
	if int16(machineId) != n.machineId {
		return nil, ErrInvalidMachineId
	}
	timestamp := (id - machineId) / int64(math.Pow(2, 10))
	if timestamp < n.startTimestampFrom {
		return nil, ErrInvalidTimestamp
	}
	snowflakeId := ID{
		sequence:  int16(sequence),
		machineId: int16(machineId),
		timestamp: timestamp,
	}
	return &snowflakeId, nil
}
