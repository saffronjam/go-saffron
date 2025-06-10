package core

import "time"

var GlobalClock *Clock

func SetGlobalClock(clock *Clock) {
	GlobalClock = clock
}

type Clock struct {
	lastTick time.Time
	dt       time.Duration
	start    time.Time
}

func NewClock() *Clock {
	return &Clock{
		dt:       0,
		lastTick: time.Now(),
		start:    time.Now(),
	}
}

func (c *Clock) Tick() {
	c.dt = time.Since(c.lastTick)
	c.lastTick = time.Now()
}

// Delta returns a float32 representing the time elapsed since the last tick in seconds.
// It uses float32 since the SFML library uses that, and it's annoying to have to convert
// between float32 and float64 all the time. If you need higher precision, use DeltaDuration instead.
func (c *Clock) Delta() float32 {
	return float32(c.dt.Seconds())
}

func (c *Clock) DeltaDuration() time.Duration {
	return c.dt
}

func (c *Clock) SinceStart() float32 {
	return float32(time.Since(c.start).Seconds())
}

func (c *Clock) SinceStartDuration() time.Duration {
	return time.Since(c.start)
}
