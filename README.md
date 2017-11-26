# chaos-scheduler ☀🌧⚡⛅ vs. 💻✨

The chaos scheduler is a non-preemptive priority-based scheduling algorithm which does
unexpected things. It has two modes: Random, and Weather. The goal here is to
experiment, possibly with terrible ideas.

## Running 

```
-r    random mode               | bool  default: true
-w    weather mode              | bool  default: true
-s    shortest-job-first mode   | bool  default: true
-t    set size of threadpool    | int   default: 5
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
