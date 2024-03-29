[![codecov](https://codecov.io/gh/nikolaydubina/openapi-inline-examples/branch/main/graph/badge.svg?token=J97ET3LIQA)](https://codecov.io/gh/nikolaydubina/openapi-inline-examples)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/nikolaydubina/openapi-inline-examples/badge)](https://securityscorecards.dev/viewer/?uri=github.com/nikolaydubina/openapi-inline-examples)

> How do I add JSON examples to `openapi.yaml` from `.json` files?

Add to your `openapi.yaml` annotation `#source <filepath>` like

```yaml
openapi: 3.0.0
info:
  version: 1.0.0
  title: Swagger Marvelstore

paths:
  /users:
    get:
      responses:
        '200':
          description: A user object.
          content:
            application/json:
              schema:
                type: object
                properties:
                  id:
                    type: integer
                    format: int64
                    example: 4
                  name:
                    type: string
                    example: Jessica Smith
              examples:
                basic:
                  value: #source testdata/user-basic.json
        '400':
          description: The specified user ID is invalid (not a number).
          content:
            application/json:
              examples:
                basic 400:
                  value: #source testdata/error-400.json
```

.. then run 

```
$ go install github.com/nikolaydubina/openapi-inline-examples@latest
$ cat openapi.yaml | openapi-inline-examples > openapi.new.yaml
```

.. which will produce

```yaml
openapi: 3.0.0
info:
  version: 1.0.0
  title: Swagger Marvelstore

paths:
  /users:
    get:
      responses:
        '200':
          description: A user object.
          content:
            application/json:
              schema:
                type: object
                properties:
                  id:
                    type: integer
                    format: int64
                    example: 4
                  name:
                    type: string
                    example: Jessica Smith
              examples:
                basic:
                  value: {"id":42,"name":"Nick Fury"} #source testdata/user-basic.json
        '400':
          description: The specified user ID is invalid (not a number).
          content:
            application/json:
              examples:
                basic 400:
                  value: {"errors":[{"message":"resource not found","status":400,"translation_data":"some translation identifier"}]} #source testdata/error-400.json
```

.. which renders nicely as multiple examples that you can select

![example-preview](docs/example.png)

Why would anyone need this?

- Keep your OpenAPI up to date with values you use in tests
- Can run multiple times
- UNIX filter
- Does not corrupt anything if fails at any stage
- 100% test coverage
