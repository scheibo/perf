# PERF

![release-candidate](http://img.shields.io/badge/status-release--candidate-green.svg)&nbsp;
[![Build Status](http://img.shields.io/travis/scheibo/perf.svg)](https://travis-ci.org/scheibo/perf)

PERF (Performance Equivalency Rating Formula) implements a method for scoring
road cycling performances on climbs for ranking and comparison purposes. See
[scheibo.github.io/perf](https://scheibo.github.io/perf) for more details.

    $ go install github.com/scheibo/perf
    $ ./perf -t=18m00s -d=4809 -gr=8.12 -h 303
    18:00 (4.81 km @ 8.12%) = 508.09

The generated GoDoc can be viewed at
[godoc.org/github.com/scheibo/perf](https://godoc.org/github.com/scheibo/perf).
