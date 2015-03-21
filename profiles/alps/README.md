Application-Level Profile Semantics (ALPS)
==========================================
Alps is (as stated in the ALPS specification):

> a data format for defining simple descriptions of application-level semantics.

ALPS documents can be used to represent the general semantics of an application
in a media type agnostic manner. This means one can use an ALPS profile to write
the high level semantics of an application and apply it to a specific media type
such as HTML, Collection+JSON, UBER, Siren, etc.

While ALPS profiles do describe the semantics of an application, they do not
fully describe all the semantics of the application, as such a specific media
type will provide additional semantics as necessary. This allows ALPS to be used
across multiple media types while not enforcing a specific server or client
implementation.

ALPS profile documents primarily describe the possible semantics of an
application. Server implementations may choose to add additional semantics or
only use a subset of the semantics provided by a profile.

This library implements the ALPS specific and provides a constructor for
creating ALPS documents and a consumer for consuming ALPS documents.

Constructor
-----------
The constructor should be used to create the profiles which should then be
cached. ALPS profiles are not meant to be dynamic, so the constructor can be
used to create a cli tool that renders JSON or XML ALPS profiles to disk which
can then be deployed onto a CDN or static asset server.
