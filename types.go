package gopgstats

type DatabasesRow struct {
	Name string
}

// --- Global Statistics (available from any DB) -----------------------------

type DiskSizesRow struct {
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

type GlobalSizesRow struct {
	DatabaseName string
	Size         uint
}

type QueryAgesRow struct {
	DatabaseName   string
	QueryAge       float64
	TransactionAge float64
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

type DiskIOsRow struct {
	DatabaseName         string
	HeapBlocksRead       uint
	HeapBlocksHit        uint
	IndexBlocksRead      uint
	IndexBlocksHit       uint
	ToastBlocksRead      uint
	ToastBlocksHit       uint
	ToastIndexBlocksRead uint
	ToastIndexBlocksHit  uint
}

type IndexIOsRow struct {
	DatabaseName    string
	IndexBlocksRead uint
	IndexBlocksHit  uint
}

type SequencesIOsRow struct {
	DatabaseName string
	BlocksRead   uint
	BlocksHit    uint
}

type ScanTypesRow struct {
	DatabaseName    string
	IndexScans      uint
	SequentialScans uint
}

type RowAccessesRow struct {
	DatabaseName     string
	InsertedTuples   uint
	UpdatedTuples    uint
	DeletedTuples    uint
	HOTUpdatedTuples uint
}

type SizeBreakdownRow struct {
	DatabaseName string
	Main         uint
	Vm           uint
	Fsm          uint
	Toast        uint
	Indexes      uint
	DiskFiles    uint
}

type ConnectionsRow struct {
	Username          string
	Idle              uint
	IdleInTransaction uint
	Unknown           uint
	QueryActive       uint
	Waiting           uint
}
