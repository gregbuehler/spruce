# Spruce - Stored Value as a Service  [![Build Status](https://travis-ci.org/gregbuehler/spruce.svg?branch=master)](https://travis-ci.org/gregbuehler/spruce)

## The Service Model

Stored-value as a service consists of three core topics: who owns the value, how the value is organized into products, and transfer of value.

Ownership is organized as a `Platform` which has the direct relationship with the operator of the service and is performing the integration (e.g. a holding corporate entity). A `Platform` has a collection of `CustomerGroup` that logically organize the various holdings of the integrator (e.g. a brand owned by a corporate entity). A `CustomerGroup` has one or more `Account` against which funds are stored and orders are placed.

End users may purchase individual products which apply to the balance of the account by placing an `Order`. The action taken when an order is created is to email a message to a recipient.


## Primary Functionality

- `Accounts`

- `Catalogs`

- `Customers`

- `Funding`

- `Orders`

## Additional Functionality

- `Status`


## Future Ideas for improvement

- Use JWT in favor of HTTP Basic Auth
- Include monitoring/telemetry


## Distribution and runtime

Ultimately Spruce runs within one of Google's [distroless runtime containers](https://github.com/GoogleContainerTools/distroless).