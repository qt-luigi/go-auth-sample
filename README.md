# go-auth-sample

Authenticating Users samples.

Main target is GAE/Go 2nd generation environment.

Now available only Google Sign-in.

## Usage

Executing this program on your local machine, you use `dev_appserver.py` tool.

```sh
$ dev_appserver.py app.yaml
```

Deploying this program to GAE/Go 2nd-gen, you use `gcloud` command.

```sh
$ gcloud app deploy
```

### Google Sign-in

You must set Google Cloud Platform OAuth2 credentials Client ID and Client Secret to an environment variables.

When executing this program on GAE/Go 2nd-gen, Setting it to `env_variables:` section in app.yaml file. 

```sh
env_variables:
  GOOGLE_CLIENT_ID: "client_id"
  GOOGLE_CLIENT_SECRET: "client_secret"
```

When executing this program by `go run` or a binary, Using `export` command on bash.

```sh
$ export GOOGLE_CLIENT_ID=client_id
$ export GOOGLE_CLIENT_SECRET=client_secret
$ go run main.go
```

## License

MIT
