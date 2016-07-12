# aivanenko vCard

## Simple golang and mongodb driven my online vCard

### Requirements

vCard version 0.1 requires Go >= 1.6 and MongoDB >=3.2.x

##### Installation

```sh
$ go get github.com/alexivanenko/web_cv/...
$ cd web_cv
$ go get -d ./...
$ make
```

Then open `http://127.0.0.1:9091` in your browser.
Admin area - `http://127.0.0.1:9091/admin/`, login and password in the config.ini file.
