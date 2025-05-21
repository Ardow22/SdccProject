package pattern

import (
	"github.com/sony/gobreaker"
	"log"
	"time"
)

// dichiarazione del circuit breaker
var (
	CBRabbitMQ *gobreaker.CircuitBreaker
)

// inizializzazione del circuit breaker
func InitCircuitBreakers() {
	CBRabbitMQ = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "RabbitMQ",
		MaxRequests: 1,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		OnStateChange: logStateChange,
	})
}

// log del cambio di stato del Circuit Breaker
func logStateChange(name string, from, to gobreaker.State) {
	log.Printf("CircuitBreaker [%s] passa dallo stato %s allo stato %s", name, from.String(), to.String())
}
