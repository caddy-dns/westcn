West.cn DNS module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [West.cn DNS](https://west.cn/).

## Caddy module name

```
dns.providers.westcn
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "westcn",
				"username": "WESTCN_USERNAME",
				"api_password": "WESTCN_API_PASSWORD"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns westcn {
		username {env.WESTCN_USERNAME}
		api_password {env.WESTCN_API_PASSWORD}
	}
}
```

```
# one site
tls {
  dns westcn {
    username {env.WESTCN_USERNAME}
    api_password {env.WESTCN_API_PASSWORD}
  }
}
```