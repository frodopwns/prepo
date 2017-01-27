# Prepo

Prepo is a tool written in Golang that helps teams apply the labels they need to get work done to new
and existing github repositories.

To build:

Clone repo into $GOPATH/src/github.com/drud/prepo
From the repo's root directory run `go build`

To use:

Create a prepo.yaml file and fill it with your lavels like so:

```
labels:
  - name: needs docs
    color: 5319e7
  - name: needs manual testing
    color: fbca04
  - name: needs tests
    color: 006b75
  - name: priority 0
    color: e11d21
  - name: priority 1
    color: e99695
  - name: proposal
    color: 0052cc
  - name: showstopper
    color: e11d21
```

Then from the same directory run `./prepo org/repo_name`