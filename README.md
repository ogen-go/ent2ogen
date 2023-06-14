# ent2ogen

Quite often there's a need to expose some database entities through service api.\
If you are using [ent](https://github.com/ent/ent) as a database abstraction layer, and [ogen](https://github.com/ogen-go/ogen) for an API, then you might need to do type conversion between ent and ogen types, like this:

```go
func (s *Server) GetUser(ctx context.Context, params openapi.GetUserParams) (openapi.User, error) {
	u, err := s.db.Users.Get(params.UserID)
	if err != nil {
		return openapi.User{}, fmt.Errorf("query user: %w", err)
	}

	return openapi.User{
		ID:       u.ID,
		Username: u.Username,
		Age:      u.Age,
		// and so on...
	}
}
```

Writing such conversion by hand is annoying and error-prone (especially when the types are big with deep nesting).\
Also these conversions may become out of sync over time because of database or api schema updates.\
Ent2ogen solves this problem by generating mapping functions automatically.

## How to use

1. Create openapi schema
2. Create ent schema
3. Create ```entc.go``` file ([sample](example/ent/entc.go))
4. Create ```generate.go``` file ([sample](example/ent/generate.go))
5. Use following ent schema annotations:

* ```ent2ogen.BindTo("")``` - generate mapping function to specified openapi schema component.
* ```ent2ogen.Bind()``` - similar to BindTo but uses ent schema name by default.

6. Run ```go generate```
