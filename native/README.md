
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
cd out

.\program.exe Hello
```
```output
starting
Count: 2
```

The compiled binary `libs` were download from: https://github.com/extism/extism/releases >
libextism-x86_64-pc-windows-msvc-v0.X.X.tar.gz


