# Umzugshelfer

The umzugshelfer fetches the current external IP address and updates a given
cloudflare DNS record.

## Usage

Set the following environment Variables:

* CF_API_KEY="your-cloudflare-api-key"
* CF_API_EMAIL="your-cloudflare-account-email-address"
* UZH_DNS_ZONE="example.com"
* UZH_DNS_A_RECORD="www.example.com"

and then execute the program.

## Umzugshelfer?

Umzugshelfer is german for relocation assistant. I named it like that because
a) it helps to "relocate" your current DNS ip address entry (old home) to the
new address (new home) and b) i just couldn't come up with a good name
illustrating this use case in english.

## Why?

There are many tools doing this. I wanted to have my own.

## No IPv6???

No, i did not need it. Merge requests are welcome though.