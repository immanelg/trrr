Simple translation tool

# Install
Install with Go tools (to something like ~/.go/bin/trrr):
```
CGO_ENABLED=0 go install -trimpath -ldflags="-s -w" github.com/immanelg/trrr@latest
```
Or the same:
```
git clone --depth=1 https://github.com/immanelg/trrr
cd trrr
CGO_ENABLED=0 go install -trimpath -ldflags="-s -w"
```

Or install from GitHub releases:
```
curl -LO https://github.com/immanelg/trrr/releases/download/v0.1/trrr-v0.1-linux-amd64.tar.gz
mkdir trrr
tar xzf trrr-v0.1-linux-amd64.tar.gz -C trrr
sudo install -v -m755 trrr /usr/local/bin  # system-wide
```

# Usage
```
trrr [OPTIONS]... [SRC]:[DST] [CONTENT]
```

With no CONTENT, reads standard input. 

If `SRC` is empty, it defaults to `auto`. 

If `DST` is empty, it defaults to `en`.

# Backends 
- Google Translate

Needs more! Send your PRs. Please... :(

# Examples
Translate from English to Spanish. Type text to stdin via cat:
```
$ cat | trrr en:es
Things!
<Ctrl-D>
¡Cosas!
```

Auto-detect source language, translate to Hungarian and read text from the second argument:
```
$ trrr :hu 'things!'
A dolgok!
```

Translate text to Russian from X11 primary clipboard (selection) and show the result in a notification:
```
xclip -o | trrr :ru | xargs -0 -I '{}' notify-send -- "trrr" '{}'
```

Same as above, but also copy the result to clipboard:
```
xclip -o | trrr :ru | tee >(xclip -selection clipboard) | xargs -0 -I '{}' notify-send -- "trrr" '{}'
```

Example for sxhkd config (~/.config/sxhkd/sxhkdrc):
```
super + {w,W}
    xclip -o | trrr :{ru,en} | xargs -0 -I '\{\}' notify-send -- "trrr" '\{\}'
```



# License 
0BSD

