package util

import "flag"

type FlagArgs struct {
	Env *string
}

func DeclareFlag() FlagArgs {
	return FlagArgs{
		Env: flag.String("env", "", "Environment to run the application in"),
	}
}