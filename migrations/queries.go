package migrations

const (
	createProductTable = "CREATE TABLE IF NOT EXISTS product (ID INT, NAME VARCHAR(255), PRICE FLOAT, CATEGORY VARCHAR(255));"
	dropProductTable   = "DROP TABLE IF EXISTS product"
)
