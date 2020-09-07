# Deploying AWS

Based off Widdix AWS scripts - amazing work.  https://github.com/widdix/aws-cf-templates

Instructions on Main Page

| Order | File | Purpose |
| --- | --- | --- |
| 1 | vpc-2azs.yaml (https://github.com/widdix/aws-cf-templates/vpc/vpc-2azs.yaml) | Networking - Based on the Widdix Cluster |
| 2 | 2-aws-planbeer-storage.yaml | Storage Layer - Planbeer (3) can be removed without losing data |
| 3 | 3-aws-planbeer-cluster.yaml | Planbeer App - The App, DB, UI Layer |

# Running Containers

## Pre-Requisites

[Docker] (https://get.docker.com)

## Config Options

## PlanBeer Server
| Option | Default | Description | Required |
| --- | --- | --- | --- |
| DNS1 | N/A | Alternate name for Generate certs (should be picobrew.com) | Yes |
| DNS2 | N/A | Alternate name for Generate certs (should be www.picobrew.com) | Yes |
| PB_SESSION_PATH | "/planbeer/sessions" | Volume in containers to import sessions  | No |
| PB_RECIPE_PATH | "/planbeer/recipes" | Volume in container to import recipes  | No |
| PB_CERTS_PATH | "/planbeer/certs" | Volume in container to generate certificates (for machine)  | No |
| PB_ENABLE_TLS | "TRUE" | Enable TLS Server for Machine Communication (if false, the machine won't connect) | No |
| PB_PORT  | ":80" | Port for UI to connect  | No |  
| PB_TLS_PORT | ":443" | Port for machine to connect  | No |
| PB_DBADMIN | "admin" | Configured Username for CouchDB  | No |
| PB_DBPASSWORD | "password" | Configured Passowrd for CouchDB  | No |
| PB_DBSERVER | "localhost" | Address to CouchDB  | No |

## Dependencies

### CouchDB Example
docker run -p 5984:5984 -e COUCHDB_USER=admin -e COUCHDB_PASSWORD=password -v C:/users/me/couchdb/data:/opt/couchdb/data couchdb

### Container Example
docker run -p 443:443 -p 80:80 -e PB_ENABLE_TLS=TRUE -e PB_SESSION_PATH=/planbeer/sessions -e PB_RECIPES_PATH=/planbeer/recipes -e PB_DBSERVER=192.168.1.104 -e DNS1=picobrew.com -e DNS2=www.picobrew.com -e PB_CERTS_PATH=/certs -v C:/users/me/downloads/test:/planbeer planbeer:latest 

Note: You'll have to loadbalance the "/ui/*" on port to this service.   Otherwise you'll have a collision on port 80 when you run the UI.

### Planbeer UI Example
docker run -p 80:80 planbeerui:latest
