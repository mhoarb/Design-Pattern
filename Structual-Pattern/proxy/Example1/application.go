package main

type Application struct {
}

func (a *Application) HandleRequestApplication(url, method string) *Response {
	if url == "/app/status" && method == "GET" {
		return &Response{
			Code: 200,
			Msg:  "OK",
		}
	}
	if url == "/create/user" && method == "POST" {
		return &Response{
			Code: 201,
			Msg:  "User Created",
		}
	}
	return &Response{
		Code: 404,
		Msg:  "Not OK",
	}

}
