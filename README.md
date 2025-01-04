# DNS Query Tool

Simple DNS Query Tool is a simple tool to send DNS queries to a DNS server with 
additional options.

Options:

* `-domain`: Domain name to query.
* `-type`: Query type (A, AAAA, CNAME, MX, NS, PTR, SOA, TXT).
* `-server`: Address of the DNS server.
* `-port`: Port of the DNS server.
* `-secret`: Secret key to protect the query.

Example:

Common DNS query:

```shell
./dns-query-tool -domain google.com -type A -server 1.1.1.1 -port 53
```

Custom DNS query with secret key:

```shell
./dns-query-tool -domain google.com -type A -server 127.0.0.1 -port 5302 -secret secret-key
```