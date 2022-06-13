package transaction

type GlobalTransactionRole int8

const (
	LAUNCHER    GlobalTransactionRole = 0
	PARTICIPANT GlobalTransactionRole = 1
)

type TransactionManager interface {
	// GlobalStatusBegin a new global transaction.
	Begin(applicationId, transactionServiceGroup, name string, timeout int64) (string, error)

	// Global commit.
	Commit(xid string) (GlobalStatus, error)

	//Global rollback.
	Rollback(xid string) (GlobalStatus, error)

	// Get current status of the give transaction.
	GetStatus(xid string) (GlobalStatus, error)

	// Global report.
	GlobalReport(xid string, globalStatus GlobalStatus) (GlobalStatus, error)
}