package main

import "ChainOfResponsibility/middlewares"

func main() {
	responser := &middlewares.Responser{}

	auth := &middlewares.Auth{}
	auth.SetNext(responser)

	logger := &middlewares.Logger{}
	logger.SetNext(auth)

	logger.Execute()
}
