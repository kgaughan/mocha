# mocha: a web-based feed reader

Let's try and reboot this again.

Mocha is intended to be a feed reader. It's also an experiment in UI design,
giving users a way to quickly triage new and unread posts on their queue. Think
of it as 'Tinder for your feed'.

I'm rebooting the project for the umpteenth time, this time using
[Go](https://golang.org/) and [gofeed](https://github.com/mmcdole/gofeed).

It's a revival of an old college project, when I wrote a simple mail client
style feed reader in Visual Basic, which I called 'Mocha' at the time. That
code is long gone, but I've kept meaning to do something similar again, though
I never really persisted with the projects beyond the initial stab.

## Build prerequisites

Go, obviously.

Also, you'll need [pkger][] for rebuilding any static assets. Install it with:

    go get github.com/markbates/pkger/cmd/pkger

[pkger]: https://github.com/markbates/pkger
