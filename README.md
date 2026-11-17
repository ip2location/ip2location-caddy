# IP2Location Caddy Middleware

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Add IP geolocation and network intelligence to requests handled by [Caddy](https://caddyserver.com/) using either a local **IP2Location BIN database** or the **IP2Location.io API**. The middleware looks up the visitor's IP address and makes the results available to downstream handlers as `X-IP2Location-*` request headers.

## Features

- **Two lookup modes:** local BIN database or remote IP2Location.io API.
- **Request enrichment:** add location, network, ISP, and ASN data as HTTP request headers.
- **Configurable header prefix:** `X-IP2Location` by default.
- **Caddy integration:** use with downstream handlers such as `php_fastcgi` or `reverse_proxy`.
- **IPv4 and IPv6:** subject to the selected database and API data coverage.

## How It Works

```text
Client request → Caddy/IP2Location middleware → IP lookup (BIN or API)
               → Enriched request headers → Application / upstream
```

The middleware enriches the **request**, not the response. Your application reads the injected headers; a browser will not see them as response headers unless the application explicitly sends them.

## Requirements

- [Go](https://go.dev/doc/install) and [xcaddy](https://github.com/caddyserver/xcaddy) to build a Caddy binary containing the module.
- **Local mode:** a compatible IP2Location `.BIN` file. [IP2Location LITE](https://www.ip2location.com/database/lite) offers free downloadable databases; available fields depend on the BIN package.
- **Remote mode:** an [IP2Location.io API key](https://www.ip2location.io/sign-up). API fields and query limits depend on the subscription plan.
- A configured Caddy site and a downstream application that can read request headers.

## Installation

Build Caddy with the middleware:

```bash
xcaddy build --with github.com/ip2location/ip2location-caddy
```

Verify that Caddy recognizes the plugin:

```bash
./caddy list-modules | grep ip2location
```

## Configuration

Choose **one** lookup mode below. These examples forward enriched requests to PHP-FPM; adapt the downstream handler for your application.

### Option 1: Local BIN database

Download and extract an IP2Location BIN database, then configure its absolute path:

```caddyfile
{
    order ip2location before php_fastcgi
}

example.com {
    root * /var/www/html

    ip2location {
        mode local
        bin_path /path/to/IP2LOCATION.BIN
        header_prefix X-IP2Location
    }

    php_fastcgi unix//run/php/php8.4-fpm.sock
    file_server
}
```

Replace the BIN path and PHP-FPM socket with the ones in your environment. Local lookups use the installed database; refresh the BIN file as new data becomes available.

### Option 2: Remote IP2Location.io API

Set the API key in the environment used by Caddy:

```bash
export IP2LOCATION_API_KEY="YOUR_API_KEY"
```

Then use Caddy's environment placeholder in your Caddyfile:

```caddyfile
{
    order ip2location before php_fastcgi
}

example.com {
    root * /var/www/html

    ip2location {
        mode remote
        api_key {$IP2LOCATION_API_KEY}
        header_prefix X-IP2Location
    }

    php_fastcgi unix//run/php/php8.4-fpm.sock
    file_server
}
```

Keep the API key out of source control. The `api_key` setting is mandatory in remote mode; the module does **not** automatically read an environment variable unless it is referenced in the Caddyfile as shown.

### Configuration reference

| Directive | Description | Default |
|---|---|---|
| `mode` | `local` for BIN lookup or `remote` for API lookup. | `local` |
| `bin_path` | Absolute path to the BIN database; required in local mode. | None |
| `api_key` | IP2Location.io API key; required in remote mode. | None |
| `header_prefix` | Prefix for injected request headers. | `X-IP2Location` |

**Note:** The Caddyfile directive is `ip2location` (without spaces).

## Usage

Validate your configuration and start the custom Caddy binary:

```bash
./caddy validate --config /path/to/Caddyfile --adapter caddyfile
./caddy run --config /path/to/Caddyfile --adapter caddyfile
```

Run Caddy with the permissions required by your chosen listening ports, BIN file, and upstream services.

### Reading the headers in PHP

The middleware adds headers to the request before forwarding it to PHP. PHP exposes them through `$_SERVER` with the `HTTP_` prefix and underscores:

```php
<?php
$fields = [
	'COUNTRY_CODE',
	'COUNTRY_NAME',
	'REGION',
	'CITY',
	'ISP',
	'LATITUDE',
	'LONGITUDE',
	'DOMAIN',
	'ZIP_CODE',
	'TIME_ZONE',
	'NET_SPEED',
	'IDD_CODE',
	'AREA_CODE',
	'WEATHER_STATION_CODE',
	'WEATHER_STATION_NAME',
	'MCC',
	'MNC',
	'MOBILE_BRAND',
	'ELEVATION',
	'USAGE_TYPE',
	'ADDRESS_TYPE',
	'CATEGORY',
	'DISTRICT',
	'ASN',
	'AS',
	'AS_DOMAIN',
	'AS_USAGE_TYPE',
	'AS_CIDR',
];

foreach ($fields as $field) {
    $name = 'HTTP_X_IP2LOCATION_' . $field;
    if (isset($_SERVER[$name])) {
        echo htmlspecialchars($field . ': ' . $_SERVER[$name], ENT_QUOTES, 'UTF-8') . "<br>\n";
    }
}
```

This is a shortened version of the original repository's PHP test page. The same headers can be consumed by other upstream applications.

## Enrichment Headers

With the default prefix, the middleware sets the following request headers after a successful lookup. Both lookup modes use the same header names, but **populated values depend on the chosen BIN database or API subscription**.

| Header | What it represents |
|---|---|
| `X-IP2Location-Country-Code` | Two-letter country code |
| `X-IP2Location-Country-Name` | Country name |
| `X-IP2Location-Region` | Region or state |
| `X-IP2Location-City` | City |
| `X-IP2Location-Latitude` | Approximate latitude |
| `X-IP2Location-Longitude` | Approximate longitude |
| `X-IP2Location-ZIP-Code` | Postal or ZIP code |
| `X-IP2Location-Time-Zone` | Time zone |
| `X-IP2Location-ISP` | Internet service provider |
| `X-IP2Location-Domain` | Associated domain |
| `X-IP2Location-Net-Speed` | Network speed category |
| `X-IP2Location-IDD-Code` | International dialing code |
| `X-IP2Location-Area-Code` | Telephone area code |
| `X-IP2Location-Weather-Station-Code` | Nearby weather station code |
| `X-IP2Location-Weather-Station-Name` | Nearby weather station name |
| `X-IP2Location-MCC` | Mobile country code |
| `X-IP2Location-MNC` | Mobile network code |
| `X-IP2Location-Mobile-Brand` | Mobile network brand |
| `X-IP2Location-Elevation` | Approximate elevation |
| `X-IP2Location-Usage-Type` | IP usage classification |
| `X-IP2Location-Address-Type` | IP address type classification |
| `X-IP2Location-Category` | Associated category |
| `X-IP2Location-District` | District or county |
| `X-IP2Location-ASN` | Autonomous System Number |
| `X-IP2Location-AS` | Autonomous System name |
| `X-IP2Location-AS-Domain` | AS-associated domain |
| `X-IP2Location-AS-Usage-Type` | AS usage classification |
| `X-IP2Location-AS-CIDR` | AS-associated CIDR range |

These header names reflect the current source implementation. Numeric fields may be formatted as numeric strings even when a particular lookup has limited data; validate values before relying on them.

## Understanding ASN Lookup

An **Autonomous System Number (ASN)** identifies an autonomous network participating in Internet routing. ASN data helps developers understand which network operator is associated with an IP address.

This middleware exposes available ASN data through `X-IP2Location-ASN`, `X-IP2Location-AS`, `X-IP2Location-AS-Domain`, `X-IP2Location-AS-Usage-Type`, and `X-IP2Location-AS-CIDR`.

ASN lookup is useful for network troubleshooting, traffic analytics, and investigation of unexpected traffic. For offline IP and ASN data, see [IP2Location LITE](https://www.ip2location.com/database/lite).

## Troubleshooting

| Issue | What to check |
|---|---|
| `unrecognized directive: ip2location` | Build Caddy using `xcaddy --with ...` and run that binary. |
| `bin_path is required when mode is local` | Supply a readable BIN file path. |
| `api_key is required when mode is remote` | Supply the key in `api_key`; check environment placeholder expansion. |
| Missing headers in the application | Check handler order, successful lookups, and how your application exposes incoming headers. |
| Some fields are unavailable | Check the BIN package or API plan; not all packages provide every field. |
| Incorrect visitor IP | Inspect your reverse-proxy topology and forwarding headers. |
| Remote lookups fail | Check credentials, outbound connectivity, API quota, and upstream availability. |

**Lookup error behavior:** The current middleware continues to the next Caddy handler when a lookup fails; it does not produce a dedicated error response or enriched headers for that failed lookup.

## Security and Responsible Use

- **Trust forwarding headers only from controlled proxies.** The current IP selection logic checks `CF-Connecting-IP`, `X-Real-IP`, `X-Forwarded-For`, then `Forwarded`, before falling back to the connection address. In an untrusted deployment, a client may be able to spoof these values.
- **Protect the API key** and avoid publishing Caddy configuration containing real credentials.
- **Treat enrichment headers as application inputs.** Do not use geolocation or ASN alone to establish identity or prove malicious activity.
- **IP geolocation is not GPS.** Location results are estimates and may be inaccurate, especially for VPNs, mobile networks, and shared infrastructure.
- **Limit access to diagnostic pages** that reveal visitor or network information.

## Related IP2Location Products

| Product | Purpose |
|---|---|
| [IP2Location.io](https://www.ip2location.io/) | Hosted IP geolocation, ASN, and network intelligence APIs. |
| [IP2Location LITE](https://www.ip2location.com/database/lite) | Free downloadable IP geolocation and ASN databases for offline use. |
| [IP2Location Database](https://www.ip2location.com/database/ip2location) | Commercial offline IP geolocation databases. |
| [IP2Proxy](https://www.ip2location.com/database/proxy) | IP proxy, VPN, TOR, and anonymizer intelligence. |
| [IP2WHOIS](https://www.ip2whois.com/) | Domain WHOIS lookup services. |

## Documentation and Support

- [IP2Location.io API documentation](https://www.ip2location.io/ip2location-documentation)
- [Report issues and request features](https://github.com/ip2location/ip2location-caddy/issues)
- [Caddy documentation](https://caddyserver.com/docs/)

## License

Licensed under the **MIT License**. See [LICENSE](https://github.com/ip2location/ip2location-caddy/blob/main/LICENSE).
