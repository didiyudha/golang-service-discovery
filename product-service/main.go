package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

type product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductConfiguration struct {
	Categories []string `json:"categories"`
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "product-service"
	registration.Name = "product-service"
	address := hostname()
	registration.Address = address
	port, err := strconv.Atoi(port()[1:len(port())])
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = port
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, port)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}

func Configuration(w http.ResponseWriter, r *http.Request) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	kvpair, _, err := consul.KV().Get("product-configuration", nil)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	if kvpair.Value == nil {
		fmt.Fprintf(w, "Configuration empty")
		return
	}
	val := string(kvpair.Value)
	fmt.Fprintf(w, "%s", val)

}

func Products(w http.ResponseWriter, r *http.Request) {
	products := []product{
		{
			ID:    1,
			Name:  "Sepeda",
			Price: 2000000.00,
		},
		{
			ID:    2,
			Name:  "Mecin",
			Price: 500.00,
		},
		{
			ID:    3,
			Name:  "Kacamata",
			Price: 1500000.00,
		},
		{
			ID:    4,
			Name:  "Kopi Gayo",
			Price: 50000.00,
		},
		{
			ID:    5,
			Name:  "Macbook Pro",
			Price: 20000000.00,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&products)
}

func main() {
	registerServiceWithConsul()
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/products", Products)
	http.HandleFunc("/product-configuration", Configuration)
	fmt.Printf("product service is up on port: %s", port())
	http.ListenAndServe(port(), nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `product service is good`)
}

func port() string {
	p := os.Getenv("PRODUCT_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8100"
	}
	return fmt.Sprintf(":%s", p)
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
