# wt

`wt` (workout timer) is an alternative to the timer in gnome-clocks with better
keyboard functionality. It takes a single argument of a simple duration, waits
that long, then shows an OS notification and plays a sound.

## Examples

```sh
# two equivalent commands to notify after 30 seconds
wt 30s
wt 0.5m

# notify after 5 minutes
wt 5m

# notify after an hour
wt 1h
```

![example](images/example1.png)

## Installation

(You can try downloading the `wt` binary under [Releases](https://github.com/GSGBen/wt/releases), but it's only for amd64 Linux and there's no guarantee it's up to date).

Install Go, then try `go install github.com/gsgben/wt@latest`. If it throws errors about dependencies (like ALSA), fix those up then try installing it again. Known dependencies:

### Linux

#### ALSA / libasound2-dev

Try `sudo apt install libasound2-dev` or your distro's equivalent.

If you get a cryptic message about dependencies and `libasound2`, and you're on an Ubuntu-like distro: make sure you're subscribed to all updates, not just security updates. Then run `sudo apt update` and try installing it again.

## Improvements

The notification icon is copied out of the embedded filesystem into a temp location in the real filesystem. It would be better to call the version of the dbus notification function that takes image-data directly.

## Videogames

Hey, do you like videogames? If so please check out my game **Grab n' Throw** on Steam, and add it to your wishlist. One gamemode is like golf but on a 256 km^2 landscape, with huge throw power, powerups, and a moving hole. And there's plenty more!

<p align="center">
  <a href=https://store.steampowered.com/app/1813590/Grab_n_Throw/?utm_source=github_t2md>
    <img src="images/throwing_3.gif">