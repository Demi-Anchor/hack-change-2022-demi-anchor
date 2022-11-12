package repository

const (
	addDonationSQL = `select * from payments.add_donation($1::text, $2::bigint)`
)
