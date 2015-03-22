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
In this example we will build out the ALPS profile documented in the
[app.js](https://github.com/alps-io/alps-contacts/blob/master/app.js) of the
[ALPS contacts](https://github.com/alps-io/alps-contacts) example.

We'll start by building out the "search" descriptor. The search descriptor is an
example of a descriptor with a state transition, which represents a state
transition inside of a server implementation of the profile, e.g. a form which
can be submitted.

Since we need to build from the bottom up we start with the "name" descriptor
inside the "search" descriptor.

```go
docOpt := NewDoc(nil, "text", "input for searching")
// Semantic is the default but we'll set it manually as well.
typeOpt := NewType(alps.Semantic)
nameDescriptor, err := NewDescriptor("name", nil, typeOpt, docOpt)
if err != nil {
    panic(err)
}
```

The nameDescriptor we've created is an Option which we can use to either create
another descriptor option by passing it into another call to NewDescriptor or we
can create a profile by passing it into a call to NewProfile.

The "name" descriptor is embedded in the "search" descriptor, so we'll create
that one next.

```go
// Since we've already added the previous doc to a descriptor we can reuse the
// variable
docOpt = NewDoc(nil, "text", "search for contacts in the list")
typeOpt = NewType(alps.Safe)
rtOpt := NewRt("contacts")
searchDescriptor , err := NewDescriptor("search", nil, docOpt, typeOpt, rtOpt, nameDescriptor)
if err != nil {
    panic(err)
}
```

Next we need to create the "contacts" descriptor which has 5 embedded
descriptors. We'll create each of those first, then attach them to the
"contacts" descriptor.

```go
typeOpt = NewType(alps.Safe)
docOpt = NewDoc(nil, "text", "link to contact")
itemDescriptor, err := NewDescriptor("item", nil, typeOpt, docOpt)
if err != nil {
    panic(err)
}

typeOpt = NewType(alps.Semantic)

// We use *url.URL to create hrefs, so we'll parse and then pass it in
href, err := url.Parse("http://schema.org/givenName")
if err != nil {
    panic(err)
}
givenDescriptor, err := NewDescriptor("givenName", href, typeOpt)
if err != nil {
    panic(err)
}

// Since each href has the same base, we can safely reuse the href and just
// update the path, alternatively one could reparse the whole url.
href.Path = "/familyName"

// We can also reuse the typeOpt as many times as we wish, each Option uses the
// value and encodes it in the various elements making it safe to reuse.
familyDescriptor, err := NewDescriptor("familyName", href, typeOpt)
if err != nil {
    panic(err)
}

href.Path = "/email"
emailDescriptor, err := NewDescriptor("email", href, typeOpt)
if err != nil {
    panic(err)
}

href.Path = "/telephone"
telephoneDescriptor, err := NewDescriptor("telephone", href, typeOpt)
if err != nil {
    panic(err)
}

// Finally we assemble the "contacts" descriptor
// We should have added each descriptor to a slice and passed in the slice using
// a variadic parameter
docOpt := NewDoc(nil, "text", "contact item")
contactsDescriptor, err := NewDescriptor("contacts", nil, docOpt, typeOpt, itemDescriptor, givenDescriptor, familyDescriptor, emailDescriptor, telephoneDescriptor)
if err != nil {
    panic(err)
}
```

Now that we have our two descriptors we can create a new profile from them, and
marshal the profile into JSON.

```go
docOpt = NewDoc(nil, "text", "List of contacts w/ search")
contactsProfile, err := NewProfile(docOpt, searchDescriptor, contactsDescriptor)
if err != nil {
    panic(err)
}
result, err := json.Marshal(contactsProfile)
if err != nil {
    panic(err)
}
var out bytes.Buffer
json.Indent(&out, result, "", "\t")
out.WriteTo(os.Stdout)
```

[This gist is a full
version](https://gist.github.com/skriptble/6bb0f97441aa879f94ee) of the above
example. It is fully runnable, so you can just copy the code into your text
editor, and issue a "go run main.go" command and see the json output. You will
need to "go get github.com/skriptble/hyper/profiles/alps/..." first.
