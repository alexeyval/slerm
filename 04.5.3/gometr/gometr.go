package gometr

type HealthCheck struct {
	ServiceID string
	Status    string
}

type GoMetrClient struct {
	url string
	sec int
}

func NewGoMetrClient(url string, sec int) *GoMetrClient {
	return &GoMetrClient{
		url: url,
		sec: sec,
	}
}

func (c *GoMetrClient) GetMetrics() string {
	return ""
}

func (c *GoMetrClient) Ping() error {
	return nil
}

func (c *GoMetrClient) GetID() string {
	return c.url
}

func (c *GoMetrClient) Health() bool {
	for _, v := range []string{"1", "2", "4"} {
		if v == c.url {
			return true
		}
	}
	return false
}

func (c *GoMetrClient) getHealth() HealthCheck {
	return HealthCheck{
		ServiceID: c.GetID(),
		Status:    "",
	}
}
