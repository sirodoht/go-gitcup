# go-gitcup

> Git repos backup utility

## Usage

### With one repo:

```sh
gitcup https://github.com/sirodoht/go-gitcup.git
```

### With multiple repos on a file

Example file `repos.txt`:
```txt
https://github.com/sirodoht/opencult.com.git
https://github.com/sirodoht/go-overlap.git
https://github.com/sirodoht/galois.git
```

Use with `-f` flag:
```sh
gitcup -f repos.txt
```

## License

MIT
