package convox

import (
	"io"
	"strings"
	"time"

	"github.com/convox/rack/client"
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

func (r *rateLimitRespectingClient) CreateApp(name, generation string) (*client.App, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.CreateApp(name, generation)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) GetApp(name string) (*client.App, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.GetApp(name)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) DeleteApp(name string) (*client.App, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.DeleteApp(name)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) ListFormation(app string) (client.Formation, error) {
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

func (r *rateLimitRespectingClient) ListParameters(app string) (client.Parameters, error) {
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

func (r *rateLimitRespectingClient) GetEnvironment(app string) (client.Environment, error) {
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

func (r *rateLimitRespectingClient) SetEnvironment(app string, body io.Reader) (client.Environment, string, error) {
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

func (r *rateLimitRespectingClient) GetResource(name string) (*client.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.GetResource(name)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) CreateResource(kind string, options map[string]string) (*client.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.CreateResource(kind, options)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) UpdateResource(name string, options map[string]string) (*client.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.UpdateResource(name, options)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) DeleteResource(name string) (*client.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.DeleteResource(name)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) CreateLink(app string, resourceName string) (*client.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.CreateLink(app, resourceName)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}

func (r *rateLimitRespectingClient) DeleteLink(app string, resourceName string) (*client.Resource, error) {
	count := 0
	for {
		count++
		rez, err := r.wrapped.DeleteLink(app, resourceName)

		if r.retry(err, count) {
			continue
		}

		return rez, err
	}
}
