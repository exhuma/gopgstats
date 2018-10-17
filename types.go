package gopgstats

type DiskSizeRow struct {
	DatabaseName string
	Size         int
}

type LocksRow struct {
	DatabaseName string
	Mode         string
	Type         string
	Granted      bool
	Count        int
}
