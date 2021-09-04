#!/bin/bash -eu

if [ $# = 0 ]; then
    echo "USAGE"
    echo "  ./init.sh PROJECT_PATH"
    exit 1
fi

rename_project () {
    a=(${1//// })
    ppath=$1 pname=${a[${#a[@]}-1]}

    find . -type f -path ./init.sh -prune -o -type d -name .git -prune -o -type f -exec sed -i "s#kzmake/skeleton#$ppath#g" {} \;
    find . -type f -path ./init.sh -prune -o -type d -name .git -prune -o -type f -exec sed -i "s#skeleton#$pname#g" {} \;
}

echo $1

rename_project $1
