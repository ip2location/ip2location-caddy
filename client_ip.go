package ip2locationcaddy

import (
	"net"
	"net/http"
	"strings"
)

func getClientIP(r *http.Request) string {
	// 1. Cloudflare header
	if ip := strings.TrimSpace(r.Header.Get("CF-Connecting-IP")); ip != "" {
		return cleanIP(ip)
	}

	// 2. Common reverse proxy header
	if ip := strings.TrimSpace(r.Header.Get("X-Real-IP")); ip != "" {
		return cleanIP(ip)
	}

	// 3. X-Forwarded-For can contain multiple IPs:
	// client, proxy1, proxy2
	if xff := strings.TrimSpace(r.Header.Get("X-Forwarded-For")); xff != "" {
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			return cleanIP(strings.TrimSpace(parts[0]))
		}
	}

	// 4. RFC 7239 Forwarded header, example:
	// Forwarded: for=203.0.113.195;proto=https;host=example.com
	if forwarded := strings.TrimSpace(r.Header.Get("Forwarded")); forwarded != "" {
		if ip := parseForwardedFor(forwarded); ip != "" {
			return cleanIP(ip)
		}
	}

	// 5. Fallback to direct connection IP
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return cleanIP(host)
	}

	return cleanIP(r.RemoteAddr)
}

func cleanIP(ip string) string {
	ip = strings.TrimSpace(ip)
	ip = strings.Trim(ip, `"`)

	// Handle IPv6 bracket format: [2001:db8::1]
	ip = strings.TrimPrefix(ip, "[")
	ip = strings.TrimSuffix(ip, "]")

	// If it is "IP:port", remove port.
	if host, _, err := net.SplitHostPort(ip); err == nil {
		ip = host
	}

	return ip
}

func parseForwardedFor(header string) string {
	// Example:
	// Forwarded: for=203.0.113.195;proto=https
	// Forwarded: for="[2001:db8:cafe::17]:4711"
	entries := strings.Split(header, ",")

	for _, entry := range entries {
		params := strings.Split(entry, ";")

		for _, param := range params {
			param = strings.TrimSpace(param)

			if strings.HasPrefix(strings.ToLower(param), "for=") {
				value := strings.TrimSpace(param[4:])
				value = strings.Trim(value, `"`)

				// Remove IPv6 bracket
				value = strings.TrimPrefix(value, "[")
				value = strings.TrimSuffix(value, "]")

				// If value still has port, remove it.
				if host, _, err := net.SplitHostPort(value); err == nil {
					return host
				}

				return value
			}
		}
	}

	return ""
}
