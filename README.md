# Onix Web API go client

A go client for the Onix Web API.

An example of how to use is below:

```go
package main

import (
    "github.com/gatblau/oxc"
)
```

```go 
model := &Model {
    Key:         "test_model",
    Name:        "Test Model",
    Description: "Test Model",
}
    
result, err := client.putModel(model)
```

More examples can be found [here](client_test.go).