

### A Golang command-line parsing library (pre-alpha).

This will be similar to python `argparse` package,  
i.e. featurefull, but not too high-level - no mvc or similar stuff,  
so this library will be good for simple or not too complex projects  
or for wrapping other more high-level libraries around.

Reason why I'm making this: all major Go command-line  
arg parsing libraries I had tried are too buggy and lack features I need.

Exemplary bugs / lacking features I encountered:

- `--arg value1 --arg value2` results in `arg==value2`
- `--arg value1 value2` results in `arg==value2`
- no way to get an array from something like the above or `--arg val1 val2` etc
- no groups / subparsers / required args etc

I am aiming at fixing all of the above here.
