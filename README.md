# wt

`wt` (workout timer) is an alternative to the timer in `gnome-clocks` with better
keyboard functionality. It takes a single argument of a simple duration, waits
that long, then shows an OS notification and plays a sound.

## examples

```sh
# two equivalent commands to notify after 30 seconds
wt 30s
wt 0.5m

# notify after 5 minutes
wt 5m

# notify after an hour
wt 1h
```

![example](example1.png)