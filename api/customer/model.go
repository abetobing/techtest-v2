package customer

type CustomerDetail struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Address string `json:"address"`
	ZipCode string `json:"zipcode"`
}

type CreateResponse struct {
	CustomerID int `json:"customer_id"`
}

type ListResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
