package main

import (
	"crypto/tls"
	"fmt"
	mdw "goapi/internal/api/middlewares"
	"goapi/internal/api/router"
	"goapi/internal/repositories/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		return
	}

	_, err = database.ConnectDb()

	if err != nil {
		log.Fatal("Error connecting to the database", err)
		return
	}

	port := os.Getenv("API_PORT")

	cert := "../../cert.pem"
	key := "../../key.pem"

	mux := router.Router()

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	/*
		rl := mdw.NewRateLimiter(5, time.Minute)

		hppOptions := mdw.HPPOptions{
			CheckQuery:                  true,
			CheckBody:                   true,
			CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
			Whitelist:                   []string{"allowedParam"},
		}
	*/
	//secureMux := mdw.Hpp(hppOptions)(rl.Middleware(mdw.Compression(mdw.ResponseTime(mdw.SecurityHeaders(mdw.Cors(mux))))))
	//secureMux := mdw.Cors(rl.Middleware(mdw.ResponseTime(mdw.SecurityHeaders(mdw.Compression(mdw.Hpp(hppOptions)(mux))))))
	//secureMux := applyMiddlewares(mux, mdw.Hpp(hppOptions), mdw.Compression, mdw.SecurityHeaders, mdw.ResponseTime, rl.Middleware, mdw.Cors)

	secureMux := mdw.SecurityHeaders(mux)
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server running on port: ", port)
	err = server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatal("Error starting the server", err)
	}

}
