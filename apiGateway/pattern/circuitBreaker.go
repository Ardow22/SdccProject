package pattern

import (
	"github.com/sony/gobreaker"
	"log"
	"time"
)

// circuit breakers globali
var (
	CBMatching       *gobreaker.CircuitBreaker
	CBNotification   *gobreaker.CircuitBreaker
	CBRecommendation *gobreaker.CircuitBreaker
)

// inizializzazione di tutti i circuit breaker per i 3 servizi
func InitCircuitBreakers() {
	CBMatching = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "MatchingService",
		MaxRequests: 1,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		OnStateChange: logStateChange,
	})

	CBNotification = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "NotificationService",
		MaxRequests: 1,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		OnStateChange: logStateChange,
	})

	CBRecommendation = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "RecommendationService",
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
