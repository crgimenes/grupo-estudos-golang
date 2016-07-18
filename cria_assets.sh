#!/bin/bash
rm -f ./assets/*.html
find ./ -depth -name "*.md" -exec sh -c 'markdown "$1">"./assets/${1%.md}.html"' _ {} \;
sed -i -e 's/.md/.html/g' ./assets/*.html
rm -f ./assets/*.html-e


