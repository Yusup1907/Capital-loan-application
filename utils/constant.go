package utils

const (
	GET_ALL_PRODUCT = `SELECT 
							mst_product.id, 
							mst_product.product_name, 
							mst_product.description, 
							mst_product.price, 
							mst_product.stok, 
							mst_product.category_product_id,
							category_product.category_product_name,
							mst_product.status, 
							mst_product.created_at, 
							mst_product.updated_at
						FROM 
							mst_product
						INNER JOIN 
							category_product ON mst_product.category_product_id = category_product.id
						ORDER BY 
							mst_product.id ASC`
	GET_PRODUCT_BY_ID = `SELECT 
							mst_product.id,
							mst_product.product_name, 
							mst_product.description, 
							mst_product.price, 
							mst_product.stok, 
							mst_product.category_product_id,
							category_product.category_product_name,
							mst_product.status, 
							mst_product.created_at, 
							mst_product.updated_at
						FROM 
							mst_product
						INNER JOIN 
							category_product ON mst_product.category_product_id = category_product.id
						WHERE
							mst_product.id = $1
						ORDER BY 
							mst_product.id ASC`
	GET_PRODUCT_BY_NAME            = "SELECT id, product_name, description, price, stok, category_product_id, status, created_at, updated_at FROM mst_product WHERE product_name = $1"
	UPDATE_PRODUCT                 = "UPDATE mst_product SET product_name = $1, description = $2, price = $3, stok = $4,  category_product_id = $5, status = $6, updated_at = $7 WHERE id = $8"
	DELETE_PRODUCT                 = "DELETE FROM mst_product WHERE id = $1;"
	ADD_PRODUCT                    = "INSERT INTO mst_product (product_name, description, price, stok, category_product_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	GET_CATEGORY_LOAN_BY_ID        = "SELECT id,category_loan_name, created_at, updated_at FROM category_loan WHERE id = $1"
	GET_CATEGORY_LOAN_BY_NAME      = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan WHERE category_loan_name = $1"
	GET_ALLCATEGORYLOAN            = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan ORDER BY id ASC"
	INSERT_CATEGORY_LOAN           = "INSERT INTO category_loan ( category_loan_name, created_at, updated_at) VALUES($1, $2, $3) RETURNING id"
	UPDATE_CATEGORY_LOAN           = "UPDATE category_loan SET category_loan_name = $1, updated_at = $2, created_at = $3 WHERE id = $4"
	DELETE_CATEGORYLOAN            = "DELETE FROM category_loan WHERE id = $1"
	GET_CATEGORY_UPDATE_ID         = "SELECT id FROM category_loan WHERE id = $1"
	INSERT_CATEGORY_PRODUCT        = "INSERT INTO category_product(category_product_name, created_at, updated_at) VALUES ($1 ,$2, DEFAULT) RETURNING id"
	DELETE_CATEGORYPRODUCT         = "DELETE FROM category_product WHERE id = $1"
	UPDATE_CATEGORY_PRODUCT        = "UPDATE category_product SET category_product_name = $1, updated_at = $2 WHERE id = $3"
	GET_CATEGORY_PRODUCT_UPDATE_ID = "SELECT id FROM category_product WHERE id = $1"
	GET_ALLCATEGORYPRODUCT         = "SELECT  id, category_product_name, created_at, updated_at FROM category_product ORDER BY id"
	GET_CATEGORY_PRODUCT_BY_ID     = "SELECT id, category_product_name, created_at, updated_at FROM category_product WHERE id = $1"
	GET_CATEGORY_PRODUCT_BY_NAME   = "SELECT id, category_product_name, created_at, updated_at FROM category_product WHERE category_product_name = $1"
	ADD_CUSTOMER                   = "INSERT INTO mst_customer(id, full_name, address, nik, phone_number, user_id, created_at) VALUES($1, $2, $3, $4, $5, $6, $7)"
	GET_CUSTOMER_BY_ID             = "SELECT id, full_name, address, nik, phone_number, user_id, created_at, updated_at FROM mst_customer WHERE id = $1"
	GET_ALL_CUSTOMER               = "SELECT * FROM mst_customer"
	UPDATE_CUSTOMER                = "UPDATE mst_customer SET full_name=$1, address=$2, nik=$3, phone_number=$4, user_id=$5, updated_at=$6 WHERE id=$7"
	DELETE_CUSTOMER                = "DELETE FROM mst_customer WHERE id=$1"
	GET_CUSTOMER_BY_NIK            = "SELECT id, full_name, address, nik, phone_number, user_id, created_at, updated_at FROM mst_customer WHERE nik = $1"
	GET_CUSTOMER_BY_NUMBER         = "SELECT id, full_name, address, nik, phone_number, user_id, created_at, updated_at FROM mst_customer WHERE phone_number = $1"
	CREATE_APLICATION_LOAN_REPO    = `INSERT INTO trx_loan (customer_id, loan_date, due_date, category_loan_id, amount, description, status, repayment_status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	GET_CUSTOMER_LOAN_BY_ID = "SELECT id, nik, nokk, emergencyname, emergencycontact, last_salary FROM mst_customer WHERE id = $1"
	GET_LOAN_APLICATION     = `SELECT la.id, la.customer_id, la.loan_date, la.due_date, la.category_loan_id, la.amount, la.description, la.status, la.repayment_status, la.created_at, la.updated_at,
	mc.full_name, mc.address, mc.nik, mc.phone_number, mc.nokk, mc.emergencyname, mc.emergencycontact, mc.last_salary FROM trx_loan la INNER JOIN mst_customer mc ON la.customer_id = mc.id ORDER BY la.id ASC OFFSET $1 LIMIT $2`
	GET_LOAN_APLICATION_BY_ID = `SELECT 
	la.id, 
	la.customer_id, 
	la.loan_date, 
	la.due_date, 
	la.category_loan_id, 
	la.amount, 
	la.description, 
	la.status,
	la.repayment_status, 
	la.created_at, 
	la.updated_at,
	mc.full_name, 
	mc.address, 
	mc.nik, 
	mc.phone_number, 
	mc.nokk, 
	mc.emergencyname, 
	mc.emergencycontact, 
	mc.last_salary
FROM 
	trx_loan la
INNER JOIN 
	mst_customer mc ON la.customer_id = mc.id
WHERE
	la.id = $1
ORDER BY la.id`

	GET_LOAN_APLCATION_REPAYMENT_STATUS = `SELECT 
					la.id, 
					la.customer_id, 
					la.loan_date, 
					la.due_date, 
					la.category_loan_id, 
					la.amount, 
					la.description, 
					la.status, 
					la.repayment_status, 
					la.created_at, 
					la.updated_at,
			   		mc.full_name, 
					mc.address, 
					mc.nik, 
					mc.phone_number, 
					mc.nokk, 
					mc.emergencyname, 
					mc.emergencycontact, 
					mc.last_salary
				FROM 
					trx_loan la
				INNER JOIN mst_customer mc ON la.customer_id = mc.id
				WHERE la.repayment_status = $3
				ORDER BY la.id ASC
				OFFSET $1 LIMIT $2`
	LOAN_REPAYMENT = "UPDATE trx_loan SET payment_date = $1, payment = $2, repayment_status = $3::loan_status, updated_at = $4 WHERE id = $5"
)
