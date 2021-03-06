# RememberMe [![Build Status](https://travis-ci.org/janekolszak/rememberme.svg?branch=master)](https://travis-ci.org/janekolszak/rememberme) [![Code Climate](https://codeclimate.com/github/janekolszak/rememberme/badges/gpa.svg)](https://codeclimate.com/github/janekolszak/rememberme) [![GoDoc](https://godoc.org/github.com/janekolszak/rememberme?status.svg)](https://godoc.org/github.com/janekolszak/rememberme) [![Join the chat at https://gitter.im/go-rememberme/Lobby](https://badges.gitter.im/go-rememberme/Lobby.svg)](https://gitter.im/go-rememberme/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
Go implementation of persistent login cookies. The implementation idea is described [here](https://paragonie.com/blog/2015/04/secure-authentication-php-with-long-term-persistence#title.2)

The library uses [gorilla sessions](http://www.gorillatoolkit.org/pkg/sessions) for saving cookies on user site. Part of the persistant login information is stored on a server side in a dedicated Store. There are two store implementations available out of the box:
- SQL database store
- RethinkDB store

## Installation
```
go get -u github.com/janekolszak/rememberme
```

## Usage
### Initialization
```go
// Create the store
dbCookieStore, err := sqlstore.New("sqlite3", "/tmp/database.db3")
if err != nil {
	panic(err)
}

// Create Rememberme
auth = &rememberme.New{
	Store:  dbCookieStore,
	MaxAge: time.Second * 30,
}
```

### Sign-in handler
```go
func(w http.ResponseWriter, r *http.Request) {

	selector, userid, err := auth.Check(r)
	if err == nil {
		// Authenticated with a RememberMe Cookie
		err = auth.UpdateCookie(w, r, selector, userid)
		if err != nil {
			panic(err)
		}
	} else {
			// Can't authenticate with "Remember Me" cookie,
			// so try with another provider.

			// userid = Authenticated User ID

			// Save the RememberMe cookie
			err = auth.SetCookie(w, r, userid)
			if err != nil {
				panic(err)
			}
	}

	// User authenticated, cookie saved
	// ...
}
