package main

type Application struct {
}

func (a *Application) HandleRequest(url, method string) *Responses {
	if url == "/app/status" && method == "GET" {
		return &Responses{
			Code: 200,
			Msg:  "OK",
		}

	}
	if url == "/create/user" && method == "POST" {
		return &Responses{
			Code: 201,
			Msg:  "User Created",
		}
	}
	return &Responses{
		Code: 404,
		Msg:  "Not Found",
	}
}
