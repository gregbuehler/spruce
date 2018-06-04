package customers

import (
	"fmt"
	"net/http"
)

type CustomerAccount struct {
	AccountIdentifier string `json:"accountIdentifier, required" description:""`
	AccountNumber     string `json:"accountNumber, required" description:""`
	CreatedAt         string `json:"createdAt, required" description:""`
	DisplayName       string `json:"displayName, required" description:""`
	AccountStatus     string `json:"accountStatus, required" description:""`
}

type CustomerDetail struct {
	Accounts       []CustomerAccount `json:"accounts, required" description:""`
	CreatedAt      string            `json:"createdAt, required" description:""`
	DisplayName    string            `json:"displayName, required" description:""`
	CustomerStatus string            `json:"CustomerStatus, required" description:""`
}

type CustomerList struct {
	Customers []CustomerDetail
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}

func CreationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}

func AccountListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}

func AccountCreationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: complete this handler\n\n %s", r.RequestURI)
}
