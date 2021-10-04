package cb

import (
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/facebookgo/clock"
	"github.com/mercari/go-circuitbreaker"
	"viniti.us/hashout/config/log"
)

type DiscountCB struct {
	CB *circuitbreaker.CircuitBreaker
}

func NewDiscountCB() DiscountCB {
	cb := circuitbreaker.New(
		circuitbreaker.WithClock(clock.New()),
		circuitbreaker.WithFailOnContextCancel(true),
		circuitbreaker.WithFailOnContextDeadline(true),
		circuitbreaker.WithHalfOpenMaxSuccesses(10),
		circuitbreaker.WithOpenTimeoutBackOff(backoff.NewExponentialBackOff()),
		circuitbreaker.WithOpenTimeout(10*time.Second),
		circuitbreaker.WithCounterResetInterval(10*time.Second),
		circuitbreaker.WithTripFunc(circuitbreaker.NewTripFuncFailureRate(10, 0.4)),
		circuitbreaker.WithOnStateChangeHookFn(func(from, to circuitbreaker.State) {
			log.Logger.Infof("[DiscountCB] state changed from %s to %s\n", from, to)
		}),
	)

	return DiscountCB{
		CB: cb,
	}

}
