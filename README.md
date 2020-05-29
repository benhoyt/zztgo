# zztgo

`zztgo` is a (not exactly finished) port of Adrian Siekierka’s source code [reconstruction of ZZT](https://github.com/asiekierka/reconstruction-of-zzt/) to Go. I created it using a [Pascal-to-Go converter](https://github.com/benhoyt/pas2go) that I wrote, as well as the [tcell](https://github.com/gdamore/tcell) terminal library for graphics.

To run it: install Go, clone the repo, type `go build`, and then run `./zztgo`. If you want to make it look a bit more authentic, you should install an [IBM EGA font](https://int10h.org/oldschool-pc-fonts/fontlist/#ibmega) and adjust the line spacing to zero. On macOS you can use [this Terminal settings file](https://github.com/benhoyt/zztgo/blob/master/zzt.terminal).

[**Read the full story here**.](https://benhoyt.com/writings/zzt-in-go/)
