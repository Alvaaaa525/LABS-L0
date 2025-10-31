package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("orders-cluster", "test-publisher", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Массив тестовых заказов
	orders := []map[string]interface{}{
		// Заказ 1: Электроника из Москвы
		{
			"order_uid":    "ORDER001",
			"track_number": "TRACK001",
			"entry":        "WBIL",
			"delivery": map[string]string{
				"name":    "Иван Иванов",
				"phone":   "+79001234567",
				"city":    "Москва",
				"address": "Тверская ул., д. 10",
				"region":  "Москва",
				"email":   "ivan@example.ru",
				"zip":     "101000",
			},
			"payment": map[string]interface{}{
				"transaction":   "tx001",
				"currency":      "RUB",
				"provider":      "wbpay",
				"amount":        3500000,
				"payment_dt":    time.Now().Unix(),
				"bank":          "sberbank",
				"delivery_cost": 50000,
				"goods_total":   3450000,
				"custom_fee":    0,
				"request_id":    "req001",
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      1001,
					"track_number": "TRACK001",
					"price":        3450000,
					"name":         "Ноутбук ASUS ROG",
					"sale":         10,
					"size":         "15.6",
					"total_price":  3450000,
					"nm_id":        2001,
					"brand":        "ASUS",
					"status":       200,
					"rid":          "rid001",
				},
			},
			"locale":             "ru",
			"customer_id":        "customer001",
			"date_created":       time.Now().Add(-48 * time.Hour),
			"delivery_service":   "cdek",
			"internal_signature": "sig001",
			"shardkey":           "1",
			"sm_id":              1,
			"oof_shard":          "1",
		},

		// Заказ 2: Одежда из Санкт-Петербурга
		{
			"order_uid":    "ORDER002",
			"track_number": "TRACK002",
			"entry":        "WBIL",
			"delivery": map[string]string{
				"name":    "Мария Петрова",
				"phone":   "+79112223344",
				"city":    "Санкт-Петербург",
				"address": "Невский пр., д. 55",
				"region":  "Ленинградская обл.",
				"email":   "maria@example.ru",
				"zip":     "190000",
			},
			"payment": map[string]interface{}{
				"transaction":   "tx002",
				"currency":      "RUB",
				"provider":      "alfabank",
				"amount":        850000,
				"payment_dt":    time.Now().Unix(),
				"bank":          "alfa",
				"delivery_cost": 35000,
				"goods_total":   815000,
				"custom_fee":    0,
				"request_id":    "req002",
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      1002,
					"track_number": "TRACK002",
					"price":        450000,
					"name":         "Куртка зимняя",
					"sale":         15,
					"size":         "L",
					"total_price":  450000,
					"nm_id":        2002,
					"brand":        "Nike",
					"status":       200,
					"rid":          "rid002a",
				},
				{
					"chrt_id":      1003,
					"track_number": "TRACK002",
					"price":        365000,
					"name":         "Кроссовки Air Max",
					"sale":         20,
					"size":         "42",
					"total_price":  365000,
					"nm_id":        2003,
					"brand":        "Nike",
					"status":       200,
					"rid":          "rid002b",
				},
			},
			"locale":             "ru",
			"customer_id":        "customer002",
			"date_created":       time.Now().Add(-24 * time.Hour),
			"delivery_service":   "boxberry",
			"internal_signature": "sig002",
			"shardkey":           "2",
			"sm_id":              2,
			"oof_shard":          "2",
		},

		// Заказ 3: Товары для дома из Казани
		{
			"order_uid":    "ORDER003",
			"track_number": "TRACK003",
			"entry":        "WBIL",
			"delivery": map[string]string{
				"name":    "Алексей Сидоров",
				"phone":   "+79505556677",
				"city":    "Казань",
				"address": "Баумана ул., д. 20",
				"region":  "Татарстан",
				"email":   "alex@example.ru",
				"zip":     "420000",
			},
			"payment": map[string]interface{}{
				"transaction":   "tx003",
				"currency":      "RUB",
				"provider":      "yandexpay",
				"amount":        1250000,
				"payment_dt":    time.Now().Unix(),
				"bank":          "tinkoff",
				"delivery_cost": 45000,
				"goods_total":   1205000,
				"custom_fee":    0,
				"request_id":    "req003",
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      1004,
					"track_number": "TRACK003",
					"price":        750000,
					"name":         "Пылесос Dyson V15",
					"sale":         0,
					"size":         "Standard",
					"total_price":  750000,
					"nm_id":        2004,
					"brand":        "Dyson",
					"status":       200,
					"rid":          "rid003a",
				},
				{
					"chrt_id":      1005,
					"track_number": "TRACK003",
					"price":        455000,
					"name":         "Кофемашина Nespresso",
					"sale":         5,
					"size":         "Compact",
					"total_price":  455000,
					"nm_id":        2005,
					"brand":        "Nespresso",
					"status":       200,
					"rid":          "rid003b",
				},
			},
			"locale":             "ru",
			"customer_id":        "customer003",
			"date_created":       time.Now().Add(-12 * time.Hour),
			"delivery_service":   "pickpoint",
			"internal_signature": "sig003",
			"shardkey":           "3",
			"sm_id":              3,
			"oof_shard":          "1",
		},

		// Заказ 4: Косметика из Екатеринбурга
		{
			"order_uid":    "ORDER004",
			"track_number": "TRACK004",
			"entry":        "WBIL",
			"delivery": map[string]string{
				"name":    "Екатерина Новикова",
				"phone":   "+79607778899",
				"city":    "Екатеринбург",
				"address": "Ленина пр., д. 40",
				"region":  "Свердловская обл.",
				"email":   "kate@example.ru",
				"zip":     "620000",
			},
			"payment": map[string]interface{}{
				"transaction":   "tx004",
				"currency":      "RUB",
				"provider":      "sberpay",
				"amount":        560000,
				"payment_dt":    time.Now().Unix(),
				"bank":          "sberbank",
				"delivery_cost": 30000,
				"goods_total":   530000,
				"custom_fee":    0,
				"request_id":    "req004",
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      1006,
					"track_number": "TRACK004",
					"price":        280000,
					"name":         "Набор косметики Dior",
					"sale":         25,
					"size":         "Set",
					"total_price":  280000,
					"nm_id":        2006,
					"brand":        "Dior",
					"status":       200,
					"rid":          "rid004a",
				},
				{
					"chrt_id":      1007,
					"track_number": "TRACK004",
					"price":        250000,
					"name":         "Парфюм Chanel No5",
					"sale":         0,
					"size":         "100ml",
					"total_price":  250000,
					"nm_id":        2007,
					"brand":        "Chanel",
					"status":       200,
					"rid":          "rid004b",
				},
			},
			"locale":             "ru",
			"customer_id":        "customer004",
			"date_created":       time.Now().Add(-6 * time.Hour),
			"delivery_service":   "ozon",
			"internal_signature": "sig004",
			"shardkey":           "4",
			"sm_id":              4,
			"oof_shard":          "2",
		},

		// Заказ 5: Спортивный инвентарь из Новосибирска
		{
			"order_uid":    "ORDER005",
			"track_number": "TRACK005",
			"entry":        "WBIL",
			"delivery": map[string]string{
				"name":    "Дмитрий Смирнов",
				"phone":   "+79131112233",
				"city":    "Новосибирск",
				"address": "Красный пр., д. 75",
				"region":  "Новосибирская обл.",
				"email":   "dmitry@example.ru",
				"zip":     "630000",
			},
			"payment": map[string]interface{}{
				"transaction":   "tx005",
				"currency":      "RUB",
				"provider":      "raiffeisen",
				"amount":        1850000,
				"payment_dt":    time.Now().Unix(),
				"bank":          "raiffeisen",
				"delivery_cost": 60000,
				"goods_total":   1790000,
				"custom_fee":    0,
				"request_id":    "req005",
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      1008,
					"track_number": "TRACK005",
					"price":        1200000,
					"name":         "Велосипед Trek Mountain",
					"sale":         5,
					"size":         "L",
					"total_price":  1200000,
					"nm_id":        2008,
					"brand":        "Trek",
					"status":       200,
					"rid":          "rid005a",
				},
				{
					"chrt_id":      1009,
					"track_number": "TRACK005",
					"price":        590000,
					"name":         "Шлем спортивный Pro",
					"sale":         0,
					"size":         "M",
					"total_price":  590000,
					"nm_id":        2009,
					"brand":        "ProSport",
					"status":       200,
					"rid":          "rid005b",
				},
			},
			"locale":             "ru",
			"customer_id":        "customer005",
			"date_created":       time.Now().Add(-3 * time.Hour),
			"delivery_service":   "dpd",
			"internal_signature": "sig005",
			"shardkey":           "5",
			"sm_id":              5,
			"oof_shard":          "1",
		},

		// Заказ 6: Книги из Владивостока
		{
			"order_uid":    "ORDER006",
			"track_number": "TRACK006",
			"entry":        "WBIL",
			"delivery": map[string]string{
				"name":    "Ольга Козлова",
				"phone":   "+79234445566",
				"city":    "Владивосток",
				"address": "Океанский пр., д. 12",
				"region":  "Приморский край",
				"email":   "olga@example.ru",
				"zip":     "690000",
			},
			"payment": map[string]interface{}{
				"transaction":   "tx006",
				"currency":      "RUB",
				"provider":      "wbpay",
				"amount":        320000,
				"payment_dt":    time.Now().Unix(),
				"bank":          "vtb",
				"delivery_cost": 25000,
				"goods_total":   295000,
				"custom_fee":    0,
				"request_id":    "req006",
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      1010,
					"track_number": "TRACK006",
					"price":        150000,
					"name":         "Энциклопедия программирования",
					"sale":         10,
					"size":         "Hardcover",
					"total_price":  150000,
					"nm_id":        2010,
					"brand":        "O'Reilly",
					"status":       200,
					"rid":          "rid006a",
				},
				{
					"chrt_id":      1011,
					"track_number": "TRACK006",
					"price":        145000,
					"name":         "Учебник по Go",
					"sale":         15,
					"size":         "Paperback",
					"total_price":  145000,
					"nm_id":        2011,
					"brand":        "Manning",
					"status":       200,
					"rid":          "rid006b",
				},
			},
			"locale":             "ru",
			"customer_id":        "customer006",
			"date_created":       time.Now().Add(-1 * time.Hour),
			"delivery_service":   "pochta",
			"internal_signature": "sig006",
			"shardkey":           "6",
			"sm_id":              6,
			"oof_shard":          "2",
		},
	}

	// Публикуем все заказы
	for i, order := range orders {
		data, _ := json.Marshal(order)
		err = sc.Publish("orders", data)
		if err != nil {
			log.Printf("❌ Ошибка при публикации заказа %d: %v", i+1, err)
			continue
		}
		log.Printf("✓ Заказ %d опубликован: %s", i+1, order["order_uid"])
		time.Sleep(100 * time.Millisecond) // Небольшая задержка между заказами
	}

	log.Println("✅ Все заказы успешно опубликованы в NATS!")
}
