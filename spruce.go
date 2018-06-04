package main

import (
	// core
	"fmt"
	"log"

	// database
	"database/sql"
	_ "github.com/gregbuehler/spruce/database"
	_ "github.com/lib/pq"

	// web/routing
	"net/http"
	"github.com/gorilla/mux"

	// internals
	"github.com/gregbuehler/spruce/accounts"
	"github.com/gregbuehler/spruce/catalogs"
	"github.com/gregbuehler/spruce/customers"
	"github.com/gregbuehler/spruce/funding"
	"github.com/gregbuehler/spruce/orders"
	"github.com/gregbuehler/spruce/status"
)

var (
	Version        = "???"
	Branch         = "???"
	BuildTimestamp = "???"
)

var (
	databaseAdapter          = "postgres"
	databaseConnectionString = "postgres://spruce:spruce@spruce/spruce?sslmode=disable"
)

var (
	SpruceEndpoint = ":8080"
)

func main() {
	fmt.Printf("spruce %s %s %s\n", Version, Branch, BuildTimestamp)

	db, err := sql.Open(databaseAdapter, databaseConnectionString)
	if err != nil {
		panic("Database connection could not be initialized: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("Database connection could not be established: " + err.Error())
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	/*
	 *
	 *	Spruce API Routes
	 *
	 */
	api := router.PathPrefix("/api/v2").Subrouter()

	// Status
	api.HandleFunc("/status", status.StatusHandler).Methods("GET")

	// Accounts
	api.HandleFunc("/accounts", accounts.ListHandler).Methods("GET")
	api.HandleFunc("/accounts/{account_id}", accounts.DetailHandler).Methods("GET")

	// Catalogs
	api.HandleFunc("/catalogs", catalogs.ItemsHandler).Methods("GET")

	// Customers
	api.HandleFunc("/customers", customers.ListHandler).Methods("GET")
	api.HandleFunc("/customers", customers.CreationHandler).Methods("GET")
	api.HandleFunc("/customers/{customer_id}", customers.DetailHandler).Methods("GET")
	api.HandleFunc("/customers/{customer_id}/accounts", customers.AccountListHandler).Methods("GET")
	api.HandleFunc("/customers/{customer_id}/accounts", customers.AccountCreationHandler).Methods("POST")

	// Funding
	api.HandleFunc("/funding", funding.HistoryHandler).Methods("GET")
	api.HandleFunc("/funding/deposits", funding.DepositFundsHandler).Methods("POST")
	api.HandleFunc("/funding/deposits/{deposit_id}", funding.DepositDetailHandler).Methods("GET")
	api.HandleFunc("/funding/sources", funding.SourceListHandler).Methods("GET")
	api.HandleFunc("/funding/sources", funding.SourceCreationHandler).Methods("POST")
	api.HandleFunc("/funding/sources/{source_id}", funding.SourceDetailHandler).Methods("GET")
	api.HandleFunc("/funding/sources/{source_id}", funding.SourceRemovalHandler).Methods("DELETE")

	// Orders
	api.HandleFunc("/orders", orders.ListHandler).Methods("GET")
	api.HandleFunc("/orders", orders.CreationHandler).Methods("POST")
	api.HandleFunc("/orders/{order_id}", orders.DetailHandler).Methods("GET")
	api.HandleFunc("/orders/{order_id}/resend", orders.ResendHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(SpruceEndpoint, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index!")
}
