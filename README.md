go-leat
====

Tools for [extending length](https://en.wikipedia.org/wiki/Length_extension_attack)
for Hash functions based on Merkle-Damgård constructions(md5, sha-1,
sha-2)

Currently implemented:
- [X] md5
- [ ] sha1
- [ ] sha256
- [ ] sha512

### Install
 Only requirement for project is Golang
 
 Run to install: `go get github.com/nodar-chkuaselidze/go-leat/cmd/leat`


### Usage
You want to extend some md5 hash and know only length of original text, say:
  * text was: `hello`
  * it's md5 is: `5d41402abc4b2a76b9719d911017c592`
  * We want to extend with: `\x20world` (\x20 is space)
  * length is: 5 (hello)
  * we want to format padding bytes like: `\x0`
  
We run `leat md5 5 5d41402abc4b2a76b9719d911017c592 ' world' '\x%x'` result will be:
```
de24e753e179eda92528062a19d30148
\x80\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x28\x0\x0\x0\x0\x0\x0\x0 world
```

You can check that md5 of `hello\x80\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x0\x28\x0\x0\x0\x0\x0\x0\x0 world` is `de24e753e179eda92528062a19d30148`
