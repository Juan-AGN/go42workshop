package reservations

// Bloque 1: seis consultas de práctica sobre las tablas products y
// reservations. Cada una es un texto SQL con nombre, igual que las de
// store.go, que vienen escritas y te sirven de ejemplo. El comentario dice
// qué tiene que devolver o hacer cada una: cambia el TODO por la sentencia.
//
// Los ? son los valores que cambian de una llamada a otra: los pone el test,
// en orden. Pruébalas antes a mano con cmd/sql, con un valor en vez del ?:
//
//	go run ./cmd/sql "SELECT name, stock FROM products"
const (
	// TODO: todos los productos, con sus tres columnas (id, name, stock),
	// ordenados por nombre.
	sqlPracticeProductsByName = `SELECT * FROM products ORDER BY name`

	// TODO: el nombre de los productos con menos stock que el que se pasa,
	// ordenados por nombre.
	sqlPracticeLowStock = `SELECT name FROM products WHERE stock < ?`

	// TODO: el id y la quantity de las reservas de un producto, de la más
	// antigua a la más nueva.
	sqlPracticeReservationsOf = `SELECT id, quantity FROM reservations WHERE product_id = ? ORDER BY id`

	// TODO: cuántas reservas hay en un estado.
	sqlPracticeCountByStatus = `SELECT COUNT(*) FROM reservations WHERE status = ?`

	// TODO: cambia el nombre de un producto. Primero llega el nombre nuevo y
	// después el id.
	sqlPracticeRename = `UPDATE products SET name = ? WHERE id = ?`

	// TODO: borra las reservas que están en un estado.
	sqlPracticeDeleteByStatus = `DELETE FROM reservations WHERE status = ?`

	sqlCountById = `SELECT COUNT(*) FROM reservations WHERE product_id = ?`

	sqlCountConfirmed = `SELECT COUNT(*) FROM reservations AS tcount WHERE product_id = ? AND status = ?`
)
