package main

type Nginx struct {
	application        *Application
	MaxAllowedRequests int
	rateLimiter        map[string]int
}

type Responses struct {
	Code int
	Msg  string
}

func NewNginxServer() *Nginx {
	return &Nginx{
		application:        &Application{},
		MaxAllowedRequests: 2,
		rateLimiter:        make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) *Responses {
	allowed := n.CheckRateLimiting(url)
	if !allowed {
		return &Responses{
			Code: 403,
			Msg:  "Not Allowed",
		}

	}
	return n.application.HandleRequest(url, method)
}

func (n *Nginx) CheckRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.MaxAllowedRequests {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
