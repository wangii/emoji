emoji
=====

[![Build Status](https://travis-ci.org/chnlr/emoji.svg?branch=master)](https://travis-ci.org/chnlr/emoji) [![GoDoc](http://godoc.org/github.com/chnlr/emoji?status.svg)](http://godoc.org/github.com/chnlr/emoji)

A golang library for emoji chars converting.

INSTALL
-------

```bash
$ go get -u github.com/chnlr/emoji
```

EXAMPLE CODE
------------

```go
package main

import (
    "fmt"

    "github.com/chnlr/emoji"
)

func main() {
    fmt.Println(emoji.EmojiTagToUnicode(`:+1:`)) // => 👍
}
```

DOCUMENTATION
-------------

  * <http://godoc.org/github.com/chnlr/emoji>

Author
------

Naoki OKAMURA a.k.a nyarla <nyarla@thotep.net>

LICENSE
-------

MIT

USGIN LIBRARIES
---------------

  1. `map.go` in this package is base by <https://github.com/melborne/emot/blob/master/lib/emot/map.rb> (ruby's [emot](https://github.com/melborne/emot))
     * ruby's [emot](https://github.com/melborne/emot) is under the [MIT](https://github.com/melborne/emot/blob/master/LICENSE.txt)
     * Copyright (c) 2014 kyoendo.
  2. this library is using to twemoji's SVG link with MaxCDN for generates emoji image link.
     * [twemoji](https://github.com/twitter/twemoji)'s Graphics is under the CC-BY 4.0: <https://creativecommons.org/licenses/by/4.0/>
     * Copyright 2014 Twitter, Inc and other contributors

