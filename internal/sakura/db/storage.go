package db

type DataType int

const (
	TypeString  DataType = iota
	TypeList    
	TypeHashMap 
	TypeSet    
)

type DataNode struct {
	Type      DataType
	ExpiresAt int
	Value     interface{}
}

func NewDataNode(t DataType, exp int, val interface{}) *DataNode {
	return &DataNode{t, exp, val}
}
