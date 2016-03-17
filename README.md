# release
A simple web service to build and publish rpm packages 

## Why?

* Get rid of all Ruby dependencies like RVM, Ruby native libs and gems.
* Do it as a web service instead of having it as a CLI utility.
* Add more types of packages, like msi for Windows.
* Golang is fun.

## Usage example
First call:

```shell
curl -v -X POST -d "@without_scripts.json" localhost:8080/release/v1/init/rpm?version=1.0.1
returns 3117776977 
``` 

Next call:

```shell
curl -v -T apache-activemq-5.12.2-bin.zip  localhost:8080/release/v1/build/rpm/3117776977

build, sign and publish an rpm
``` 

## Todo
* Upload to webdav and nexus
* Revisit the DoBuild handler to make it more RAM efficient.
* Test cover for more than 80%.
* Add benchmarks
