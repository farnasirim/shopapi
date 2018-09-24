package main

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/gorilla/mux"

	"github.com/farnasirim/shopapi/api/graphql"
	"github.com/farnasirim/shopapi/data/mongodb"
	shopapihttp "github.com/farnasirim/shopapi/http"
)

var (
	serveAddress            string
	initDB                  bool
	mongodbConnectionString string
	mongodbDatabaseName     string
)

// represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the api on the given address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if serveAddress == "" {
			log.Fatalln("you need to supply the serve address")
		}
		if mongodbConnectionString == "" {
			log.Fatalln("you need to supply mongodb address")
		}

		log.Println("will serve on:", serveAddress)
		// This should be removed. We shouldn't write important stuff
		// right onto the terminal. Also should make it possible to
		// pass in the sensitive information using env variables.
		log.Println("will connect to mongodb at", mongodbConnectionString)
		log.Println("will initialize db?", initDB)
		log.Println("mongodb database name:", mongodbDatabaseName)

		mongodbService := mongodb.NewMongodbService(mongodbConnectionString, mongodbDatabaseName)

		if initDB && len(mongodbService.Shops()) == 0 {
			apple := mongodbService.NewShop("apple")
			iphoneX := mongodbService.CreateProductInShop(apple.ID(), "iphone X", 999, 99)
			ipad := mongodbService.CreateProductInShop(apple.ID(), "ipad", 665, 50)

			samsung := mongodbService.NewShop("samsung")
			galaxyA5 := mongodbService.CreateProductInShop(samsung.ID(), "galaxy A5", 700, 20)
			_ = mongodbService.CreateProductInShop(samsung.ID(), "note", 600, 30)

			orderInApple := mongodbService.CreateOrderInShop(apple.ID())
			mongodbService.AddProductToOrder(orderInApple.ID(), iphoneX.ID(), 2)
			mongodbService.AddProductToOrder(orderInApple.ID(), ipad.ID(), 2)

			orderInSamsung := mongodbService.CreateOrderInShop(samsung.ID())
			mongodbService.AddProductToOrder(orderInSamsung.ID(), galaxyA5.ID(), 1)
			mongodbService.AddProductToOrder(orderInSamsung.ID(), galaxyA5.ID(), 1)
		}

		graphqlService := graphql.NewGrpahqlService(mongodbService)

		router := mux.NewRouter()

		router.Handle("/api", graphqlService.GraphqlHTTPHandler)
		router.Handle("/debug", shopapihttp.NewDebugClient())

		http.ListenAndServe(serveAddress, router)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVar(&serveAddress, "address", "", "ip:port to listen on")
	serveCmd.Flags().BoolVar(&initDB, "initdb", false, "pass to create initial data in the database")
	serveCmd.Flags().StringVar(&mongodbConnectionString, "mongodb-uri", "mongodb://localhost:27017", "mongodb connection string: mongodb://user:pass@ip:port")
	serveCmd.Flags().StringVar(&mongodbDatabaseName, "dbname", "shopapidb", "mongodb database name to be used")
}
