package utils

const (
	INSERT_USER       = "INSERT INTO mst_user(id, user_name, email, password, roles_name, is_active, phone_number, created_at, updated_at ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	UPDATE_USER       = "UPDATE mst_user SET user_name = $1,  email = $2, password = $3, roles_name = $4, is_active = $5, phone_number = $6 ,created_at = $7, updated_at = $8 WHERE id = $9"
	DELETE_USER       = "DELETE FROM mst_user WHERE id = $1"
	GET_ALL_USER      = "SELECT id,user_name, email, password, roles_name, is_active, phone_number ,created_at, updated_at FROM mst_user ORDER BY id ASC"
	GET_USER_BY_NAME  = "SELECT id,user_name, email, password, roles_name, is_active, phone_number ,created_at, updated_at FROM mst_user WHERE user_name = $1"
	GET_USER_BY_ID    = "SELECT id, user_name, email, password, roles_name, is_active, phone_number ,created_at, updated_at FROM mst_user WHERE id = $1"
	GET_USER_BY_EMAIL = "SELECT id,user_name, email, password, roles_name, is_active, phone_number,created_at, updated_at  FROM mst_user WHERE email = $1"
)
