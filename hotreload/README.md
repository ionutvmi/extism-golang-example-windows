
# Example hot reloading WASM plugin

To compile the main program + the plugin:
```shellscript
.\build.bat
```

To run the main program:
```shellscript
cd out
.\program.exe

# the program will print the result from the plugin once per second
# leave the program running and start a new shell for future commands
```

Make changes to vowels.go to change the result.

To re-compile only the plugin:
```shellscript
.\build_plugin.bat

# In the initial shell of the program you will see the new result
```
