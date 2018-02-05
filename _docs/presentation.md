<!-- $theme: default -->

Integration Tests in Go
===

<br><br><br><br><br>
<div style="text-align: right">Oleh Herych</div>
<div style="text-align: right">2018</div>

---

# What is integration testing

**Integration testing** tests integration or interfaces between components, interactions to different parts of the system such as an operating system, file system and hardware or interfaces between systems.

(c) ISTQB

---

# setup test environment
- Docker
- Minio server
- http.Transport{...}

---

# Time for Demo

```sh
docker-compose up -d
go test -v ./...
```


--- 

# Links
- https://divan.github.io/posts/integration_testing/
- http://github.com/gavv/httpexpect/
- github.com/h2non/gock