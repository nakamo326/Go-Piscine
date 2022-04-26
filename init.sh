#! /bin/zsh

if [ $# != 2 ]; then
  return 1
fi

DIR=$1
N=$2
cd $DIR
mkdir -p  ex{00..$N}/vendor/ft ex{00..$N}/vendor/piscine
touch ex{00..$N}/main.go 
find . -name "*ft" -type d -exec cp ../printrune.go {} \;
cd ..