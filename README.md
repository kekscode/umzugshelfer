# Umzugshelfer

The umzugshelfer fetches the current external IP address and updates a given
cloudflare DNS record.

## Usage (client)

Set the following environment Variables:

* CF_API_KEY="your-cloudflare-api-key"
* CF_API_EMAIL="your-cloudflare-account-email-address"
* UZH_DNS_ZONE="example.com"
* UZH_DNS_A_RECORD="www.example.com"
* UZH_PUBLIC_IP_PROVIDER="https://ifconfig.me/ip"

and then execute the program.

The `UZH_PUBLIC_IP_PROVIDER` has no default to protect your privacy, so it
has to be set to either a trusted provider or this tool in Servermode hosted
on a trusted system you control yourself.

## Usage (Server-mode)

In Server-Mode the tool is started like:

$> umzugshelfer 0.0.0.0:8080

to bind on a local IPv4:Port an serve as an public ip address provider you
can host for yourself. This is *not* supposed to work for IPv6 yet.

## Umzugshelfer?

Umzugshelfer is german for relocation assistant. I named it like that because
a) it helps to "relocate" your current DNS ip address entry (old home) to the
new address (new home) and b) i just couldn't come up with a good name
illustrating this use case in english.

## Why?

There are many tools doing this. I wanted to have my own.

## No IPv6 support?

No, i did not need it. Merge requests are welcome though.