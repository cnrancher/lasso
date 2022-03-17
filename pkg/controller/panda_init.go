package controller

import (
	"os"
	"time"

	"github.com/rancher/lasso/pkg/log"
)

const (
	fastSlowRateLimiterSlowDelayEnv          = "CATTLE_LASSO_FSRL_SD"
	exponentialFailureRateLimiterMaxDelayEnv = "CATTLE_LASSO_EFRL_MD"
)

var (
	defaultFastSlowRateLimiterSlowDelay          = 2 * time.Minute
	defaultExponentialFailureRateLimiterMaxDelay = 30 * time.Second
)

func init() {
	log.Infof("got lasso SlowRateLimiterSlowDelay %s", getFastSlowRateLimiterSlowDelay())
	log.Infof("got lasso ExponentialFailureRateLimiterMaxDelay %s", getExponentialFailureRateLimiterMaxDelay())
}

func getFastSlowRateLimiterSlowDelay() time.Duration {
	d, err := time.ParseDuration(os.Getenv(fastSlowRateLimiterSlowDelayEnv))
	if err != nil || d <= 0 {
		return defaultFastSlowRateLimiterSlowDelay
	}
	return d
}

func getExponentialFailureRateLimiterMaxDelay() time.Duration {
	d, err := time.ParseDuration(os.Getenv(exponentialFailureRateLimiterMaxDelayEnv))
	if err != nil || d <= 0 {
		return defaultExponentialFailureRateLimiterMaxDelay
	}
	return d
}
