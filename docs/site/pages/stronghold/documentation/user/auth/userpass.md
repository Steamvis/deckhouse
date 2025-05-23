---
title: "Userpass method"
permalink: en/stronghold/documentation/user/auth/userpass.html
lang: en
description: >-
  The "userpass" auth method allows users to authenticate with Stronghold using a
  username and password.
---

## Userpass auth method

The `userpass` auth method allows users to authenticate with Stronghold using
a username and password combination.

The username/password combinations are configured directly to the auth
method using the `users/` path. This method cannot read usernames and
passwords from an external source.

The method lowercases all submitted usernames, e.g. `Mary` and `mary` are the
same entry.

## Authentication

### Via the CLI

```shell-session
$ d8 stronghold login -method=userpass \
    username=mitchellh \
    password=foo
```

### Via the API

```shell-session
$ curl \
    --request POST \
    --data '{"password": "foo"}' \
    http://127.0.0.1:8200/v1/auth/userpass/login/mitchellh
```

The response will contain the token at `auth.client_token`:

```json
{
  "lease_id": "",
  "renewable": false,
  "lease_duration": 0,
  "data": null,
  "auth": {
    "client_token": "c4f280f6-fdb2-18eb-89d3-589e2e834cdb",
    "policies": ["admins"],
    "metadata": {
      "username": "mitchellh"
    },
    "lease_duration": 0,
    "renewable": false
  }
}
```

## Configuration

Auth methods must be configured in advance before users or machines can
authenticate. These steps are usually completed by an operator or configuration
management tool.

1. Enable the userpass auth method:

   ```shell-session
   d8 stronghold auth enable userpass
   ```

   This enables the userpass auth method at `auth/userpass`. To enable it at a different path, use the `-path` flag:

   ```shell-session
   d8 stronghold auth enable -path=<path> userpass
   ```

1. Configure it with users that are allowed to authenticate:

   ```shell-session
   $ d8 stronghold write auth/<userpass:path>/users/mitchellh \
       password=foo \
       policies=admins
   ```

   This creates a new user "mitchellh" with the password "foo" that will be
   associated with the "admins" policy. This is the only configuration
   necessary.

## User lockout

If a user provides bad credentials several times in quick succession,
Stronghold will stop trying to validate their credentials for a while, instead returning immediately
with a permission denied error. We call this behavior "user lockout". The time for which
a user will be locked out is called “lockout duration”. The user will be able to login after the lockout
duration has passed. The number of failed login attempts after which the user is locked out is called
“lockout threshold”. The lockout threshold counter is reset to zero after a few minutes without login attempts,
or upon a successful login attempt. The duration after which the counter will be reset to zero
after no login attempts is called "lockout counter reset". This can defeat both automated and targeted requests
i.e, user-based password guessing attacks as well as automated attacks.

The user lockout feature is enabled by default. The default values for "lockout threshold" is 5 attempts,
"lockout duration" is 15 minutes, "lockout counter reset" is 15 minutes.

The user lockout feature can be disabled as follows:

- It can be disabled globally using environment variable `VAULT_DISABLE_USER_LOCKOUT`.
- It can be disabled for all supported auth methods (ldap, userpass and approle) or a specific supported auth method using the `disable_lockout`
  parameter within `user_lockout` stanza in configuration file.
  Please see [user lockout configuration](/docs/configuration/user-lockout#user_lockout-stanza) for more details.
- It can be disabled for a specific auth mount using "auth tune". Please see [auth tune command](/docs/commands/auth/tune)
  or [auth tune api](/api-docs/system/auth#tune-auth-method) for more details.

{% alert level="warning" %}

**NOTE**: This feature is only supported by the userpass, ldap, and approle auth methods.

{% endalert %}

## API

The Userpass auth method has a full HTTP API. Please see the [Userpass auth
method API](/api-docs/auth/userpass) for more details.
