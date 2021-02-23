# gotest [![Build Status](https://travis-ci.com/mfroes/gotest.svg?token=onas63qBL39kCorPgzps&branch=main)](https://travis-ci.com/mfroes/gotest)

Simple http server written in Go
## Overview

This is a boilerplate written in Go for http endpoints.

* Dockerised service written in Go.
* CI Pipeline which triggers on push to master branch

## Architecture

Continuous Integration Architecture

![arch](docs/arch.png)

## Processes

List of the processes used.

| Requirement       | Tech or Process                                              |
| ----------------- | ------------------------------------------------------------ |
| Development Model | [Acceptence Test Driven Development](https://en.wikipedia.org/wiki/Acceptance_test%E2%80%93driven_development) |
| Branching Model   | [GitHub Flow](https://githubflow.github.io/) |
| Versioning        | [Semantic Versioning](https://semver.org/)                   |

## Technologies

List of the technologies used.

| Requirement                 | Tech or Process                                              |
| --------------------------- | ------------------------------------------------------------ |
| Language                    | [Go](https://golang.org/doc/) |
| Dependency Management       | [go.mod](https://golang.org/doc/modules/gomod-ref) |
| Linters                     | [go lint](https://github.com/golang/lint) |
| Formatter                   | [gofmt](https://golang.org/cmd/gofmt/) |
| Testing Framework           | [go test](https://golang.org/doc/code#Testing) [go mock](https://github.com/golang/mock)|
| SCM                         | [Github](https://github.com/) |
| Packaging                   | [Docker](https://www.docker.com/) |
| CI Pipeline                 | [travis CI](https://travis-ci.org) |


| Supporting tools | Description |
| --------------------------- | ------------------------------------------------------------ |
| [git](https://git-scm.com) | distributed version control system |
| [docker](https://www.docker.com) | tool for building and running virtual containers |
| [jq](https://stedolan.github.io/jq/) | lightweight command-line JSON processor |

## Project structure

| Directory | Description |
| --------------------------- | ------------------------------------------------------------ |
| ./       | project files including go |
| bin/     | stand alone scripts that are run by the pipeline and can be run locally |
| docs     | documentation support files |
| pkg      | go project support files |
| routes   | contains all the handlers for processing http requests |
| METADATA | File containing the VERSION number and Description as requested |

## Development

All the development have been set up so there would be a minimal amount of time and number of tools a developer would need.
The only software requires are the ones described on the `Supporting tools` section.
By doing this, we can guarantee that the environment on dev, on the pipelines and production are the same.

### Development workflow

The deployment process is descibed only to the the continuous integration environment.

1. Check out latest version of code

```bash
git pull
```

2. Write tests / acceptence criteria (ATDD)

3. Tests fail :(

```bash
./bin/test
```

4. Develop

5. Tests pass :)

```bash
./bin/test
```

6. Commit code

```bash
git commit -a -m "my new endpoint"
```

7. if any dependency or metadata, as release number, have been changed, re-run build.

```bash
./bin/build
```

8. Push code

```bash
git push
```

9. Pipeline will now run and deploy to the CI environment

10. Remember your number one priority in TBD is to not break the build.  So make sure it is all good before moving onto the next story.

## Updating release number

To generate a new release number, you will have to manually edit the METADATA file located in the root directoruy of the project and update the version using semantic versioning. On that file you will find the versioning and the description of the application that would be served by the /status endpooint.


## Packaging

Packaging is done using docker.  It is automated in the pipeline but if you wish to build or run the image manually you can use the following commands.

* Build docker image with a version
this will only build a `gotest:latest` image.

```bash
./bin/build
```

* Run the container

```bash
./bin/run-local
```

* Debug container

```bash
docker-compose run --rm --service-ports shell
```

## Metadata decision

At the moment all metadata is being passed at compilation time to the application by using the `-ldflags="-X var=value` as these information are immutable and should not be able to be changed on runtime.
This is the only used case that I would use this solution.

## Limitations and risks

* As the development environment is automated for fast developer set up time, lots of processes have been hidden, abstracted, from the developer which may cause issues when things go wrong as developers wouldnt undestand each step of what is being done. The reason for it being there is for simplicity and repeatability of environment and not to promote lack of knowledge on the application.

* there are no restrictions on method type on any http endpoint request at the moment. they both accept all method types.

* status endpoint is not formatted, it's just a json at the moment.
if it were to have users manually checking i would format it, if just for scrapers to look, id leave as is.

* because the need of setting up the metadata for the docker image(image labels), the METADATA processing is being done on a shellscript and passed as parameters to docker so it could be set as docker label.

* As there doesnt seem to be a limitation on the tags/releases with the same number, i've followed the semantic version as much as i could and also added a suffix to the tag names with the short git commit sha. example: v1.0.0.QWERTY and v1.0.0.YTREWQ are acceptable tags and will be pushed together with the v1.0.0 tag
this was done to try prevent conflicting tags and release numbers.

* At the moment there is no automatic semantic versioning done by the pipeline once application is released/merged into main. This would be a good implementation to add into the release phase of the pipeline. would probably implement by using pre-commit hooks or any tool similar.

* No non-functional tests (performance/security)

* No integration tests, not really required in this example... but could be something for the future.

* could mock the `pkg.GetMetadata()` call from the status_test.go, but how it is implemented on the code, i did not see any issue at the moment.