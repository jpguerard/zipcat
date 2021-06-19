#!/bin/dash

build(){
  local OPTIONS=""
  local GO="go"
  if [ "$3" = "linux64" ] ; then
    OPTIONS="-buildmode=pie"
  fi
  CGO_ENABLED=0 GOOS="$1" GOARCH="$2" ${GO} build -trimpath ${OPTIONS}
}

make_tar(){
  echo " * $3"
  build "$1" "$2"
  file zipcat
  tar cvjf "release/zipcat_$3.tar.bz2" zipcat LICENSE README.md
}

make_zip(){
  echo " * $3"
  build "$1" "$2"
  file zipcat.exe
  zip "release/zipcat_$3.zip" zipcat.exe LICENSE README.md
}

rm -f release/*
rmdir release
mkdir -p release
go mod tidy
make_zip windows 386   win32
make_zip windows amd64 win64
make_tar darwin  386   darwin32
make_tar darwin  amd64 darwin64
make_tar linux   386   linux32
make_tar linux   amd64 linux64
rm -f zipcat.exe

