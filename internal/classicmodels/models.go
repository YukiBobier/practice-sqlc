// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package classicmodels

import (
	"database/sql"
	"time"
)

type Customer struct {
	Customernumber         int32
	Customername           string
	Contactlastname        string
	Contactfirstname       string
	Phone                  string
	Addressline1           string
	Addressline2           sql.NullString
	City                   string
	State                  sql.NullString
	Postalcode             sql.NullString
	Country                string
	Salesrepemployeenumber sql.NullInt32
	Creditlimit            sql.NullString
}

type Employee struct {
	Employeenumber int32
	Lastname       string
	Firstname      string
	Extension      string
	Email          string
	Officecode     string
	Reportsto      sql.NullInt32
	Jobtitle       string
}

type Office struct {
	Officecode   string
	City         string
	Phone        string
	Addressline1 string
	Addressline2 sql.NullString
	State        sql.NullString
	Country      string
	Postalcode   string
	Territory    string
}

type Order struct {
	Ordernumber    int32
	Orderdate      time.Time
	Requireddate   time.Time
	Shippeddate    sql.NullTime
	Status         string
	Comments       sql.NullString
	Customernumber int32
}

type Orderdetail struct {
	Ordernumber     int32
	Productcode     string
	Quantityordered int32
	Priceeach       string
	Orderlinenumber int32
}

type Payment struct {
	Customernumber int32
	Checknumber    string
	Paymentdate    time.Time
	Amount         string
}

type Product struct {
	Productcode        string
	Productname        string
	Productline        string
	Productscale       string
	Productvendor      string
	Productdescription string
	Quantityinstock    int32
	Buyprice           string
	Msrp               string
}

type Productline struct {
	Productline     string
	Textdescription sql.NullString
	Htmldescription sql.NullString
	Image           sql.NullString
}
