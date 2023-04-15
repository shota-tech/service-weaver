# Service Weaver Tutorial

Tutorial for Service Weaver.

## What I learned
- What is Service Weaver.
- How to implement components and listeners.
- How to run an application in multiple processes.
- How to use a dashboard.

## Usage

### Install tools
```sh
make install-tools
```

### Run in a single process
Run the application in a single process.
```sh
go run .
```

Send a request to the application.
```sh
curl "localhost:12345/hello?name=Weaver"
```

Display the application status.
```sh
weaver single status
```

Open the dashboard.
```sh
weaver single dashboard
```

### Run in multiple processes
Run the application in a single process.
```sh
weaver multi deploy weaver.toml
```

Send a request to the application.
```sh
curl "localhost:12345/hello?name=Weaver"
```

Display the application status.
```sh
weaver multi status
```

Open the dashboard.
```sh
weaver multi dashboard
```

## References
- https://serviceweaver.dev/docs.html
