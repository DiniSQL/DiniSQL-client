package Type

type OperationType = int

const (
	CreateDatabase OperationType = iota
	UseDatabase
	CreateTable
	CreateIndex
	DropTable
	DropIndex
	DropDatabase
	Insert
	Update
	Delete
	Select
	ExecFile
)

type PacketType = int

const (
	KeepAlive      PacketType = iota // RegionServer send to Master
	Ask                              // Client send to Master to know which RegionServer it should visit
	Answer                           // Master answer the Ask packet from Client, tell which RegionServer
	SQLOperation                     // Client send to RegionServer to execute a SQL operation
	Result                           // RegionServer send to Client, the result of the SQL operation
	UploadRegion                     // Master send to RegionServer, tell the RegionServer to upload a region to etcd
	DownloadRegion                   //Master send to RegionServer, tell the RegionServer to download a region from etcd
)

// type Column struct {
// 	Name      string
// 	Type      ColumnType
// 	Unique    bool
// 	NotNull   bool
// 	ColumnPos int //the created position when table is created, this value is fixed
// }

// type ColumnType struct {
// 	TypeTag ScalarColumnTypeTag
// 	Length  int
// 	IsArray bool
// }
// type ScalarColumnTypeTag = int

// const (
// 	Bool ScalarColumnTypeTag = iota
// 	Int64
// 	Float64
// 	String
// 	Bytes
// 	Date
// 	Timestamp
// )
// type CreateTableStatement struct {
// 	TableName   string
// 	ColumnsMap  map[string]Column
// 	PrimaryKeys []Key
// 	Cluster     Cluster
// }
// // Cluster is a Spanner table cluster.
// type Cluster struct {
// 	TableName string
// 	OnDelete  OnDelete
// }
// // Key is a table key.
// type Key struct {
// 	Name     string
// 	KeyOrder KeyOrder
// }
// //KeyOrder order for key
// type KeyOrder = int

// const (
// 	Asc KeyOrder = iota
// 	Desc
// )
// //OnDelete is used for on delete behave
// type OnDelete = int

// const (
// 	NoAction OnDelete = iota
// 	Cascade
// )
