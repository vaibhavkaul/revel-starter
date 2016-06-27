# Welcome to Revel Starter

## Getting Started

This strips out all the View/Index stuff from the revel starter. This can serve as a starter for a simple REST app

### Start the web server:

    get revel: go get github.com/revel/revel
    get revel command line: go get github.com/revel/cmd/revel
    
    Make sure $GOPATH is set and $GOPATH/bin is on the PATH

    revel run revel-starter

### Go to http://localhost:9000/orders and you'll see:

```json
[
    {
        "ID": 1,
        "Event": "Event 1",
        "UserId": 1
    },
    {
        "ID": 2,
        "Event": "Event 2",
        "UserId": 2
    },
    {
        "ID": 3,
        "Event": "Event 3",
        "UserId": 3
    }
]
```

### Description of Contents

The default directory structure of a generated Revel application:

    revel-starter               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates (Empty folder in this case)
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.

test

    Tests are kept in the tests directory.
    
### Running Tests

Tests can be run in 2 ways:

Interactively:
    
    revel run revel-starter
    navigate to /@tests
    
Non Interactively:

    revel test revel-starter
    
