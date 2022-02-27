package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rtpa25/banking/domain"
	"github.com/rtpa25/banking/logger"
	"github.com/rtpa25/banking/service"
	"github.com/spf13/viper"
)

func Start(serverUrl string) {
	router := mux.NewRouter()

	//wiring
	dbClient := getDbClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDB(dbClient)
	ch := CustomerHandler{service: service.NewCustomerService(&customerRepositoryDB)}
	ah := AccountHandler{service: service.NewAccountService(&accountRepositoryDB)}
	th := TransactionHandler{service: service.NewTransactionService(&transactionRepositoryDB)}

	//definig routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/transaction", th.NewTransaction).Methods(http.MethodPost)

	//starting server
	err := http.ListenAndServe(serverUrl, router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getDbClient() *sqlx.DB {
	connStr := viper.Get("DB_URL")
	db, err := sqlx.Open("postgres", connStr.(string))
	if err != nil {
		logger.Error(err.Error())
	}
	return db
}
