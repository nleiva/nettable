# Parsing Telemetry data from IOS XR YANG models

[![GoDoc](https://godoc.org/github.com/nleiva/nettable?status.svg)](https://godoc.org/github.com/nleiva/nettable) 
[![Build Status](https://travis-ci.org/nleiva/nettable.svg?branch=master)](https://travis-ci.org/nleiva/nettable) 
[![codecov](https://codecov.io/gh/nleiva/nettable/branch/master/graph/badge.svg)](https://codecov.io/gh/nleiva/nettable) 
[![Go Report Card](https://goreportcard.com/badge/github.com/nleiva/nettable)](https://goreportcard.com/report/github.com/nleiva/nettable)
[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

`nettable` reads an IOS XR Telemetry file and parses relevant info in order to produce a summary table. This exercise is inspired on what was done for [tlvdecode](https://github.com/nleiva/tlvdecode). The goal is to illustrate how to organize telemetry data to facilitate connecting the dots between different data sets. 

While the examples presented here were created in an IPv6-only enviroment, adding IPv4 support is straight forward (TODO).

## Use

`nettable` reads an IOS XR Telemetry file, parses fields of interest and produces a summary table. This helps visualizing the data, however this info should be actually sent to a database.

### Options

- Option -`f` points to the input file 
- Option -`i` Type of information:
  - From *Cisco-IOS-XR-clns-isis-oper*
    - **isis-int**: IS-IS Interfaces (*:isis/instances/instance/interfaces/interface*)
    - **isis-nbr**: IS-IS Neigbors (*:isis/instances/instance/neighbors/neighbor*)
    - **isis-lsp**: IS-IS LSPs (*:isis/instances/instance/levels/level/detailed-lsps/detailed-lsp*)
  - From *Cisco-IOS-XR-infra-statsd-oper*
    - **int-count**: Interface Counters (*:infra-statistics/interfaces/interface/latest/generic-counters*)
    - **int-rate**: Interface Data Rates (*:infra-statistics/interfaces/interface/data-rate*)

### IS-IS Interfaces

From [showtable](example/showtable) example:

```console
$ ./showtable -f ../../input/isis-int.json -i isis-int
+------------------------+--------------------+---------------+-----------+--------------------------+--------------------+
|        HOSTNAME        |     INTERFACE      |    CONFIG     |  STATUS   |        FW ADDRESS        |       PREFIX       |
+------------------------+--------------------+---------------+-----------+--------------------------+--------------------+
| mrstn-5502-1.cisco.com | Loopback60         | isis-disabled | isis-up   | <nil>                    | 2001:558:2::1/128  |
| mrstn-5502-1.cisco.com | Bundle-Ether20     | isis-disabled | isis-down | fe80::28a:96ff:fe46:6cdc | 2001:f00:bc::/64   |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/0 | isis-enabled  | isis-up   | fe80::28a:96ff:fe46:6c00 | 2001:f00:ba::/64   |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/1 | isis-enabled  | isis-up   | fe80::28a:96ff:fe46:6c04 | 2001:db8:cafe::/64 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/1 | isis-enabled  | isis-up   | fe80::28a:96ff:fe46:6c04 | 2001:f00:bb::/64   |
+------------------------+--------------------+---------------+-----------+--------------------------+--------------------+
```

```console
$ ./showtable -f ../../input/isis-int2.json -i isis-int
+------------------------+--------------------+---------------+-----------+--------------------------+--------------------+
|        HOSTNAME        |     INTERFACE      |    CONFIG     |  STATUS   |        FW ADDRESS        |       PREFIX       |
+------------------------+--------------------+---------------+-----------+--------------------------+--------------------+
| mrstn-5502-2.cisco.com | Loopback60         | isis-disabled | isis-up   | <nil>                    | 2001:558:2::2/128  |
| mrstn-5502-2.cisco.com | Bundle-Ether20     | isis-disabled | isis-down | fe80::28a:96ff:fe46:34db | 2001:f00:bc::/64   |
| mrstn-5502-2.cisco.com | Bundle-Ether30     | isis-disabled | isis-down | fe80::28a:96ff:fe46:34da | 2001:f00:bd::/64   |
| mrstn-5502-2.cisco.com | Bundle-Ether40     | isis-disabled | isis-down | fe80::28a:96ff:fe46:34d9 | 2001:f00:be::/64   |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/0 | isis-enabled  | isis-up   | fe80::28a:96ff:fe46:3400 | 2001:f00:ba::/64   |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/1 | isis-enabled  | isis-up   | fe80::28a:96ff:fe46:3404 | 2001:db8:cafe::/64 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/1 | isis-enabled  | isis-up   | fe80::28a:96ff:fe46:3404 | 2001:f00:bb::/64   |
+------------------------+--------------------+---------------+-----------+--------------------------+--------------------+
```

### IS-IS Neigbors

From [showtable](example/showtable) example:

```console
$ ./showtable -f ../../input/isis-nbr.json -i isis-nbr
+------------------------+--------------------+--------------+----------------+--------------------------+
|        HOSTNAME        |     INTERFACE      |     AREA     |   REMOTE ID    |        FW ADDRESS        |
+------------------------+--------------------+--------------+----------------+--------------------------+
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/0 | 49.0000.0162 | 0151.0250.0002 | fe80::28a:96ff:fe46:3400 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/1 | 49.0000.0162 | 0151.0250.0002 | fe80::28a:96ff:fe46:3404 |
+------------------------+--------------------+--------------+----------------+--------------------------+
```

```console
$ ./showtable -f ../../input/isis-nbr2.json -i isis-nbr
+------------------------+--------------------+--------------+----------------+--------------------------+
|        HOSTNAME        |     INTERFACE      |     AREA     |   REMOTE ID    |        FW ADDRESS        |
+------------------------+--------------------+--------------+----------------+--------------------------+
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/0 | 49.0000.0162 | 0151.0250.0001 | fe80::28a:96ff:fe46:6c00 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/1 | 49.0000.0162 | 0151.0250.0001 | fe80::28a:96ff:fe46:6c04 |
+------------------------+--------------------+--------------+----------------+--------------------------+
```

### IS-IS LSPs

From [showtable](example/showtable) example:

```console
$ ./showtable -f ../../input/isis-lsp.json -i isis-lsp
+----------------+------------------------+----------------+--------+
|    LOCAL ID    |        HOSTNAME        |   REMOTE ID    | METRIC |
+----------------+------------------------+----------------+--------+
| 0151.0250.0001 | mrstn-5502-1.cisco.com | 0151.0250.0002 |    333 |
| 0151.0250.0002 | mrstn-5502-2.cisco.com | 0151.0250.0001 |     10 |
+----------------+------------------------+----------------+--------+
```

```console
$ ./showtable -f ../../input/isis-lsp2.json -i isis-lsp
+----------------+------------------------+----------------+--------+
|    LOCAL ID    |        HOSTNAME        |   REMOTE ID    | METRIC |
+----------------+------------------------+----------------+--------+
| 0151.0250.0001 | mrstn-5502-1.cisco.com | 0151.0250.0002 |    333 |
| 0151.0250.0002 | mrstn-5502-2.cisco.com | 0151.0250.0001 |     10 |
+----------------+------------------------+----------------+--------+
```

### Interface Counters

From [showtable](example/showtable) example:

```console
$ ./showtable -f ../../input/int-count.json -i int-count
+------------------------+---------------------+-----------+-----------+-------+---------+----------+
|        HOSTNAME        |      INTERFACE      | PKTS SENT | PKTS RECV | TRANS | IN ERRS | OUT ERRS |
+------------------------+---------------------+-----------+-----------+-------+---------+----------+
| mrstn-5502-1.cisco.com | Null0               |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-1.cisco.com | Bundle-Ether20      |    769131 |    768015 |     0 |       0 |        0 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/0  |     43764 |     43753 |     5 |       0 |        0 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/47 |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/46 |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/1  |   1846705 |   1740573 |    19 |       0 |        0 |
...
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/21 |    461053 |    461742 |     4 |       0 |        0 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/22 |    308078 |    306273 |     4 |       0 |        0 |
...
| mrstn-5502-1.cisco.com | FortyGigE0/0/0/34   |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-1.cisco.com | FortyGigE0/0/0/38   |         0 |         0 |     0 |       0 |        0 |
+------------------------+---------------------+-----------+-----------+-------+---------+----------+
```

```console
$ ./showtable -f ../../input/int-count2.json -i int-count
+------------------------+---------------------+-----------+-----------+-------+---------+----------+
|        HOSTNAME        |      INTERFACE      | PKTS SENT | PKTS RECV | TRANS | IN ERRS | OUT ERRS |
+------------------------+---------------------+-----------+-----------+-------+---------+----------+
| mrstn-5502-2.cisco.com | Null0               |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | Bundle-Ether20      |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | Bundle-Ether30      |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | Bundle-Ether40      |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | MgmtEth0/RP0/CPU0/0 |   1886191 |    426258 |     3 |       0 |        0 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/0  |    313546 |    313561 |     5 |       0 |        0 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/47 |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/46 |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/45 |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/1  |    741369 |    753244 |     1 |       0 |        0 |
...
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/21 |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/22 |         0 |         0 |     0 |       0 |        0 |
...
| mrstn-5502-2.cisco.com | FortyGigE0/0/0/34   |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | FortyGigE0/0/0/14   |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/0   |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/1   |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/2   |         0 |         0 |     0 |       0 |        0 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/3   |         0 |         0 |     0 |       0 |        0 |
+------------------------+---------------------+-----------+-----------+-------+---------+----------+
```

### Interface Data Rates

From [showtable](example/showtable) example:

```console
$ ./showtable -f ../../input/int-rate.json -i int-rate
+------------------------+---------------------+---------+----------+-----------+
|        HOSTNAME        |      INTERFACE      | IN KBPS | OUT KBPS |    BW     |
+------------------------+---------------------+---------+----------+-----------+
| mrstn-5502-1.cisco.com | Null0               |       0 |        0 |         0 |
| mrstn-5502-1.cisco.com | Bundle-Ether20      |       0 |        0 |         0 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/0  |     293 |      293 | 100000000 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/47 |       0 |        0 | 100000000 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/46 |       0 |        0 | 100000000 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/1  |       7 |       10 | 100000000 |
...
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/21 |       0 |        0 | 100000000 |
| mrstn-5502-1.cisco.com | HundredGigE0/0/0/22 |       0 |        0 | 100000000 |
...
| mrstn-5502-1.cisco.com | FortyGigE0/0/0/34   |       0 |        0 |  40000000 |
| mrstn-5502-1.cisco.com | FortyGigE0/0/0/38   |       0 |        0 |  40000000 |
+------------------------+---------------------+---------+----------+-----------+
```

```console
$ ./showtable -f ../../input/int-rate2.json -i int-rate
+------------------------+---------------------+---------+----------+-----------+
|        HOSTNAME        |      INTERFACE      | IN KBPS | OUT KBPS |    BW     |
+------------------------+---------------------+---------+----------+-----------+
| mrstn-5502-2.cisco.com | Null0               |       0 |        0 |         0 |
| mrstn-5502-2.cisco.com | Bundle-Ether20      |       0 |        0 |         0 |
| mrstn-5502-2.cisco.com | Bundle-Ether30      |       0 |        0 |         0 |
| mrstn-5502-2.cisco.com | Bundle-Ether40      |       0 |        0 |         0 |
| mrstn-5502-2.cisco.com | MgmtEth0/RP0/CPU0/0 |       1 |        4 |   1000000 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/0  |     178 |      176 | 100000000 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/47 |       0 |        0 | 100000000 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/46 |       0 |        0 | 100000000 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/45 |       0 |        0 | 100000000 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/1  |       7 |        7 | 100000000 |
...
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/21 |       0 |        0 | 100000000 |
| mrstn-5502-2.cisco.com | HundredGigE0/0/0/22 |       0 |        0 | 100000000 |
...
| mrstn-5502-2.cisco.com | FortyGigE0/0/0/34   |       0 |        0 |  40000000 |
| mrstn-5502-2.cisco.com | FortyGigE0/0/0/14   |       0 |        0 |  40000000 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/0   |       0 |        0 |  10000000 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/1   |       0 |        0 |  10000000 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/2   |       0 |        0 |  10000000 |
| mrstn-5502-2.cisco.com | TenGigE0/0/0/10/3   |       0 |        0 |  10000000 |
+------------------------+---------------------+---------+----------+-----------+
```

### IPv6 Routing Table

From [showtable](example/showtable) example:

```console
$ ./showtable -f ../../input/rib-ipv6.json -i rib-ipv6
+------------------------+-------------------------+-----------+--------------------------+--------------------------+--------+
|        HOSTNAME        |         PREFIX          | PROTOCOL  |         NEXT HOP         |          SOURCE          | METRIC |
+------------------------+-------------------------+-----------+--------------------------+--------------------------+--------+
| mrstn-5502-1.cisco.com | ::/0                    | static    | 2001:420:2cff:1204::1    | 2001:420:2cff:1204::1    |      0 |
| mrstn-5502-1.cisco.com | 2001:420:2cff:1204::/64 | bgp       | 2001:db8:cafe::2         | 2001:db8:cafe::2         |      0 |
| mrstn-5502-1.cisco.com | 2001:558:2::1/128       | local     | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:558:2::2/128       | isis      | fe80::28a:96ff:fe46:3400 | fe80::28a:96ff:fe46:3400 |    334 |
| mrstn-5502-1.cisco.com | 2001:db8:33::1/128      | local     | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:db8:55::/64        | bgp       | 2001:db8:cafe::2         | 2001:db8:cafe::2         |      0 |
| mrstn-5502-1.cisco.com | 2001:db8:77::/64        | bgp       | 2001:db8:cafe::2         | 2001:db8:cafe::2         |      0 |
| mrstn-5502-1.cisco.com | 2001:db8:88:88::/64     | connected | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:db8:88:88::1/128   | local     | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:db8:cafe::/64      | connected | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:db8:cafe::1/128    | local     | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:f00:ba::/64        | connected | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:f00:ba::1/128      | local     | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:f00:bb::/64        | connected | ::                       | ::                       |      0 |
| mrstn-5502-1.cisco.com | 2001:f00:bb::1/128      | local     | ::                       | ::                       |      0 |
+------------------------+-------------------------+-----------+--------------------------+--------------------------+--------+
```

```console
$ ./showtable -f ../../input/rib-ipv62.json -i rib-ipv6
+------------------------+--------------------------------+-------------+--------------------------+--------------------------+--------+
|        HOSTNAME        |             PREFIX             |  PROTOCOL   |         NEXT HOP         |          SOURCE          | METRIC |
+------------------------+--------------------------------+-------------+--------------------------+--------------------------+--------+
| mrstn-5502-2.cisco.com | ::/0                           | static      | 2001:420:2cff:1204::1    | 2001:420:2cff:1204::1    |      0 |
| mrstn-5502-2.cisco.com | 4::4/128                       | static      | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 100::100/128                   | application | 4::4                     | 4::4                     |      0 |
| mrstn-5502-2.cisco.com | 101::101/128                   | application | 5::5                     | 5::5                     |      0 |
| mrstn-5502-2.cisco.com | 102::102/128                   | application | 6::6                     | 6::6                     |      0 |
| mrstn-5502-2.cisco.com | 2001:420:2cff:1204::/64        | connected   | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:420:2cff:1204::5502:2/128 | local       | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:558:2::1/128              | isis        | fe80::28a:96ff:fe46:6c00 | fe80::28a:96ff:fe46:6c00 |     11 |
| mrstn-5502-2.cisco.com | 2001:558:2::2/128              | local       | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:33::1/128             | bgp         | 2001:db8:cafe::1         | 2001:db8:cafe::1         |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:55::/64               | connected   | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:55::1/128             | local       | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:77::/64               | connected   | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:77::1/128             | local       | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:88:88::/64            | bgp         | 2001:db8:cafe::1         | 2001:db8:cafe::1         |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:cafe::/64             | connected   | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:db8:cafe::2/128           | local       | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:f00:ba::/64               | connected   | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:f00:ba::2/128             | local       | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:f00:bb::/64               | connected   | ::                       | ::                       |      0 |
| mrstn-5502-2.cisco.com | 2001:f00:bb::2/128             | local       | ::                       | ::                       |      0 |
+------------------------+--------------------------------+-------------+--------------------------+--------------------------+--------+
```

## Links

- [Pipeline](https://github.com/cisco/bigmuddy-network-telemetry-pipeline): Telemetry Collector
