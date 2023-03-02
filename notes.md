`internal/comment/comment.go` fo;llow
**accepting interfaces - returning structs**
`git remote add origin git@github.com:7350206/go-api-pg.git`

### context propagation
btween diff layers of service becomes invaluable
observability (comment id, trace id, ...): 
- when all services are talk to each other - becomes diff to trace full request flow to identify where error happened

### dockerize app
- create Dockerfile
- `docker build -t go-api-pg .`

### compose
- new file `docker-compose.yml`

`docker compose version`
Docker Compose version v2.15.1

`docker-compose up`

can now start the database server using:
`pg_ctl -D /var/lib/postgresql/data -l logfile start`


### taskfiles [instead of make]
[github](https://github.com/go-task/task/blob/master/Taskfile.yml)
[site](https://taskfile.dev/)
`go install github.com/go-task/task/v3/cmd/task@latest`
make a  `Taskfile.yml` file

