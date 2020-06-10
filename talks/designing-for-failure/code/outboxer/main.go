package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/database/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// START OUTBOXER_DEPS OMIT
	db, err := sql.Open("postgres", os.Getenv("DS_DSN"))
	if err != nil {
		log.Fatalf("could not connect to amqp: %s", err)
	}

	conn, err := amqp.Dial(os.Getenv("ES_DSN"))
	if err != nil {
		log.Fatalf("could not connect to amqp: %s", err)
	}

	// we need to create a data store instance first
	ds, err := postgres.WithInstance(ctx, db)
	if err != nil {
		log.Fatalf(("could not setup the data store: %s", err)
	}
	defer ds.Close()
	// END OUTBOXER_DEPS OMIT

	// START OUTBOXER_SETUP OMIT
	o, err := outboxer.New(
		outboxer.WithDataStore(ds),
		outboxer.WithEventStream(amqpOut.NewAMQP(conn)),
		outboxer.WithCheckInterval(1*time.Second),
		outboxer.WithCleanupInterval(5*time.Second),
		outboxer.WithCleanUpBefore(time.Now().AddDate(0, 0, -5)),
	)
	if err != nil {
		log.Fatalf("could not create an outboxer instance: %s", err)
	}
	// END OUTBOXER_SETUP OMIT

	// START OUTBOXER OMIT
	// here we initialize the outboxer checks and cleanup go rotines
	o.Start(ctx)
	defer o.Stop()

	// finally we are ready to send messages
	if err = o.Send(ctx, &outboxer.OutboxMessage{
		Payload: []byte("test payload"),
		Options: map[string]interface{}{
			amqpOut.ExchangeNameOption: "test",
			amqpOut.ExchangeTypeOption: "topic",
			amqpOut.RoutingKeyOption:   "test.send",
		},
	}); err != nil {
		log.Fatalf("could not send message: %s", err)
	}
	// END OUTBOXER OMIT

	// we can also listen for errors and ok messages that were send
	for {
		select {
		case err := <-o.ErrChan():
			log.Printf("could not send message: %s", err)
		case <-o.OkChan():
			log.Printf("message received")
			return
		}
	}
}
