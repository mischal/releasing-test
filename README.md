# releasing-test


## ghcr
https://docs.github.com/en/free-pro-team@latest/packages/managing-container-images-with-github-container-registry/pushing-and-pulling-docker-images




## github actions

### pullrequest.yaml
build only für pullrequests

### release.yml
erstellt ein docker image und pusht es in die ghcr

- wird ein tag gesetzt = release
- kann auch manuell gestartet werden (z.b. für images ab branches)

