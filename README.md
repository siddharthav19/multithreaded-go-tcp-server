## Table of Contents

- [Test 1](#test-1)

  - [Summary](#summary:)

  - [Response time histogram:](#response-time-histogram:)
  - [Response time distribution:](#response-time-distribution:)
  - [Details (average, fastest, slowest):](<#details-(average,-fastest,-slowest):>)
  - [Status code distribution:](#status-code-distribution:)

- [Test 2](#test-2)

  - [Summary:](#summary:)

  - [Response time histogram:](#response-time-histogram:)
  - [Response time distribution:](#response-time-distribution:)
  - [Details (average, fastest, slowest):](<#details-(average,-fastest,-slowest):>)
  - [Status code distribution:](#status-code-distribution:)

- [Test 3](#test-3)

  - [Summary:](#summary:)

  - [Response time histogram:](#response-time-histogram:)
  - [Response time distribution:](#response-time-distribution:)
  - [Details (average, fastest, slowest):](<#details-(average,-fastest,-slowest):>)
  - [Status code distribution:](#status-code-distribution:)

# Introduction

- A basic TCP server which handles multiple concurrent requests written from scratch in go, benchmarked the server using oha

# Some Extensions

- Headers Compression
- Make the streamFile actual stream the file,lol
- Pipelining ?

# Test 1

➜ ~ oha -n 1000 -c 1000 http://localhost:4221/echo/heyfromoha

## Summary:

- Success rate: 100.00%
- Total: 0.6430 secs
- Slowest: 0.5486 secs
- Fastest: 0.1003 secs
- Average: 0.4432 secs
- Requests/sec: 1555.2087

#

- Total data: 9.77 KiB
- Size/request: 10 B
- Size/sec: 15.19 KiB

## Response time histogram:

- 0.100 [1] |
- 0.145 [204] |■■■■■■■■
- 0.190 [0] |
- 0.235 [0] |
- 0.280 [0] |
- 0.324 [0] |
- 0.369 [0] |
- 0.414 [0] |
- 0.459 [0] |
- 0.504 [0] |
- 0.549 [795] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■

## Response time distribution:

- 10.00% in 0.1215 secs
- 25.00% in 0.5124 secs
- 50.00% in 0.5223 secs
- 75.00% in 0.5314 secs
- 90.00% in 0.5362 secs
- 95.00% in 0.5387 secs
- 99.00% in 0.5437 secs
- 99.90% in 0.5486 secs
- 99.99% in 0.5486 secs

## Details (average, fastest, slowest):

- DNS+dialup: 0.4242 secs, 0.0212 secs, 0.5282 secs
- DNS-lookup: 0.0000 secs, 0.0000 secs, 0.0001 secs

## Status code distribution:

[200] 1000 responses

# Test 2

➜ ~ oha -n 2500 -c 2500 http://localhost:4221/echo/heyfromoha

## Summary:

- Success rate: 100.00%
- Total: 1.6117 secs
- Slowest: 1.5200 secs
- Fastest: 0.0590 secs
- Average: 0.6642 secs
- Requests/sec: 1551.1707

#

- Total data: 24.41 KiB
- Size/request: 10 B
- Size/sec: 15.15 KiB

## Response time histogram:

- 0.059 [1] |
- 0.205 [23] |
- 0.351 [44] |
- 0.497 [175] |■■■
- 0.643 [1606] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
- 0.789 [0] |
- 0.936 [0] |
- 1.082 [641] |■■■■■■■■■■■■
- 1.228 [0] |
- 1.374 [0] |
- 1.520 [10] |

## Response time distribution:

- 10.00% in 0.5203 secs
- 25.00% in 0.5489 secs
- 50.00% in 0.5738 secs
- 75.00% in 1.0118 secs
- 90.00% in 1.0259 secs
- 95.00% in 1.0290 secs
- 99.00% in 1.0319 secs
- 99.90% in 1.5197 secs
- 99.99% in 1.5200 secs

## Details (average, fastest, slowest):

- DNS+dialup: 0.6196 secs, 0.0020 secs, 1.5197 secs
- DNS-lookup: 0.0000 secs, 0.0000 secs, 0.0002 secs

## Status code distribution:

- [200] 2500 responses

# Test 3

➜ ~ oha -n 4000 -c 4000 http://localhost:4221/echo/heyfromoha

## Summary:

- Success rate: 100.00%
- Total: 2.2648 secs
- Slowest: 2.0513 secs
- Fastest: 0.0499 secs
- Average: 1.0802 secs
- Requests/sec: 1766.1650

#

- Total data: 39.06 KiB
- Size/request: 10 B
- Size/sec: 17.25 KiB

## Response time histogram:

- 0.050 [1] |
- 0.250 [35] |■
- 0.450 [89] |■■■
- 0.650 [731] |■■■■■■■■■■■■■■■■■■■■■■■■
- 0.850 [548] |■■■■■■■■■■■■■■■■■■
- 1.051 [244] |■■■■■■■■
- 1.251 [949] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
- 1.451 [439] |■■■■■■■■■■■■■■
- 1.651 [886] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
- 1.851 [0] |
- 2.051 [78] |■■

## Response time distribution:

- 10.00% in 0.5836 secs
- 25.00% in 0.7110 secs
- 50.00% in 1.1041 secs
- 75.00% in 1.3750 secs
- 90.00% in 1.5721 secs
- 95.00% in 1.5891 secs
- 99.00% in 2.0407 secs
- 99.90% in 2.0490 secs
- 99.99% in 2.0513 secs

## Details (average, fastest, slowest):

- DNS+dialup: 0.9782 secs, 0.0016 secs, 2.0482 secs
- DNS-lookup: 0.0000 secs, 0.0000 secs, 0.0008 secs

## Status code distribution:

- [200] 4000 responses

---
