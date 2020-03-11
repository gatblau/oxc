# Onix Web API go client

A go client for the Onix Web API.

An example of how to use is below:

```go 
import "github.com/gatblau/oxc"

model := &oxc.Model {
    Key:         "test_model",
    Name:        "Test Model",
    Description: "Test Model",
}
    
result, err := client.PutModel(model)
```

More examples can be found [here](client_test.go).