package utils

const (
	GET_CATEGORY_LOAN_BY_ID   = "SELECT id,category_loan_name, created_at, updated_at FROM category_loan WHERE id = $1"
	GET_CATEGORY_LOAN_BY_NAME = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan WHERE category_loan_name = $1"
	GET_ALLCATEGORYLOAN       = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan ORDER BY id ASC"
	INSERT_CATEGORY_LOAN      = "INSERT INTO category_loan ( category_loan_name, created_at, updated_at) VALUES($1, $2, $3) RETURNING id"
	UPDATE_CATEGORY_LOAN      = "UPDATE category_loan SET category_loan_name = $1, updated_at = $2, created_at = $3 WHERE id = $4"
	DELETE_CATEGORYLOAN       = "DELETE FROM category_loan WHERE id = $1"
	GET_CATEGORY_UPDATE_ID    = "SELECT id FROM category_loan WHERE id = $1"
)
