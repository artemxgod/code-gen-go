# Code generation in golang

## Ifacemaker

> Download: `brew install ifacemaker`

#### Usage
- Creates interface for some impelentation. now when you add some method, just run go generate ./... and your interface is up to date

#### Steps
1. Create implementation for your interface
2. generate interface using ifacemaker

- command: `ifacemaker -f <file> -o <interface_filename> -i <inteface_name> -s <struct_name> -p <package_name>`

- [Example](./ifmaker)

## stringer

> Download: `go install golang.org/x/tools/cmd/stringer@latest`

#### Usage
- Stringer makes your type suitable for fmt.Stringer interface

#### Steps
1. create enum for type
2. generate String method for type

- command: `stringer -type=<typename> <filename>`

- [Example](./stringer/)

## mockgen

> Download: `go install go.uber.org/mock/mockgen@latest`

#### Usage
- mockgen allows you to isolate unit of code you are testing by replacing dependencies with contolled substitutes

#### Steps
1. Select interface for mocking
2. Mock the interface

- command: `mockgen -source=<filename> -destination=<filename> -package=<pkgname>`

- [Example](./mockgen/)