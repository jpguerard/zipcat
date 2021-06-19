Zipcat
======

Zipcat is a small tool to quickly uncompress the content of a zip file.
It reads all files in the zip file and dump their contents to the
standard output.

Some sample usage :

1. Find some text in a zip file including text and binary files:

    zipcat mylib.jar | strings | grep mytext 

2. Count the lines of text in a zip file containing only text files:

    zipcat myarchive.zip | wc -l

