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

##Struct & Struct Tag Construction
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
