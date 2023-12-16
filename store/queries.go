package store

const (
	createQuery  = `INSERT INTO product (id,name,price,category)values(?,?,?,?)`
	updateQuery  = `UPDATE product SET name=?, price=?, category=? WHERE id=?`
	getByIDQuery = `SELECT id,name,price,category FROM product WHERE id=?`
	deleteQuery  = `Delete from product where id =?`
)
