#!/bin/bash

main() {
    local update="$1"

    git tag -d 1.0
    git push origin --delete 1.0

    git add .
    git commit -m "$update"
    git push

    git tag -a 1.0 -m "1.0"
    git push origin 1.0
}

main "$1"
