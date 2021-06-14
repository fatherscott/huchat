FOR %%I in (proto/*.proto) DO protoc -I=proto --go_out=../packet/%%~nI %%~nxI
pause