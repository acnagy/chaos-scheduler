# chaos-scheduler ☀🌧⚡⛅ vs. 💻✨

The chaos scheduler is a non-preemptive priority-based scheduling algorithm which does unexpected things. It has four modes: Random, Weather - Static, Weather - Variable, and Shortest Job First. The goal here is to experiment, possibly with terrible ideas.

### Definitions
* Random: threads are prioritized... at [pseudo-random](https://golang.org/pkg/math/rand/)
* Weather - Static: threads are prioritized based on a static latitude/longitude (Boston)
* Weather - Variable: threads are prioritized based on the weather at a _generated_ latitude/longitude. The coordinates are generated based off of the thread id
* Shortest Job First: threads are prioritized based on what should be the shortest running thread

### Work
Note: no real work is done here - these threads are executing a difficult job of sleeping for a pseudo-randomly generated period of nanoseconds

## Running 

```
-r    random mode               | bool  default: true
-w    weather mode (static)     | bool  default: true
-v    weather mode (variable)   | bool  default: true
-j    shortest-job-first mode   | bool  default: true
-t    set size of threadpool    | int   default: 5
-n    number of threads created | int   default: 10
```

Examples: 

```
chaos-scheduler                     | runs all modes, default threadpool size
chaos-scheduler -w=false -t=100     | does not run weather mode, 100 threads in pool
```

## Weather calculation

The weather-based priority factor is calculated as follows:

```
priority = (temp / pressure) * (wind_gust + precip_total)
```

This roughly favors stronger weather (especially in warmer temperatures), as precipitation is more probable in areas with lower pressure. Note - this little calculation would breakdown with division by zero if you were checking a weather station near the edge of the atmosphere. As of 2017, this is not a problem we face yet :) 
