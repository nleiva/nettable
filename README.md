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

### Interface Counters

From [showtable](example/showtable) example:

```console
$ /showtable -f ../../input/int-count.json -i int-count
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

## Links

- [Pipeline](https://github.com/cisco/bigmuddy-network-telemetry-pipeline): Telemetry Collector
