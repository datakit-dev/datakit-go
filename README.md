# datakit-go
Go types for DataKit Cloud GraphQL services.

This Go module is used primarily internally for `dtkt` (DataKit CLI) and other services. Can be used for Go projects looking to integrate with DataKit Cloud.

Install using `go get`:

```
go get github.com/datakit-dev/datakit-go
```

Code generated from DataKit Cloud GraphQL Schema using: https://github.com/jkrajniak/graphql-codegen-go using the -config option.

Command to generate types:

```
graphql-codegen-go -config ./config/codegen.yaml
```