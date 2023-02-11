# schedulii

A web application that allows you to coordinate meeting times with your friends. We provide easy import from Google (Microsoft and other platforms are TBD) and a way to easily add new events into your calendars.

### Tools and construction
The backend of the application is written in Golang, using the popular Gin library to run our web server. We save non-sensitive information about users and the events they create on a PostgreSQL instance running on AWS RDS. Session storage for users is done using a redis instance running on AWS Elasticache.

The frontend is written in Typescript and React.

## Deployment
### CI
Currently we are using Github Actions to run CI checks on code before merges to main. With each merge, we also create an updated Dockerfile with deployable binaries to our image repository hosted on AWS Elastic Container Repository (ECR).

### CD
We are planning to frequently update the live application with every merge by using AWS CodeDeploy on an EC2 autoscaling group. The live domain URL will be hosted on AWS Route53.

## Running locally
To run a local copy of the server, you will need access to database and oauth credentials that need to be loaded into your environment. If you are interested in collaborating, message the current collaborators for the secrets and place them in a .env file at the root of the project. We have a Makefile that automatically loads the environment variables into the local environment before running the server.

## Setting up a local database for testing
1. `docker pull postgres`
2. `docker run --name postgres -e POSTGRES_PASSWORD=<password> -d -p 5432:5432 postgres`
3. `psql`
4. Copy and paste the `.psql` contents from the assets folder to load data into postgres
5. Update env variable reference to database to `postgresql://postgres:<password>@localhost:5432/postgres`
