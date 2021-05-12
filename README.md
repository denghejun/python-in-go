## How to embed python3 into golang (MacOS + Docker as an example)

ONLY `Docker` is required for this experiment. Install `Docker` if you haven't install it yet.

What this repo try to explain?

- How to call python3 code in golang
- How to call python3 distribution in golang

## Description

We have two repos

- [python-model](https://github.com/denghejun/python-model) : Hosting the python3 source code.
- [go-app](https://github.com/denghejun/python-in-go): Application based on golang, will interact
  with [python-model](https://github.com/denghejun/python-model)

## How to Run

We assume the situation is :

- `python-model` repo already did some changes, built and published a docker image named `python-model:v1.0.1` into
  docker registry.
- We need to change `go-app` also to apply the changes from `python-model`.

So that we can do the following steps:

- Update the `python-model` docker image version in file: ${ProjectRoot}/docker/go/Dockerfile (e.g.: python-model:
  v1.0.1)
- Update golang code if needed
- Build `golang-app` by running in
  terminal: `docker build --no-cache --progress=plain -t golang_server:v1.0.1 -f docker/go/Dockerfile .`
- Checking if the changes have been applied or not: `docker run golang_server:v1.0.1`

## Reference docs

- Python2 into Golang
    - https://www.datadoghq.com/blog/engineering/cgo-and-python/
    - https://github.com/sbinet/go-python
- Python3 into Golang
    - https://github.com/DataDog/go-python3
  