package main

import (
	"log"
	"net/http"
	"order-service/internal/cache"
	"order-service/internal/database"
	"order-service/internal/handlers"
	natsservice "order-service/internal/nats"

	"github.com/gorilla/mux"
)

func main() {
	// Подключение к PostgreSQL
	log.Println("Connecting to PostgreSQL...")
	db, err := database.NewDB("localhost", "5432", "orderuser", "orderpass", "ordersdb")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	log.Println("✓ Connected to PostgreSQL")

	// Создание кэша
	cache := cache.NewCache()

	// Восстановление кэша из БД
	log.Println("Restoring cache from database...")
	orders, err := db.GetAllOrders()
	if err != nil {
		log.Printf("Warning: Failed to restore cache: %v", err)
	} else {
		cache.LoadFromMap(orders)
		log.Printf("✓ Cache restored: %d orders loaded", len(orders))
	}

	// Подключение к NATS Streaming
	log.Println("Connecting to NATS Streaming...")
	subscriber, err := natsservice.NewSubscriber(
		"nats://localhost:4222",
		"orders-cluster",
		"order-service-1",
		cache,
		db,
	)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer subscriber.Close()
	log.Println("✓ Connected to NATS Streaming")

	// Подписка на канал
	if err := subscriber.Subscribe("orders"); err != nil {
		log.Fatal("Failed to subscribe:", err)
	}
	log.Println("✓ Subscribed to NATS channel 'orders'")

	// Настройка HTTP сервера
	handler := handlers.NewHandler(cache)
	router := mux.NewRouter()

	router.HandleFunc("/", handler.GetOrderPage).Methods("GET")
	router.HandleFunc("/api/order/{id}", handler.GetOrder).Methods("GET")
	router.HandleFunc("/api/orders", handler.GetAllOrderIDs).Methods("GET")

	log.Println("\n========================================")
	log.Println("HTTP server starting on http://localhost:8080")
	log.Println("========================================\n")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}
