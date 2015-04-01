Collection+JSON Consumer Architecture
=====================================
The Collection+JSON, hereafter C+J, consumer should act like a web crawler,
allowing the client code to traverse between multiple C+J documents, retrieving
items, filling out and submitting templates, and traversing links.

Basic C+J Functionality
-----------------------
###Templates & Query Templates
This library should allow setting the values of templates and query templates by
individual "Set" methods on the templates, in addition to passing in a struct.
The struct can either implement an interface (TemplateMarshaler) or reflection
can used to inspect the struct, follow any struct tags given, and do a best
attempt to match the data of the template to the properties of the struct.

```go
myStruct := MyQueryStruct{Foo: "Bar"}
query := collection.Query("search foo")
query.Marshal(myStruct)
query.Set("Hello", "World")
responseCollection, err := query.Submit()
if err != nil {
    // handle error
}
```

For templates, one should be able to unmarshal the template into the provided
struct given the struct implements an interface (TemplateUnmarshaler) or via
reflection using struct tags or best guessing.

```go
template := collection.Template()
myStruct := MyStruct{}
err := template.Unmarshal(&myStruct)
if err != nil {
    // handle error
}
myStruct.Hello = "World"
err = template.Marshal(myStruct)
if err != nil {
    // handle error
}
// Potentially, err can non-nil if the status code is non-200
status, err := template.Submit()
if err != nil {
    // handle error
}
```

###Retrieving Items
Clients should be able to retrieve individual items from the collection using
either key/value pairs from the data, the href of the item, or rel/href pairs
from the links attached to an item. The library will provide an Option type for
these three types of queries and for ANDing and ORing them together. This allows
for a level of composability that ensure flexibility while keeping complexity at
a minimum. This does come with a slight increase in verbosity, but allows code
to remain clear as to the intent of a specific retrieval query.

###Finding & Traversing Links
Clients should be able to retrieve links via their rel or name attribute. They
should then be able to traverse those links to another C+J document. Links
returned from the collection should implement a link interface, allowing
several different implementation backends to be used.

###C+J Error
When the C+J document being parsed contains an error a special error type should
be returned that contains the information provided in response. Since a C+J
error document can also contain other types of data, the collection returned may
be populated as well.

Interfaces
----------
To support a wide range of uses, this library should provide several interfaces
that can be implemented by the library itself or client code. These interfaces
include marshalling and unmarshalling and link traversing interfaces. Along with
the interfaces, reflection and struct tags should be implemented to allow
automatic unmarshalling of C+J structures.

For instance, in the case of attempting to unmarshal into a given struct, by
default if there have been no marshaling interfaces implemented and no struct
tags assigned, this library will attempt to first assign the keys in the data
array to the struct properties, if those keys are either a string, []byte, or
number. If the struct properties implement the link interface and the property
name matches either the rel or name of a link in the links array, then the link
will be assigned to that property.

Struct tags can be used to override this behavior via either simply renaming the
property, as in Foo to foo, or by explicitly setting the type, as in cj:",link".

Overall this allows clients to use their own structs to consume C+J which
enables this library to be used as a base level of composition. This also allows
for the definition of conversion to be made with static code allowing static
analysis tools to be created that can help ensure the intent matches the
implementation. Static analysis tools give an additional ability to write a
client based on a specific profile and ensure that before runtime the given code
matches the profile as provided to said tools.

Creating a new collection
-------------------------
Initial collection creation would usually be done via URL. However, there are
feasible situations when the bytes could be created from some other method or
retrieved via another method, e.g. file on disk, parsing from the body of a
request handled, etc.. For this reason, while this initial implementation takes
a direct href string, in the future there should be a general "fetcher"
implemented and self referential options should be used to configure the initial
collection to support these other methods of instantiating and browsing a
collection.

These are not being added in the initial version because it adds a large amount
of verbosity and API space without a clear indication of actual usefulness.
Additionally, adding the extra functionality for creating a Collection with self
referential options could easily be added on without needing a breaking API
change, although the resulting API would be slightly strange in situations where
an href string is not being used.

Immutability
------------
This library is designed with immutability, meaning that when a fetch is called
a new collection is always returned instead of modifying the caller's
collection. This allows following multiple paths from the same collection. It
also helps prevent the accidental overwriting of a collection.
