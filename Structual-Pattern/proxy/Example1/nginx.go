package main

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

type Response struct {
	Code int
	Msg  string
}

func (n *Nginx) HandleRequestNginx(url, method string) *Response {
	allowed := n.CheckRateLimiting(url)
	if !allowed {
		return &Response{
			Code: 403,
			Msg:  "Not Allowed",
		}
	}
	return &Response{
		Code: 404,
		Msg:  "Not Found",
	}

}

func (n *Nginx) CheckRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
