<!--
parent:
  order: false
-->

<div align="center">
  <h1> treasure-backend  repo </h1>
</div>

treasure project backed service , contracts event lister ,handler log event and business api 

**Tips**: need [Go 1.23.4+](https://golang.org/dl/)

## Install

### Install dependencies
```bash
go mod tidy
```

### build
```bash
make
```

### start
```bash
./treasure-backend
```

### Start the RPC interface test interface

```bash
grpcui -plaintext 127.0.0.1:8989
```

## Contribute

### 1.fork repo

fork treasure-backend to your github

### 2.clone repo

```bash
git@github.com:kitshy/treasure-backend.git
```

### 3. create new branch and commit code

```bash
git branch -C xxx
git checkout xxx

coding

git add .
git commit -m "xxx"
git push origin xxx
```

### 4.commit PR

Have a pr on your github and submit it to the treasure-backend repository

### 5.review

After the treasure-backend code maintainer has passed the review, the code will be merged into the treasure-backend library. At this point, your PR submission is complete

### 6.Disclaimer

This code has not yet been audited, and should not be used in any production systems.
