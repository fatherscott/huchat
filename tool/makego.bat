FOR %%I in (../input/*.proto) DO protoc -I=../Input --go_out=../../../../Packet/%%~nI %%~nxI
pause