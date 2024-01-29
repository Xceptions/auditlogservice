The microservice runs on two ports

TCP: port 80 -> used for accepting events
Http: port 8080 -> used for querying the events

for the docker pass, use environment variables

When using pgadmin to view postgres db, use "host.docker.internal" as the hostname instead of localhost


For the http requests, I take care of:
1. Rate limiting
2. Caching
3. Versioning