
# Extism in golang with msys2

To run this on windows we need https://www.msys2.org/

Install gcc with pacman

The compiled binary `libs` were download from: https://github.com/extism/extism/releases >
libextism-x86_64-pc-windows-msvc-v0.X.X.tar.gz

```shellscript
# installed tinygo + binaryen dependency
scoop install tinygo
scoop install binaryen

sh ./build.sh
```

To run the program:
```shellscript
cd out
./program.exe Hello
```
```output
starting
Count: 2
```
