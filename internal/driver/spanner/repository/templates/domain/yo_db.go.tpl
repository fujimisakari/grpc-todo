// YODB is the common interface for database operations.
type YODB interface {
	YORODB
}

// YORODB is the common interface for database operations.
type YORODB interface {
	ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)
	Read(ctx context.Context, table string, keys spanner.KeySet, columns []string) *spanner.RowIterator
	ReadUsingIndex(ctx context.Context, table, index string, keys spanner.KeySet, columns []string) (ri *spanner.RowIterator)
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}
