# Local development

## 1. Initialize th app

Codegen with the OpenAPI spec file, resolve dependencies & generate RSA keys.

Using [openapi.yaml](https://github.com/scaleshift/scaleshift/blob/master/spec/openapi.yaml), you can generate its source code.

```console
make init
```

## 2. Run the application

```console
make run
open http://localhost:8080
```

## 3. Test it locally

```console
make test/src
```
