package backoff

import (
	"errors"
	"math"
	"math/rand/v2"
	"time"
)

type Config struct {
	MaxAttempts    int
	MinDelay       time.Duration
	MaxDelay       time.Duration
	JitterFraction float64
	Factor         float64
}

type Backoff struct {
	maxAttempts    int
	maxDelay       time.Duration
	minDelay       float64
	floatMaxDelay  float64
	jitterFraction float64
	factor         float64
	attempts       int
}

func NewBackoff(cfg Config) *Backoff {
	if cfg.MaxAttempts == 0 {
		cfg.MaxAttempts = 10
	}
	if cfg.MinDelay == 0 {
		cfg.MinDelay = 1 * time.Second
	}
	if cfg.MaxDelay == 0 {
		cfg.MaxDelay = 5 * time.Minute
	}
	if cfg.JitterFraction == 0 {
		cfg.JitterFraction = 0.1
	}
	if cfg.Factor == 0 {
		cfg.Factor = 2.0
	}

	return &Backoff{
		maxAttempts:    cfg.MaxAttempts,
		minDelay:       float64(cfg.MinDelay),
		maxDelay:       cfg.MaxDelay,
		floatMaxDelay:  float64(cfg.MaxDelay),
		jitterFraction: cfg.JitterFraction,
		factor:         cfg.Factor,
	}
}

func (c *Backoff) NextDelay() (time.Duration, error) {
	if c.attempts > c.maxAttempts {
		return 0, errors.New("max attempts reached")
	}

	pow := math.Pow(c.factor, float64(c.attempts))

	delayFloat := pow * c.minDelay
	if delayFloat > c.floatMaxDelay {
		delayFloat = c.floatMaxDelay
	}

	delay := time.Duration(delayFloat + rand.NormFloat64()*c.jitterFraction*delayFloat)
	c.attempts++

	return delay, nil
}
