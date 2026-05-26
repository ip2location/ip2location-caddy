# IP2Location Caddy Middleware

This is a Caddy Middleware that set geolocation data in headers using the client IP address and querying either the IP2Location BIN database or the IP2Location.io API.


# Requirement

When running in the local query mode, you will need to download an IP2Location BIN database to provide the geolocation data.
You can [register](https://www.ip2location.com/database/lite) for an account to download the free LITE database.
Then download the zipped file containing the BIN file. Extract and upload the BIN file to your server.

Another option is using the remote query mode. In this case, you can [sign up](https://www.ip2location.io/sign-up) for a free IP2Location.io API key and enjoy up to 50,000 queries per month.


# Compiling Caddy with the IP2Location Caddy Middleware

1. Make sure you have Go and xcaddy installed. Please refer to https://go.dev/doc/install and https://github.com/caddyserver/xcaddy for the installation steps.

2. Run the below to compile Caddy with the IP2Location Caddy Middleware.

```bash
xcaddy build --with github.com/ip2location/ip2location-caddy
```


# Configure the Caddyfile

Examples of Caddyfile configurations for both local and remote geolocation queries are shown below.
In the below examples, we also pass the geolocation headers to PHP but you can modify that for your own use case.

## Option 1: Using the IP2Location BIN database file for the local query mode

Modify the path to the IP2Location BIN file that you're using.

```caddy
{
    # Globally instruct Caddy to run your plugin before it processes FastCGI
    order ip2location before php_fastcgi
}

example.com {
    # Set the root folder where your PHP files (like index.php) live 
    root * /var/www/html

    # 1. First, look up the IP and inject the geolocation headers
    ip2location {
        mode local
        bin_path /path/to/IPV6-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE-ADDRESSTYPE-CATEGORY-DISTRICT-ASN.BIN
        header_prefix X-IP2Location
    }

    # 2. Then, pass the request (with headers included) to PHP-FPM
    # (Adjust the path to your actual php-fpm socket or use 127.0.0.1:9000)
    php_fastcgi unix//run/php/php8.4-fpm.sock

    # 3. Serve static files (images, CSS, JS) directly if they exist
    file_server
}
```

## Option 2: Using the IP2Location.io API for the remote query mode

Modify <YOUR_API_KEY> with your actual IP2Location.io API key.

```caddy
{
    # Globally instruct Caddy to run your plugin before it processes FastCGI
    order ip2location before php_fastcgi
}

example.com {
    # Set the root folder where your PHP files (like index.php) live 
    root * /var/www/html

    # 1. First, look up the IP and inject the geolocation headers
    ip2location {
        mode remote
        api_key <YOUR_API_KEY>
        header_prefix X-IP2Location
    }

    # 2. Then, pass the request (with headers included) to PHP-FPM
    # (Adjust the path to your actual php-fpm socket or use 127.0.0.1:9000)
    php_fastcgi unix//run/php/php8.4-fpm.sock

    # 3. Serve static files (images, CSS, JS) directly if they exist
    file_server
}
```


# Usage

Start the Caddy server to serve requests. Modify the path to the Caddyfile.

```bash
sudo ./caddy run --config /path/to/Caddyfile --adapter caddyfile
```


# Testing

Below is a simple PHP page to output the geolocation headers passed by Caddy. Upload this file to website folder and navigate to this page in your browser to view the geolocation output.

```php
<?php
$prefix = 'HTTP_X_IP2LOCATION_';
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

foreach ($fields as $val) {
	$key = $prefix . $val;
	if (isset($_SERVER[$key])) {
		echo $key . ': ' . $_SERVER[$key] . "\n<br>";
	}
}
```