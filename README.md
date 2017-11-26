# chaos-scheduler ☀🌧⚡⛅ vs. 💻✨

The chaos scheduler is a non-preemptive priority-based scheduling algorithm which does
unexpected things. It has two modes: Random, and Weather. The goal here is to
experiment, possibly with terrible ideas.

## Running 

```
-w    weather mode              | bool  default: true
-r    random                    | bool  default: true
-t    set size of threadpool    | int)  default: 5
```

Example: 

`chaos-scheduler -w=true -t=100`

## Weather calculation

The weather-based priority factor is calculated as follows:

```
priority = (temp / pressure) * (wind_gust + precip_total)
```

This roughly favors stronger weather (especially in warmer temperatures), as precipitation is more probable in areas with lower pressure. Note - this little calculation would breakdown with division by zero if you were checking a weather station near the edge of the atmosphere. As of 2017, this is not a problem we face yet :) 
