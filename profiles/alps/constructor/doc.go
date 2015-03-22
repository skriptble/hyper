/*
Package constructor implements the structures and methods to create ALPS
profiles. It currently implements version 01 (http://tools.ietf.org/html/draft-amundsen-richardson-foster-alps-01)
of the ALPS spec.

Profiles are built starting at the bottom and building up from there. Start
by creating a descriptor. There are several configuration options for a
descriptor. For most profiles we won't actually need any options outside of
NewName, NewType, NewDoc, and possibly NewRt.

	docOpt := NewDoc(nil, "text", "input for searching")
	descriptorOpt, err := NewDescriptor("contacts", nil, docOpt)
	if err != nil {
		panic(err)
	}

Once the descriptors are built up, create a Profile by calling NewProfile.
A Doc can also be set on the Profile.

	docOpt = NewDoc(nil, "text", "List of contacts w/ search")
	prof, err := NewProfile(descriptorOpt, docOpt)
	if err != nil {
		panic(err)
	}
	result, err := json.Marshal(prof)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(os.Stdout, result)

For a full example see the README.md (https://github.com/skriptble/hyper/blob/master/profiles/alps/constructor/README.md)
*/
package constructor
