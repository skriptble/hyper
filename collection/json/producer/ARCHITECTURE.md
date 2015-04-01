Collection+JSON Producer Architecture
=====================================
The Collection+JSON, hereafter C+J, producer should act similar to the
encoding/json marshaler. It should allow constructing "structless" C+J documents
via bottom-up self-referential configuration functions and struct configurations
using a user defined struct. The struct version should allow reflection, struct
tags, and interface methods of configuration.

Constructing C+J Documents
--------------------------
###Option Construction
Each element inside of a C+J document can be created via a constructor that
takes the required parameters and potentially other options, and returns either
an option or an implementation of an interface. This allows for a bottom up
construction of C+J documents.

###Struct & Struct Tag Construction
Each element of a C+J document has an interface that can be implemented. Any
struct can be marshaled into C+J. If it has no tags and doesn't implement the
collection/json Marshaler interface then reflection will be used to marshal the
properties of the struct into a C+J document. Only properties that are either a
string, a number (int, uint, float, etc...), or implement the datum interface
will be marshaled into datum.

Fields on a struct that implement the links or link interface will be marshalled
into the links element. Fields on a struct that implement the queries or query
interface will be marshalled into the query element. If multiple implementations
of either link or query interfaces exist on a struct, they will be aggregated
together into the links or query elements, respectively.

Fields on a struct that implement the template or error interface will be
marshalled into the template or error elements, respectively. If more than one
implementation of each of these interfaces is attached only the last one will
be used.

####Struct Tags
There are struct tags for each type of element in C+J: link, item, query, datum,
data, template, and error. Struct tags for link and datum can reasonably be used
by many implementations instead of implementing the associated interface. This
will benefit static analysis tools. The remainder of the struct tags can only be
used to avoid ambiguity in the case where a particular type implements more than
one interface.

*Link*
The link struct tag is used to indicate a property is a link. This tag can be
used on a struct that is an item of a collection or a collection itself. It is
formatted as follows:

```go
type Foo struct {
    Bar string `cj:"link,profile"`
    Baz url.URL `cj:"link,next"`
}
```

The first parameter of the tag is the word "link" followed by the rel of the
link. If additional properties are required (e.g. name, render, or prompt) then
the link interface must be implemented instead. The type of the struct property
can either be a string, url.URL, a pointer to a string, a pointer to a url.URL,
a type that implements the link interface, or a slice of a type that implements
the link interface.

*Datum*
The datum struct tag is used to indicate a property is a datum. This tag can be
used on a struct that is a query, template, or item. It is formatted as follows:

```go
type Foo struct {
    Bar string `cj:"datum,A Bar,bar"`
}
```

The first parameter of the tag is the word "datum" followed by the propmpt for
the datum, followed by the alternative Name. The type of the property can be
anything that can be marshaled into JSON or a type that implements the datum
interface.
