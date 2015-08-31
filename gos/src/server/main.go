package main

import (
	"app/register"
	"app/register/callbacks"
	"gslib"
)

func main() {
	register.Load()
	register.RegisterDataLoader()
	register.CustomRegisterDataLoader()
	callbacks.RegisterBroadcast()
	gslib.Run()
}
