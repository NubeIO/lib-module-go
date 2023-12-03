package parser

import "encoding/json"
import "github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"

func SerializeArgs(args nargs.Args) (string, error) {
	argsData, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	argsString := string(argsData)
	return argsString, nil
}

func DeserializeArgs(args string) (*nargs.Args, error) {
	deserializedArgs := nargs.Args{}
	if len(args) == 0 {
		return &deserializedArgs, nil
	}
	err := json.Unmarshal([]byte(args), &deserializedArgs)
	if err != nil {
		return nil, err
	}
	return &deserializedArgs, nil
}
