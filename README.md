# zztgo

`zztgo` is a (not exactly finished) port of Adrian Siekierka’s source code [reconstruction of ZZT](https://github.com/asiekierka/reconstruction-of-zzt/) to Go. I created it using a [Pascal-to-Go converter](https://github.com/benhoyt/pas2go) that I wrote, as well as the [tcell](https://github.com/gdamore/tcell) terminal library for graphics.

To run it, simply clone the repo, type `go build`, and then run `zztgo`. To make it look a bit more like the real thing, you should install an [IBM EGA font](https://int10h.org/oldschool-pc-fonts/fontlist/#ibmega) and adjust the line spacing to zero. On macOS you can use the [zzt.terminal](https://github.com/benhoyt/zztgo/blob/master/zzt.terminal) Terminal settings file.

[**Read the full story here**.](https://benhoyt.com/writings/zzt-in-go/)
