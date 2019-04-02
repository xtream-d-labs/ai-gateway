# Local development

## 1. Codegen with the OpenAPI spec file

Using [openapi.yaml](https://github.com/rescale/scaleshift/blob/master/spec/openapi.yaml), you can generate its source code.

```console
docker-compose --file tools/dev/utility.yml run --rm codegen-api
docker-compose --file tools/dev/utility.yml run --rm codegen-web
```

## 2. Resolve dependencies

```console
docker-compose --file tools/dev/utility.yml run --rm deps
```

## 3. Run the application

```console
docker-compose --file tools/dev/runtime.yml up
open http://localhost:8080
```

## 4. Test it locally

```console
docker-compose --file tools/dev/utility.yml run --rm test-go-lint
docker-compose --file tools/dev/utility.yml run --rm test-go-unit
```
