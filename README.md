# Go Basic Auth

A fairly secure app with basic auth following some best practices to improve security. Prompts you for a password when
accessing the protected route. Uses a custom middleware to implement basic auth guard.

### Practices to improve security

1. Hash passwords and usernames to make their lengths obscure from an attacker.
2. Use crypto.subtleTimeCompare to check the entirety of the hashes before informing us of a match or mismatch.
3. Served over HTTPS

### Routes

- [Unprotected route - `/`]("https://go-basic-auth.herokuapp.com/")
- [Protected route - `/dashboard`]("https://go-basic-auth.herokuapp.com/dashboard")

