## Teste

```shell
go test -v -coverprofile cover.out ./...
go tool cover -func cover.out
# go tool cover -html cover.out -o cover.html
# open cover.html
```

erro do elasticsearch ao subir sonarqube com o dockercompose
https://stackoverflow.com/questions/57175156/cannot-start-sonarqube-because-of-memory-problem
sysctl -w vm.max_map_count=262144
sysctl -w fs.file-max=65535

admin:pedeaiAPI123!