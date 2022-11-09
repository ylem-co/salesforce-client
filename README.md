Salesforce REST client
===========

How to use?
------

First of initiate the client with stateful oauth credentials. E.g., in init

```go
func init() {
	salesforceclient.Initiate(salesforceclient.Config{
		ClientID:     "<your client id>",
		ClientSecret: "<your client secret>",
		RedirectUrl:  "<guess what: redirect url>",
		Scopes:       []string{"api", "chatter_api"},
	})
}
```

After the grant & token exchange (it's pretty straight forward), create a client and... use it.

```go
sf, _ := salesforceclient.CreateInstance(ctx, token)
_ = sf.CreateCase(request)
```
