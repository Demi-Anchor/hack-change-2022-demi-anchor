package repository

const (
	addPaymentSQL = `select * from payments.add_payment($1::text, $2::bigint)`
)
