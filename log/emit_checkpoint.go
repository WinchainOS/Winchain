package log

//TODO【Winchain】================create by myself==================
const (
	TxCreated          = "TX-CREATED"
	TxAccepted         = "TX-ACCEPTED"
	BecameMinter       = "BECAME-MINTER"
	BecameVerifier     = "BECAME-VERIFIER"
	BlockCreated       = "BLOCK-CREATED"
	BlockVotingStarted = "BLOCK-VOTING-STARTED"
)

//TODO【Winchain】================the part had been changed==================
var DoEmitCheckpoints = false

func EmitCheckpoint(checkpointName string, logValues ...interface{}) {
	args := []interface{}{"name", checkpointName}
	args = append(args, logValues...)
	if DoEmitCheckpoints {
		Info("Winchain-CHECKPOINT", args...)
	}
}
