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
### Graphql
I've used [This graphql library](../api/graphql/schema.graphql) to parse graphql calls throughout this project. The more popular [graphql-go](https://github.com/graph-gophers/graphql-go) library lacks static types (`interface{}` is used instead) and doesn't need an actual schema to validate the code against, which in my opinion were two important drawbacks, making me not want to go down the other path.

### Data
I've decided to go with `mongodb` as the datastore. A schemaless database will save a lot of time in development of the shops api as we are more likely to see different purchase scenarios and new entities with changing relations between them.

Also I think `mongodb` will go much better with `graphql` than any relational dbms would in terms of how different the languages (api language and db query language) look, since the amount of code that we should write and maintian is roughly proportional to this difference. For example deep queries is a nice feature to have in the data source when a very similar thing can happen in graphql.

For anything related to data as sensitive as payments though I would still stick to a more reliable dbms in terms of Consistency.

I have decided to store all the data in a single shops collection, since accessing/creating products, orders, and line items outside of the context of a shop would be meaningless. If that were to be the case though, depending on the usage I would either duplicate all the products to another products collection too, or just created a products collection without duplicating everything under the corresponding shop item.

Also for the line items, I have decided to again keep them under the orders (which are themselves kept inside shops) since line items would be meaningless without an order. Inside the line items I will duplicate the "important" fields of product (which is pretty much everything now: name and dollar value) but will also keep a reference to the actual Product. This has implications which may or may not be wanted:
- "Freezing" the state of a product by adding it to a shopping cart (well, it's order right now, but shopping cart can make it worse) and not finalizing the cart, will keep an old "version" of a product in the system.
- On the other hand when querying the orders in the past, one would probably like expect the names and prices to be as they were at the time of the order. At worst the client can look up the product reference from each line item to find the current price/name/etc.

I believe we can come up with different solutions for issues similar to the first one and therefore find the first solution (not only keep the reference but also duplicate important data for calculations and business logic) better:
1. It makes more sense logically
2. When querying for anything that needs the whole order or an aggregation of it's products this will saves us a round trip to the db

Anything in the data modeling (keep prices precalculated or do it on the fly, removing items from orders, indexes, etc.) can be discussed but I've decided to keep this document in a level higher than that.


## Next steps
- How do we do discounts and special offers? Should single item extras be handled inside the line item (buy 2 cookies to get 1 free)? How about the multi item extras (buy cookie + milk to get something free)? 
- Add pagination to plural api nouns
- Use `dep` for dependency management instead of bare `go get` calls and a `Makefile`
- Can be a little more sophisticated in the modelToGrpahQL (e.g. `/api/graphql/shop_resolver:shopModelToGraphQL`) adapters. There is no need to go back to the data service for the name when you could cache it alongside the ID. This depends on the field and how much we are keep the in memory objects consistent with the DB though, but for something like Name this would be fine. On the other hand, this could interfere with DB fetching strategies. For example even the Name could have been provided through a lazy strategy and a eventhough the client thinks that prefetching the Name property would be a good idea performance wise, DB's autonomy in this strategy is now violated. Maybe we should handoff the query optimization completely to the data service so that it could do all sorts of black magic (e.g. caching everything very close by - for example in process memory, making the data service stateful! Only for a a number of back and forth calls = about a second) in these scenarios.
