package sources

import (
	"testing"

	"github.com/OWASP/Amass/requests"
	"github.com/OWASP/Amass/resolvers"
)

func TestWayback(t *testing.T) {
	if *networkTest == false {
		return
	}

	c := setupConfig(domainTest)
	bus, out := setupEventBus(requests.NewNameTopic)
	defer bus.Stop()

	pool := resolvers.NewResolverPool(nil)
	defer pool.Stop()

	srv := NewWayback(c, bus, pool)

	result := testService(srv, out)
	if result < expectedTest {
		t.Errorf("Found %d names, expected at least %d instead", result, expectedTest)
	}
}
