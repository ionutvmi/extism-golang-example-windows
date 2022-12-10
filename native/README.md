# Run extism natively on windows

## Install tinygo and gcc
```shellscript
scoop install tinygo
# dependency for tinygo
scoop install binaryen
scoop install gcc
```

## To compile the main program + the plugin
```shellscript
./build.bat
```

## To run the program
```shellscript
cd build

.\program.exe Hello
```
```output
starting
Count: 2
```
