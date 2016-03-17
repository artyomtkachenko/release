# release
A simple web service to build and publish rpm packages 

## Why?

* Get rid of all Ruby dependencies like RVM, Ruby native libs and gems.
* Do it as a web service instead of having it as a CLI utility.
* Add more types of packages, like msi for Windows.
* Golang is fun.

## Usage example
#Not implemented yet. Should replace the origianl bad idea with the two steps apporach.

### Building RPM
```bash
curl -X POST -F "data=@path/to/artefact.zip" \
             -F "config=@without_scripts.json" \
             -F "script=@path/to/first-script.sh" \
             -F "script=@path/to/second-script.sh" \
             -F 'meta={"version": "1.0.1", "revision", "de1b85fbf4f855f48c6362ea26401e46" }' \
             localhost:9999/release/v1/build/rpm
```

### Building MSI
```bash
curl -X POST -F "data=@path\to\artefact.zip" \
             -F "config=@without_scripts.json" \
             -F 'meta={"version": "1.0.1", "revision", "de1b85fbf4f855f48c6362ea26401e46" }' \
             localhost:9999/release/v1/build/msi

```

# Original, bad implementation.
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
