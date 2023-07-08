# Using oidc for login in gin

This example shows how to include an OIDC Identity Provider in your gin webserver to handle logins.

The whole login flow is handled by the oidc package and you only have to start the session in gin when the client successfully logged in on the IP.

This was successfully tested with Keycloak v21.1.1.

I highly recommend you to read about OIDC and SAML if you haven't already, or this will not make much sense to you.
