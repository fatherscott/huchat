FOR %%I in (proto/*.proto) DO protoc -I=proto --go_out=../%%~nI %%~nxI
pause