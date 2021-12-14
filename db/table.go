package db

type Table struct {
	*DBCon
	table string
}

func (t *Table) TableName() string {
	return t.table
}

func (t *Table) InsertMap(data Data) int64 {
	return t.InsertByMap(t.table, data)
}
