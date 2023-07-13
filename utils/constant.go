package utils

const (
	ADD_CUSTOMER           = "INSERT INTO mst_customer(id, full_name, address, nik, phone_number, user_id, created_at) VALUES($1, $2, $3, $4, $5, $6, $7)"
	GET_CUSTOMER_BY_ID     = "SELECT id, full_name, address, nik, phone_number, user_id, created_at, updated_at FROM mst_customer WHERE id = $1"
	GET_ALL_CUSTOMER       = "SELECT * FROM mst_customer"
	UPDATE_CUSTOMER        = "UPDATE mst_customer SET full_name=$1, address=$2, nik=$3, phone_number=$4, user_id=$5, updated_at=$6 WHERE id=$7"
	DELETE_CUSTOMER        = "DELETE FROM mst_customer WHERE id=$1"
	GET_CUSTOMER_BY_NIK    = "SELECT id, full_name, address, nik, phone_number, user_id, created_at, updated_at FROM mst_customer WHERE nik = $1"
	GET_CUSTOMER_BY_NUMBER = "SELECT id, full_name, address, nik, phone_number, user_id, created_at, updated_at FROM mst_customer WHERE phone_number = $1"
)
