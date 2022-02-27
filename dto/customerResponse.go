package dto

//to ommit any response from the response provided by the API we do not need to change the customer in the domain level we can just change the dto and ommit that specific field
type CustomerResponse struct {
	Id      int64  `db:"id"`
	Name    string `db:"name"`
	City    string `db:"city"`
	Zipcode string `db:"zipcode"`
	Status  string `db:"status"`
}
