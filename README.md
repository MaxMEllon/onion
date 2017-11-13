# onion (玉ねぎ)

##### *Onion is able to check the format of mail in tmng lab.*

> <img width="300" src="./.github/logo.png"/>

INSTALLATION
---

```
$ go get -u github.com/onion/cmd/onion
```

USAGE
---

#### CLI

```
$ onion *.tmng
```

**Please execute with `--fix` option, if you want to auto fix to document file**

```
$ onion *.tmng --fix  # auto fix!!
```

#### Vim

```vim
set makeprg=onion \-f=errorformats\ %
set errorformat=%f\|%l\ col\ %c\|\ %m
```

TODO
---

- lint
  - [x] `[新規], [継続], [完了]`
  - [x] `__◎ タイトル`
  - [x] `less than 75 character`
  - [x] `dont use ambi width space char`
  - [x] `dont use ◯ , should use ○`
  - [ ] `○____タイトル`
  - [ ] `2000.00.00(__)`

- autofix

LICENSE
---

[MIT License](./LICENSE.txt)
