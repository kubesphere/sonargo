# sonargo

![GitHub version](https://img.shields.io/badge/version-v0.0.1-brightgreen.svg?logo=appveyor&longCache=true&style=flat)
![](https://sonarcloud.io/api/project_badges/measure?project=kubesphere_sonargo&metric=alert_status)  ![](https://goreportcard.com/badge/github.com/magicsong/sonargo)
[![GoDoc](https://godoc.org/github.com/magicsong/sonargo/sonar?status.svg)](https://godoc.org/github.com/kubesphere/sonargo/sonar)

Client of [sonarqube api](https://sonarcloud.io/web_api) in golang
## Compatibility
As `sonarqube api` is updated frequently, this client now only support [v7.4](https://www.sonarqube.org/sonarqube-7-4/). Forward compatibility is not  guaranteed in this version (To be done).

## Usage

```go
import github.com/kubesphere/sonargo/sonar
```

Construct a new Sonarqube client, then use the various services on the client to access different parts of the Sonarqube API. For example, to list all projects:

```go
sonarURL := os.Getenv("SONAR_URL")
if sonarURL == "" {
    fmt.Println("Sonar URL has not been set")
    os.Exit(1)
}
client, err := sonargo.NewClient(sonarURL+"/api", "user", "password")
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}
//set options
opt := &ProjectsSearchOption{
    AnalyzedBefore:    "",
    OnProvisionedOnly: "",
    P:                 "",
    ProjectIds:        "",
    Projects:          "",
    Ps:                "",
    Q:                 "",
    Qualifiers:        "",
}
v, _, err := client.Projects.Search(opt)
if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}
fmt.Print(v.Components[0])
```

Other services is also like `Projects`, using following steps:
1. Get a client  by use `NewClient()`
2. Get the service by simply using dot. The service name is `CamelCase` of the name in sonarqube api
3. Use functions of each service to do what you want to do

## Notes
- Most code in first release version is generated by code. So some naming is not suitable. See [generate-go-for-sonarqube](https://github.com/magicsong/generate-go-for-sonarqube)
- Sonarqube community does not build a v7.4 docker image for public, so a v7.4 dockerfile is provided in [build](https://github.com/kubesphere/sonargo/tree/master/build) folder. 

## ToDo
* [ ] Add godoc for each api
* [ ] Add more tests
* [ ] test forward compatibility
## Author

宋雪涛 (Xuetao Song, magicsong@yunify.com)

## Licence
Copyright 2018 MagicSong 

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.