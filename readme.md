## Quake

This project is an example of read a complex log file utilizing specific functions of Golang like slice and strings functions.

## Stack

- [Golang](https://go.dev) - Build fast, reliable, and efficient software at scale
- [Docker](https://www.docker.com) - Accelerate how you build, share, and run modern applications

## Project structure

```
$PROJECT_ROOT
├── reader               # Logic of read the log file
├── logs                 # To save log files
├── configs              # Dot env configs
└── utils                # Utilities functions
```


## Make commands:

Assuming that you have already cloned the project and the [Go](https://golang.org/doc/install) is installed, ensure that all dependencies are vendored in the project:

```
make install
```

To build the application:

```
make build
```

To run the application local:

```
make run
```

## Author

- [@pedrogaletti](https://www.github.com/PedroGaletti)