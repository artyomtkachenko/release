# release
A simple web service to build and publish rpm, deb and msi packages.
Inspired by Ruby based FPM project.

## Why?

* Would like to get rid of all Ruby dependencies like RVM, Ruby native libs and gems.
* Do it as a web service instead of having it as a CLI utility.
* Add more types of packages, like msi for Windows.
* Golang is fun.
* Golang is really fun :).

## Usage example

### Building RPM
```bash
curl -X POST -F "data=@path/to/artefact.zip" \
             -F "config=@without_scripts.json" \
             -F "post=@path/to/first-script.sh" \
             -F "preun=@path/to/second-script.sh" \
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

## Todo
* Upload to webdav and nexus.
* Revisit the DoBuild handler to make it more RAM efficient.
* Better config handling
* Add debugging functionality.
* Tests, tests, tests.
* Add benchmarks.
* Add more security.
* Add msi packaging for Windows.
* Add deb packaging for Debian based distros.
