# simplefileserver
[![Go Report Card](https://goreportcard.com/badge/github.com/rbague/simplefileserver)](https://goreportcard.com/report/github.com/rbague/simplefileserver)
[![MIT License](https://badgen.net/github/license/rbague/simplefileserver)](https://github.com/rbague/simplefileserver/blob/master/LICENSE)

A simple zero-configuration static file server.

## Install
Download the latest binary for your platform from the [releases page](https://github.com/rbague/simplefileserver/releases), and run it from the command line

### Go
```sh
GO111MODULE=on go get -v -u github.com/rbague/simplefileserver
```

### From source
```sh
git clone https://github.com/rbague/simplefileserver.git && cd simplefileserver
make setup build-current # generates sfs binary
```

### Cross-Compile
```sh
git clone https://github.com/rbague/simplefileserver.git && cd simplefileserver
make # will generate all the binaries under bin with the following format sfs-GOOS-GOARCH
```

## Configuration
The server is ready to be used without any configuration, however, there a few options to customize it:

Name | Short | Required | Default | Description
--- | :---: | :---: | :---: | ---
help | - | false | - | Usage instructions
directory | d | false | `.` | The directory where to serve the files from
path | - | false | `/` | The URL path where files should be served from
host | h | false | - | The host where to bind the server to
port | p | false | `8080` | The port where the server should listen to
cert-file | - | false *(`true` if key-file is present)* | - | SSL certificate file
key-file | - | false *(`true` if cert-file is present)* | - | SSL certificate's private key

## HTTPS
In order to run the server with TLS activated, pass both the `cert-file` and `key-file` flags to the startup command.

to enable it locally I recommend using the [mkcert](https://github.com/filosottile/mkcert) tool to gerate the files, or you can generate them with the following command
```sh
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365
```

## License
The software is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
