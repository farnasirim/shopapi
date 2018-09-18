# Documentation

## Overal review
The language used for implementation is `go`.

The project layout is based on [this](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) blog post. "Standard Package Layout" segments the projects to different parts based on third party dependencies. No two parts will have cross references to each other and they will communicate through abstract interfaces that the root package provides: "Program to the interface, not the implementation".

The actual concrete instances will be supplied to the users of the interfaces at runtime by the `main` function (if executable), or by the client code (if it's library code).

This has many important and useful implications:
 - Each third party package (e.g. `data/mongodb`) which "implements" an interface (e.g. `UserInfoService`), can be switched for another third party implementation without the other packages even noticing: They had been only using the interface. They could also 
 - One can add support for other API protocols such as `GRPC` over `protobuf`.
 - Unit testing the individual packages can be done minimally and there won't be need of deep initializations.
 - Each component will by definition be "wrappable" for applications such as tracing or stats.

And some bad ones:
 - The API layer being pushed back from touching the details in the data access/business logic layer (aside from it's obvious plus side: no redundancy on api endpoints/resolvers) could result in very inefficient database calls and worse than that, a lot of round trips to the DB. This issue can later on be approached by introducing a query planner/query builder service between the two layers to let the (in particular graphql) callers decide which path they want to go in a certain operation.

## Important design decisions
- [This graphql schema](../api/graphql/schema.graphql) is used for working with graphql throughout this project. The more popular [graphql-go](https://github.com/graph-gophers/graphql-go) library lacks static type checkings (`interface{}` is used instead) and doesn't need an actual schema to validate the code against which in my opinion were two important drawbacks, making me not want to go down the other path.

## Next steps
- How do we do discounts and special offers? Should single item extras be handled inside the line item (buy 2 cookies to get 1 free)? How about the multi item extras (buy cookie + milk to get something free)? 
- Add pagination to plural api nouns
- Use `dep` for dependency management instead of bare `go get` calls and a `Makefile`