// +build acceptance

package app_test

import (
	"github.com/DATA-DOG/godog"
	"github.com/checkinhq/checkin/test"
	"github.com/checkinhq/checkin/test/acceptance"
)

func init() {
	test.RegisterFeaturePath("../features")
	test.RegisterFeatureContext(FeatureContext)
}

func FeatureContext(s *godog.Suite) {
	grpcFeatureContext := new(acceptance.GrpcFeatureContext)

	grpcFeatureContext.FeatureContext(s)

	// Add steps here
}
