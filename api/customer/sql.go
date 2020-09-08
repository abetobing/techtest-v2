package customer

import (
	"database/sql"
	"strconv"
)

func CreateCustomer(db *sql.DB, req CustomerDetail) (id int, err error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	var cid int
	err = tx.QueryRow("INSERT INTO customers (name, phone) VALUES ($1, $2) RETURNING id", req.Name, req.Phone).Scan(&cid)
	if err != nil {
		return -1, err
	}

	for _, addr := range req.Addresses {
		_, err = tx.Exec("INSERT INTO customer_addresses (customer_id, address, zipcode) VALUES ($1, $2, $3)", cid, addr.Address, addr.ZipCode)
		if err != nil {
			return -1, err
		}
	}

	err = tx.Commit()
	return cid, err
}

func GetCustomer(db *sql.DB, id int) (cust CustomerDetail, err error) {
	cust.ID = id
	err = db.QueryRow("SELECT name, phone FROM customers WHERE id = $1", strconv.Itoa(id)).Scan(&cust.Name, &cust.Phone)
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT address, zipcode FROM customer_addresses WHERE customer_id = $1", strconv.Itoa(id))
	if err != nil {
		return cust, err
	}

	for rows.Next() {
		var addr Address
		rows.Scan(&addr.Address, &addr.ZipCode)
		cust.Addresses = append(cust.Addresses, addr)
	}

	return
}

func ListCustomer(db *sql.DB) (cs []ListResponse, err error) {
	rows, err := db.Query("SELECT id, name FROM customers")
	if err != nil {
		return cs, err
	}

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		cs = append(cs, ListResponse{id, name})
	}

	return
}

func DeleteCustomer(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM customers WHERE id = $1", strconv.Itoa(id))
	if err != nil {
		return err
	}
	return nil
}
