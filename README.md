# learn-cicd (Notely)

![workflow badge](https://github.com/bailey4770/learn-cicd/actions/workflows/ci.yml/badge.svg)

Notely is a simple CRUD app that allows users to create accounts and write notes.
This project uses Google Cloud Platform to host the application code and Turso
to host the database.

The purpose of the project is to explore how to write CI/CD pipelines. The CI
section ensures that all pull requests meet certain code standards, such as passing
all tests, and the CD section automatically pushes accepted code changes to the
remote app server and database.

The automation of this pipeline helps ensure teams can focus more on delivering
bug fixes and new features, rather than performing 'op' tasks. This setup is much
more efficient and less error-prone than manually SSHing into servers to ensure
the latest docker image is being used in production. Therefore, it is very important
to incorporate this into real projects.

To ensure costs are kept to a minimum, all cloud services have been deleted. This
project was intended to learn the principles of CI/CD, not to deploy a real web
app.

## Continuous Integration

Project runs CI when pull request is made to main.

CI Pipeline runs following tasks:

- Checks out code.
- Sets up Go.
- Discovers and runs all test funcs in the project.
- Formats code with `go fmt`
- Runs `staticcheck` for linting tests.
- Runs `gosec` to help catch glaring security failures.

## Continuous Deployment

When main is pushed to, CD action is run.

CD Pipeline runs following tasks:

- Checks out code.
- Sets up Go.
- Installs goose.
- Builds binary from new code.
- Authorises and sets up GCloud SDK.
- Builds new docker image and pushes it to Google Artefact Registry.
- Runs any SQL migrations with Goose.
- Deploys new image to Cloud Run.
