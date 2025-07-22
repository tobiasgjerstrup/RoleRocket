# RoleRocket
I was tired of making a user system for all the random projects I made. So now I'm just gonna make this, also this seems like an alright way of learning GO.

### Errors
If you get something like this
```
Error creating logs Error=Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
2025/07/22 09:57:08 Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
exit status 1
```
install a C compiler [Microsofts recommendation](https://code.visualstudio.com/docs/cpp/config-mingw). Once installed type
```bash
set CGO_ENABLED=1
```
in a terminal.

if you get stuck trying to install MinGW with pacman, open C:/msys/etc/pacman.conf and set `SigLevel = Never` and then run the MinGW install. Please revert it back to `SigLevel = Required` when you're done