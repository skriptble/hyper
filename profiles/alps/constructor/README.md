ALPS Constructor
================
[![GoDoc](https://godoc.org/github.com/skriptble/hyper/profiles/alps/constructor?status.svg)](https://godoc.org/github.com/skriptble/hyper/profiles/alps/constructor)

This package implements a constructor for ALPS profiles. It currently implements
[version 01 of the ALPS
Spec](http://tools.ietf.org/html/draft-amundsen-richardson-foster-alps-01).

Currently this package provides a rather low level API. It uses interfaces and
functions to handle the build up of a descriptors and profiles. Both profiles
and descriptors are constructed from the bottom up. So first all of the options,
such as doc, ext, type, name, etc.. are created and then passed in during the
construction of the profile or descriptor.

There are plans to build a better API implementation, perhaps with code
generation. This way it won't feel so backwards to build up a profile.

Currently there are few optimizations involving the creation of the JSON. This
is on purpose as the ideal situation for using this library is in a tool to
generate the profiles to disk as json files. These files would then be uploaded
to a CDN or static file server. A server could be built to render directly, but
if done it is advisable that the server builds all the profiles on startup and
cache them for serving.

ALPS profiles are not meant to be dynamically changed. Even though many clients
will be able to cope with the changes, changing the profiles means there has
been a change in the application semantics, indicating a version tick. This
should be handled in a managed fashion. For more information and discussion on
the topic see the [ALPS Google
Group](https://groups.google.com/forum/?fromgroups=#!forum/alps-io)

Example
-------
