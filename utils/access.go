package utils

const (
	GET_CATEGORY_LOAN_BY_ID   = "SELECT id,category_loan_name, created_at, updated_at FROM category_loan WHERE id = $1"
	GET_CATEGORY_LOAN_BY_NAME = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan WHERE category_loan_name = $1"
	GET_ALLCATEGORYLOAN       = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan  id = $1"
	INSERT_CATEGORY_LOAN      = "INSERT INTO category_loan ( category_loan_name, created_at, updated_at) VALUES($1, $2, $3) RETURNING id"
	UPDATE_CATEGORY_LOAN      = " UPDATE INTO category_loan (category_loan_name, created_at, updated_at) VALUES($1, $2, $3)"
	DELETE_CATEGORYLOAN       = "DELETE INTO category_loan WHERE id = $1"
)
