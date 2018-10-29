#!/usr/bin/env bash

# Rofi dmenu
# rofi_dmenu(){
#   rofi -dmenu -p "Bookmarks"
# }


# rofi_dmenu(){
#   rofi -dmenu -i -matching fuzzy -only-match -p "Select a repository > "\
#     -location 0 -bg "#F8F8FF" -fg "#000000" -hlbg "#ECECEC" -hlfg "#0366d6"
# }

rofi_dmenu(){
  rofi -dmenu -i -matching fuzzy -p "Select bookmark"\
    -location 0 -bg "#F8F8FF" -fg "#000000" -hlbg "#ECECEC" -hlfg "#0366d6"
}

main() {
  url=$(chrome-bookmarks | rofi_dmenu | xargs -i -0 chrome-bookmarks --find-bookmark-url "{}")
  xdg-open $url
}

main

exit 0
