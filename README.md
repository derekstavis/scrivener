# Scrivener
> Display [goreporter][goreporter] results in terminal.

## Using

Install it:

```sh
$ go get github.com/derekstavis/scrivener
```

Run `goreporter` in your project with JSON output:

```sh
$ goreporter -p ../scrivener -f json
```

>Be sure to replace `../scrivener` by your project's package

Run `scrivener` with JSON file as argument:

```sh
$ scrivener scrivener-1497297258.json

+-----------+----------------------------------+---------------------------------------------------------------+
|   TEST    |             SUMMARY              |                            ERROR                              |
+-----------+----------------------------------+---------------------------------------------------------------+
| CycloTips | github.com/derekstavis/scrivener | /home/derek/src/github.com/derekstavis/scrivener/main.go:91:1 |
+           +                                  +---------------------------------------------------------------+
|           |                                  | /home/derek/src/github.com/derekstavis/scrivener/main.go:74:1 |
+           +                                  +---------------------------------------------------------------+
|           |                                  | /home/derek/src/github.com/derekstavis/scrivener/main.go:45:1 |
+-----------+----------------------------------+---------------------------------------------------------------+
```

You can also ignore some tests using `--exclude`:

```sh
$ scrivener scrivener-1497297258.json --exclude=CycloTips

+-----------+----------------------+------------------------------+
|   TEST    |       SUMMARY        |           ERROR              |
+-----------+----------------------+------------------------------+
+-----------+----------------------+------------------------------+
```

---
[goreporter]: https://github.com/wgliang/goreporter
