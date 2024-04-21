# net/http login with cookies

This small example shows how to handle user logins with cookies.  
This is useful if basic authentication is not enough but using jwt or oauth2 is too much.

This example uses an `sync.Map` for "storing" logins because a inbuilt map (e.g. `map[string]string`) is not save for concurrent use from multiple goroutines.  
If the application will run behind load balancing with many instances running at once, then you should use e.g. Redis to save and manage these logins.

The `GenerateRandomKey()` function should be more random in a production environment, consider using google's `uuid4` package or `crypto/rand`.  
Same with the `IsUserLoginCorrect` function, this should query the database and verify the user login that way.  
The user password also has to be hashed with e.g. `bcrypt` when comparing it with the database stored one.

To stay logged in you should set the cookie with a max age again and return it to the client, or else the client will be logged out X-hours after login.  
To log the client out manually just set the MaxAge of the cookie to -1, this will tell the client to delete it.
