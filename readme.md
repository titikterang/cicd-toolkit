## Overview
* CI/CD Toolkit is small command line tool helper to integrate with vault secret kv management & github api
* We can use simple command to generate json/env/raw text file from vault secret kv
* Or to verify github PR approval status and then merge PR if pull request approval is validated

## Build
```shell
make build
```
it will build two binaries, binary file for linux and binary based on local OS

## Generate json file from vault
```shell
./cmd/toolkit/toolkit  -vault -output=conf.json -secret=vaultkv/data/yourdata
```

## Generate env file from vault
```shell
./cmd/toolkit/toolkit  -vault  -env -secret=vaultkv/data/envdata
```

## Generate raw text file from vault
```shell
./cmd/toolkit/toolkit  -vault  -raw -secret=vaultkv/data/rawtextfile
```

## Verify github PR approval
```shell
./cmd/toolkit/toolkit -approval -repo=ujunglangit-id/some-repo -id=23
```

## Merge github pull request
```shell
./cmd/toolkit/toolkit -merge -repo=ujunglangit-id/some-repo -id=23
```

## Merge squash github pull request
```shell
./cmd/toolkit/toolkit  -debug -squash -repo=ujunglangit-id/some-repo -id=23
```

## Hashicorp vault secret kv management
<hr/>
https://www.vaultproject.io/api-docs

## Github api integration for pr status check & merging
<hr/>
https://docs.github.com/en/rest/reference/pulls