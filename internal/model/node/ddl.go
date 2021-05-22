package node

func(node) SqlDDLStatement() string {
	return `CREATE TABLE node (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"enc_key" BLOB,
		"name" TEXT,
		"program" TEXT		
	);
`
}
