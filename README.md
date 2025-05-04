Simple translation tool

# Install
## Build from source
Build with Go and install to something like ~/.go/bin/trrr:
```sh
CGO_ENABLED=0 go install -trimpath -ldflags="-s -w" github.com/immanelg/trrr@latest
```

## Download a binary
Install from [Github releases](https://github.com/immanelg/trrr/releases). From there, you can figure out how to unzip a file.

# Usage
```sh
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
```sh
$ cat | trrr en:es
Things!
<Ctrl-D>
¡Cosas!
```

Auto-detect source language, translate to Hungarian and read text from the second argument:
```sh
$ trrr :hu 'things!'
A dolgok!
```

Translate text to Russian from X11 primary clipboard (selection) and show the result in a notification:
```sh
xclip -o | trrr :ru | xargs -0 -I '{}' notify-send -- "trrr" '{}'
```

Same as above, but also copy the result to clipboard:
```sh
xclip -o | trrr :ru | tee >(xclip -selection clipboard) | xargs -0 -I '{}' notify-send -- "trrr" '{}'
```

Prompt for text with rofi and display the translation with rofi:
```sh
rofi -dmenu -p 'translate' -l 0 | trrr :ru | xargs -0 -I\{\} rofi -p 'translation' -e \{\}
```

Example for sxhkd config (~/.config/sxhkd/sxhkdrc):
```conf
super + {w,W}
    xclip -o | trrr :{ru,en} | xargs -0 -I '\{\}' notify-send -- "trrr" '\{\}'

super + alt + {w,W}
    rofi -dmenu -p 'translate' -l 0 | trrr :{ru,en} | xargs -0 -I\{\} rofi -p 'translation' -e \{\}
```



# License 
0BSD

