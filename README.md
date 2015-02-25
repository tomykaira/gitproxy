# Gitproxy -- allow someone to push in place of you

**Use this software under your own responsibility.**

## provider

In case of GitHub, create a personal access tokens.
Reference: [Creating an access token for command-line use - User Documentation](https://help.github.com/articles/creating-an-access-token-for-command-line-use/).

Start the server with following environment variables.

- `PORT`: HTTP port on which the server runs.
- `GITPROXY_REMOTE_URL`: Remote url, like `https://github.com`
- `GITPROXY_ACCESSOR_USER`: Username used to authenticate who pushes in place of you.
- `GITPROXY_ACCESSOR_PASS`: Password.
- `GITPROXY_REMOTE_USER`: Your basic authentication account on the remote server. Username in case of GitHub.
- `GITPROXY_REMOTE_PASS`: Your password. Token in case of GitHub. Keep it secret.

```
go run main.go
```

## accessor

Receive the accessor user, pass and the server address.

Clone the target repository like `git clone https://github.com/USER/REPO.git`.

In the git directory, change push URL.

```
git remote set-url --push origin http://GITPROXY.example.com/USER/REPO.git
```

When you `git push`, git asks accessor username/password.
