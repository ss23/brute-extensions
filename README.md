# brute-extensions
Small script to verify which extensions return a 403 vs non 403 error for a given URL/domain

Simply `go run brute-extensions.go http://www.example.com/foo`

## why?
I was on a job where someone had a misconfiguration such that the IP whitelist only applied to some extensions, but PHP files for example were accessible to everyone.
This was a quick way to enumerate all extensions with the same issue.
