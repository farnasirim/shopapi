# Documentation

## Overal review
The language used for implementation is `go`.

The project layout is based on [this](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) blog post. "Standard Package Layout" segments the projects to different parts based on third party dependencies. No two parts will have cross references to each other and they will communicate through abstract interfaces that the root package provides: "Program to the interface, not the implementation".

The actual concrete instances will be supplied to the users of the interfaces at runtime by the `main` function (if executable), or by the client code (if it's library code).

This has many important and useful implications:
 - Each third party package (e.g. `data/mongodb`) which "implements" an interface (e.g. `UserInfoService`), can be switched for another third party implementation without the other packages even noticing: They had been only using the interface.
 - Unit testing the individual packages can be done minimally and there won't be need of deep initializations.
 - Each component will by definition be "wrappable" for applications such as tracing or stats.

## Important design decisions
[The graphql schema](../api/graphql/schema.graphql)
Todo: Why this graphql library

## Next steps
Product unrelated things handled in line item
