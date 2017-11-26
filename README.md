# chaos-scheduler ☀🌧⚡⛅ vs. 💻✨

The chaos scheduler is a non-preemptive priority-based scheduling algorithm which does
unexpected things. It has two modes: Random, and Weather. The goal here is to
experiment, possibly with terrible ideas.

## Running 

```
-w    weather mode
-r    random
-t    set size of threadpool (integers required)
```

Example: 

`chaos-scheduler -w=true -t=100`
