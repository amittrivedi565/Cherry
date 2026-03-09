## Cherry
Cherry is a query builder & executer for SQL Dialect, It uses DASL (Data Agnostic Service Language) to query and execute operations.

Cherry takes JSON (DASL) as an input, in which all the required information is contained, it returns a JSON object with required data.

It is a primitive block, independent of any service, created for the sole purpose of data query for our application, services relies upon DASL for communication.

DASL is an ecosystem language used by our services to communicate with each other, every service has its own DASL contract upon which it relies, DASL provides an abstraction to the application developer to not rely upon the domain knowledge.

## Architechure

```arch
Input (DASL) -> Query Router -> Query Builder -> Query Executer -> Output
```

## Upcoming

- Ensure service has database dialect.
- Provides complete abstraction over operations.
- Make DASL more versatile and robust for communication.

