# Local development

## 1. Codegen with the OpenAPI spec file

Using [openapi.yaml](https://github.com/rescale-labs/scaleshift/blob/master/spec/openapi.yaml), you can generate its source code.

```console
docker-compose --file tools/dev/utility.yml run --rm codegen-api
docker-compose --file tools/dev/utility.yml run --rm codegen-web
```

## 2. Resolve dependencies

```console
docker-compose --file tools/dev/utility.yml run --rm deps-api
```

## 3. Test it locally

```console
docker-compose --file tools/dev/utility.yml run --rm lint-api
docker-compose --file tools/dev/utility.yml run --rm test-api
```

## 4. Generate RSA keys

```console
docker-compose --file tools/dev/utility.yml run --rm gen-private-key
docker-compose --file tools/dev/utility.yml run --rm gen-public-key
```

## 5. Run the application

```console
docker-compose --file tools/dev/runtime.yml up
open http://localhost:8080
```
