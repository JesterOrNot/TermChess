nim c --verbosity:0 --hints:off --opt:speed -d:release src/lib/changeToBlue.nim && mv src/lib/changeToBlue src/lib/bin
nim c --verbosity:0 --hints:off --opt:speed -d:release src/lib/changeToRed.nim && mv src/lib/changeToRed src/lib/bin
nim c --verbosity:0 --hints:off --opt:speed -d:release src/lib/resetColor.nim && mv src/lib/resetColor src/lib/bin
clear
go run src/main.go