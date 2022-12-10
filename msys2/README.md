
# Extism in golang with msys2

To run this on windows we need https://www.msys2.org/

Install gcc

Download the compiled binary from: https://github.com/extism/extism/releases >
libextism-x86_64-pc-windows-msvc-v0.X.X.tar.gz

```shellscript
# extism.h file added to C:\tools\msys64\mingw64\include
# extism.dll file added to C:\tools\msys64\mingw64\lib

# installed tinygo + binaryen dependency
scoop install tinygo
scoop install binaryen

sh ./build_main_program.sh

sh ./build_plugin.sh
```

To run the program:
```shellscript
cd build
./program.exe Hello
```
```output
starting
Count: 2
```
