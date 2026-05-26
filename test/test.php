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
?>
