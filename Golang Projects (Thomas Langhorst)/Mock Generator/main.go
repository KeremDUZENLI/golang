package main

import (
	"fmt"
	"mockgenerator/reader"
	"mockgenerator/writer"
)

func main() {
	interfaceLines, err := reader.Read("./service/service.go")
	if err != nil {
		panic(err)
	}

	for i, line := range interfaceLines {
		fmt.Printf("line: %d: %s\n", i+1, line)
	}

	mockDefinition := fakeService()
	writer.Write("./service/servicefromtemplate_mock.go", "./mock.tmpl", mockDefinition)
}

func fakeService() writer.MockDefinition {
	return writer.MockDefinition{
		Package: "service",
		Name:    "service",
		FunctionDefinitions: []writer.FunctionDefinition{
			{
				Name:             "Login",
				Params:           "username string, password string",
				CalledLine:       "args := m.Called(username, password)",
				ReturnValues:     "error",
				TypeCastingLines: []string{},
				ReturnLine:       "return args.Error(0)",
			},
			{
				Name:         "ListUsers",
				Params:       "",
				CalledLine:   "args := m.Called()",
				ReturnValues: "([]User, error)",
				TypeCastingLines: []string{
					"var u []User",
					"if args.Get(0) != nil {",
					"\tu = args.Get(0).([]User)",
					"}",
				},
				ReturnLine: "return u, args.Error(1)",
			},
		},
	}
}
