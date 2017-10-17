package certificatetpr

import (
	"testing"
	"time"

	"k8s.io/client-go/kubernetes/fake"

	"github.com/giantswarm/micrologger/microloggertest"
)

// Test_NewService tests the NewService function.
func Test_NewService(t *testing.T) {
	tests := []struct {
		config func() ServiceConfig

		expectedErrorHandler func(error) bool
	}{
		// Test that providing neither a kubernetes client,
		// or a logger, returns an error.
		{
			config: func() ServiceConfig {
				config := DefaultServiceConfig()

				config.K8sClient = nil
				config.Logger = nil

				return config
			},

			expectedErrorHandler: IsInvalidConfig,
		},

		// Test that only providing a kubernetes client returns an error.
		{
			config: func() ServiceConfig {
				config := DefaultServiceConfig()

				config.K8sClient = fake.NewSimpleClientset()
				config.Logger = nil

				return config
			},

			expectedErrorHandler: IsInvalidConfig,
		},

		// Test that only providing a logger returns an error.
		{
			config: func() ServiceConfig {
				config := DefaultServiceConfig()

				config.K8sClient = nil
				config.Logger = microloggertest.New()

				return config
			},

			expectedErrorHandler: IsInvalidConfig,
		},

		// Test that providing both a kubernetes client and a logger,
		// returns no error.
		{
			config: func() ServiceConfig {
				config := DefaultServiceConfig()

				config.K8sClient = fake.NewSimpleClientset()
				config.Logger = microloggertest.New()

				return config
			},

			expectedErrorHandler: nil,
		},

		// Test that providing a kubernetes client and a logger,
		// and a timeout,
		// returns no error.
		{
			config: func() ServiceConfig {
				config := DefaultServiceConfig()

				config.K8sClient = fake.NewSimpleClientset()
				config.Logger = microloggertest.New()

				config.WatchTimeOut = 5 * time.Second

				return config
			},

			expectedErrorHandler: nil,
		},
	}

	for index, test := range tests {
		config := test.config()
		service, err := NewService(config)

		if err != nil && test.expectedErrorHandler == nil {
			t.Fatalf("%d: unexpected error returned: %s\n", index, err)
		}
		if err != nil && !test.expectedErrorHandler(err) {
			t.Fatalf("%d: incorrect error returned: %s\n", index, err)
		}
		if err == nil && test.expectedErrorHandler != nil {
			t.Fatalf("%d: expected error not returned\n", index)
		}

		if test.expectedErrorHandler == nil && service == nil {
			t.Fatalf("%d: no error handler specified, but no service returned")
		}
	}
}
