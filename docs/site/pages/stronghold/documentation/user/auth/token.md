---
title: "Token"
permalink: en/stronghold/documentation/user/auth/token.html
lang: en
description: The token store auth method is used to authenticate using tokens.
---

## Token auth method

The `token` auth method is built-in and automatically available at `/auth/token`. It
allows users to authenticate using a token, as well to create new tokens, revoke
secrets by token, and more.

When any other auth method returns an identity, Stronghold core invokes the
token method to create a new unique token for that identity.

The token store can also be used to bypass any other auth method:
you can create tokens directly, as well as perform a variety of other
operations on tokens such as renewal and revocation.

Please see the [token concepts](/docs/concepts/tokens) page dedicated
to tokens.

## Authentication

### Via the CLI

```shell-session
d8 stronghold login token=<token>
```

### Via the API

The token is set directly as a header for the HTTP API. The header should be
either `X-Vault-Token: <token>` or `Authorization: Bearer <token>`.

## API

The Token auth method has a full HTTP API. Please see the
[Token auth method API](/api-docs/auth/token) for more
details.
