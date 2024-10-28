# Boot.dev CI/CI Course

![CI](https://github.com/EnsIaver/boot-dev-cicd/actions/workflows/ci/badge.svg)

This repository comprises the starting template, implementation and other materials created throughout the [boot.dev CI/CD course](http://boot.dev/learn/learn-ci-cd).

## Local Development

Make sure you're on Go version 1.20+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8000"
```

Run the server:

```bash
go build -o notely && ./notely
```
