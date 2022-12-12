# Sentinel - Secure policy based workflow authorization system

Solving the problem where you have an application that has an endpoint which
needs have a second pair of eyes approve the request before the command is
actually allowed to run. This is common in regulated environments where it
is best to ensure that no single corrupt person can perform a destructive
operation.

Sentinel is a policy based authorization system that allows you to have a
central location where policies are stored and enforced to ensure a credible
audit trail which is run outside the scope of the service.

> NOTE: This service is a proof-of-concept and is not production ready.
> It also assumes the service gating access is trusted as enforcement is only
> working if the service is trusted. This is a problem that could be solves by
> having it be the gatekeeper to credentials if the downstream service requires
> authentication.

# Usage

Go service handler endpoint.

```go

```
