## Teste

```shell
go test -v -coverprofile coverage.out ./...
go tool cover -func coverage.out
# go tool cover -html coverage.out -o cover.html
# open cover.html
```

erro do elasticsearch ao subir sonarqube com o dockercompose
https://stackoverflow.com/questions/57175156/cannot-start-sonarqube-because-of-memory-problem
sysctl -w vm.max_map_count=262144
sysctl -w fs.file-max=65535

admin:pedeaiAPI123!

sqp_9ab9301c990fead9702c66af8113d7fd0627018b

/home/filipe/projetos/sonarqube-scanner/sonar-scanner-cli-6.2.1.4610-linux-x64/sonar-scanner-6.2.1.4610-linux-x64/bin \
  -Dsonar.projectKey=pedeai-clientes \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.token=sqp_9ab9301c990fead9702c66af8113d7fd0627018b


/home/filipe/projetos/sonarqube-scanner/sem-java/bin/sonar-scanner   -Dsonar.projectKey=pedeai-clientes   -Dsonar.sources=.   -Dsonar.host.url=http://localhost:9000   -Dsonar.token=sqp_9ab9301c990fead9702c66af8113d7fd0627018b

https://stackoverflow.com/questions/52962493/how-to-exclude-golang-tests-structs-and-constants-from-counting-against-code-co