# RememberMe [![Build Status](https://travis-ci.org/janekolszak/rememberme.svg?branch=master)](https://travis-ci.org/janekolszak/rememberme)
Go implementation of persistent login cookies. The implementation idea is described [here](https://paragonie.com/blog/2015/04/secure-authentication-php-with-long-term-persistence#title.2)

## Usage
### Initialization
```go
// Create the store
dbCookieStore, err := sqlstore.New("sqlite3", "/tmp/database.db3")
if err != nil {
	panic(err)
}

// Create Rememberme
rememberme = &rememberme.New{
	Store:  dbCookieStore,
	MaxAge: time.Second * 30,
}
```

### Sign-in handler
```go
func(w http.ResponseWriter, r *http.Request) {

	selector, userid, err := cookieProvider.Check(r)
	if err == nil {
		// Authenticated with a RememberMe Cookie
		err = cookieProvider.UpdateCookie(w, r, selector, userid)
		if err != nil {
			panic(err)
		}
	} else {
			// Can't authenticate with "Remember Me" cookie,
			// so try with another provider.

			// userid = Authenticated User ID

			// Save the RememberMe cookie
			err = cookieProvider.SetCookie(w, r, userid)
			if err != nil {
				panic(err)
			}
	}

	// User authenticated, cookie saved
	// ...
}
