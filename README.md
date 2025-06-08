# go-saffron

## Errors

* Undefined references
  You get this error if you are not linking against the SFML libraries correctly.
  libraries).
    ```
    /tmp/go-build/cgo-gcc-prolog:8853:(.text+0x5377): undefined reference to 'sfWindowBase_setVisible'
    ```

  Fix by environment variables:
  ```
  CGO_CFLAGS=-I/path/to/sfml/include
  CGO_LDFLAGS=-L/path/to/sfml/lib
  LD_LIBRARY_PATH=/path/to/sfml/lib:$LD_LIBRARY_PATH
  ```