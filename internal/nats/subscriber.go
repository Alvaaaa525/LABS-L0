package nats

import (
	"encoding/json"
	"log"
	"order-service/internal/cache"
	"order-service/internal/database"
	"order-service/internal/models"

	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	conn  stan.Conn
	cache *cache.Cache
	db    *database.DB
}

func NewSubscriber(natsURL, clusterID, clientID string, cache *cache.Cache, db *database.DB) (*Subscriber, error) {
	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		conn:  conn,
		cache: cache,
		db:    db,
	}, nil
}

func (s *Subscriber) Subscribe(subject string) error {
	_, err := s.conn.Subscribe(subject, func(msg *stan.Msg) {
		var order models.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Error unmarshaling order: %v", err)
			return
		}

		// Сохраняем в БД
		if err := s.db.SaveOrder(&order); err != nil {
			log.Printf("Error saving order to DB: %v", err)
			return
		}

		// Сохраняем в кэш
		s.cache.Set(order.OrderUID, &order)
		log.Printf("Order %s received and cached", order.OrderUID)
	}, stan.DurableName("orders-durable"))

	return err
}

func (s *Subscriber) Close() error {
	return s.conn.Close()
}
