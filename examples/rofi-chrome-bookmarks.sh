#!/usr/bin/env bash

rofi_dmenu(){
  rofi -dmenu -i -matching fuzzy -p "Select bookmark"\
    -location 0 -bg "#F8F8FF" -fg "#000000" -hlbg "#ECECEC" -hlfg "#0366d6"
}

main() {
  url=$(/home/sendhil/go/bin/chrome-bookmarks  | rofi_dmenu | xargs -i -0 /home/sendhil/go/bin/chrome-bookmarks --find-bookmark-url "{}")
  xdg-open $url
}

main

exit 0
