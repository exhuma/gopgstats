package gopgstats

type DatabaseRow struct {
	Name string
}

// --- Global Statistics (available from any DB) -----------------------------

type DiskSizeRow struct {
	DatabaseName string
	Size         uint
}

type LocksRow struct {
	DatabaseName string
	Mode         string
	Type         string
	Granted      bool
	Count        uint
}

type GlobalSizeRow struct {
	DatabaseName string
	Size         uint
}

type QueryAgeRow struct {
	DatabaseName   string
	QueryAge       uint
	TransactionAge uint
}

type TransactionsRow struct {
	DatabaseName string
	Committed    uint
	Rolledback   uint
}

type TempBytesRow struct {
	DatabaseName   string
	TemporaryBytes uint
}

// --- DB detail statistices (must be connected to the respective DB) --------

type DiskIORow struct {
	HeapBlocksRead       uint
	HeapBlocksHit        uint
	IndexBlocksRead      uint
	IndexBlocksHit       uint
	ToastBlocksRead      uint
	ToastBlocksHit       uint
	ToastIndexBlocksRead uint
	ToastIndexBlocksHit  uint
}

type IndexIORow struct {
	IndexBlocksRead uint
	IndexBlocksHit  uint
}

type SequencesIORow struct {
	BlocksRead uint
	BlocksHit  uint
}

type ScanTypesRow struct {
	IndexScans      uint
	SequentialScans uint
}

type RowAccessRow struct {
	InsertedTuples   uint
	UpdatedTuples    uint
	DeletedTuples    uint
	HOTUpdatedTuples uint
}

type SizeBreakdownRow struct {
	Main      uint
	Vm        uint
	Fsm       uint
	Toast     uint
	Indexes   uint
	DiskFiles uint
}

type ConnectionsRow struct {
	Username          string
	Idle              uint
	IdleInTransaction uint
	Unknown           uint
	QueryActive       uint
	Waiting           uint
}
