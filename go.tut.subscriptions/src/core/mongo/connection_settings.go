package mongo

import (
	"time"
)

type ConnectionSettings struct {
	connectionString string
	timeout          time.Duration
	retryCount       int
	// sleepBetweenRetry The time to wait between retries in seconds
	sleepBetweenRetry time.Duration
	// retrySleepTimeout the maximum time in seconds allowed for each retry step
	retrySleepTimeout time.Duration
}

func NewConnectionSettings(connectionString string) ConnectionSettings {
	return ConnectionSettings{
		connectionString: connectionString,
	}
}

func (settings ConnectionSettings) WithTimeout(duration time.Duration) ConnectionSettings {
	settings.timeout = duration
	return settings
}

func (settings ConnectionSettings) WithRetry(
	retryCount int,
	sleepBetweenRetry time.Duration,
	retrySleepTimeout time.Duration) ConnectionSettings {
	settings.retryCount = retryCount
	settings.sleepBetweenRetry = sleepBetweenRetry
	settings.retrySleepTimeout = retrySleepTimeout
	return settings
}
