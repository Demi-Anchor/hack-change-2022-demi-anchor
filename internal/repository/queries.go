package repository

const (
	addDonationSQL       = `select * from payments.add_donation($1::integer, $2::text, $3::bigint, $4::text, $5::timestamp)`
	getDailyDonationsSQL = `select * from payments.get_daily_donations($1::date, $2::date)`
)
