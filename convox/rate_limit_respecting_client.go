package convox

import (
	"io"
	"strings"
	"time"

	"github.com/convox/rack/pkg/structs"
	"github.com/jpillora/backoff"
)

//NewRateLimitRespectingClient will wrap the given client with rate-limit handling goodness.
func NewRateLimitRespectingClient(c Client) Client {
	if c == nil {
		panic("c Client must not be nil")
	}

	return &rateLimitRespectingClient{
		wrapped: c,
		backoff: &backoff.Backoff{
			Min:    500 * time.Millisecond,
			Max:    10 * time.Second,
			Factor: 2,
			Jitter: true,
		},
	}
}

type rateLimitRespectingClient struct {
	wrapped Client
	backoff *backoff.Backoff
}

func (r *rateLimitRespectingClient) retry(err error, attempts int) bool {
	if err == nil || attempts > 5 {
		r.backoff.Reset()
		return false
	}

	if strings.Contains(err.Error(), "Throttling: Rate exceeded") {
		time.Sleep(r.backoff.Duration())
		return true
	}

	return false
}

func (r *rateLimitRespectingClient) AppCreate(name string, opts structs.AppCreateOptions) (*structs.App, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.AppCreate(name, opts)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) AppGet(name string) (*structs.App, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.AppGet(name)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) AppDelete(name string) error {
	count := 0
	for {
		count++
		err := r.wrapped.AppDelete(name)

		if r.retry(err, count) {
			continue
		}

		return err
	}
}

func (r *rateLimitRespectingClient) ListFormation(app string) (structs.Formation, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.ListFormation(app)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) ListParameters(app string) (map[string]string, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.ListParameters(app)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) SetParameters(app string, params map[string]string) error {
	count := 0
	for {
		count++
		err := r.wrapped.SetParameters(app, params)

		if r.retry(err, count) {
			continue
		}

		return err
	}
}

func (r *rateLimitRespectingClient) GetEnvironment(app string) (structs.Environment, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.GetEnvironment(app)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) SetEnvironment(app string, body io.Reader) (structs.Environment, string, error) {
	count := 0
	for {
		count++
		rez, str, err := r.wrapped.SetEnvironment(app, body)

		if r.retry(err, count) {
			continue
		}

		return rez, str, err
	}
}

func (r *rateLimitRespectingClient) ResourceGet(name string) (*structs.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.ResourceGet(name)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) ResourceCreate(kind string, options structs.ResourceCreateOptions) (*structs.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.ResourceCreate(kind, options)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) ResourceUpdate(name string, options structs.ResourceUpdateOptions) (*structs.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.ResourceUpdate(name, options)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) ResourceDelete(name string) error {
	count := 0
	for {
		count++
		err := r.wrapped.ResourceDelete(name)

		if r.retry(err, count) {
			continue
		}

		return err
	}
}

func (r *rateLimitRespectingClient) ResourceLink(resourceName, app string) (*structs.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.ResourceLink(resourceName, app)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) ResourceUnlink(resourceName, app string) (*structs.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.ResourceUnlink(resourceName, app)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}
