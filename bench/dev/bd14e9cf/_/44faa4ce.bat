::WIN MYSQL  备份脚本
::## http://www.aqzt.com
::##email: ppabc@qq.com
::##robert yu
@echo off
set name=%date:~0,10%_%TIME:~0,2%-%TIME:~3,2%-%TIME:~6,2%
set name=%name: =%
set name=%name:/=-%
md D:\backup\%name%
"D:\MySQL Server 5.5\bin\mysqldump.exe" --opt -u test1 --password=I8g7Z0l1bub test1 > D:\backup\%name%\test1_%name%.sql
echo test1_%name% >>D:\backup_time.txt
"D:\MySQL Server 5.5\bin\mysqldump.exe" --opt -u test2 --password=S6k6x9U8lzw test2 > D:\backup\%name%\test2_%name%.sql
echo test2_%name% >>D:\backup_time.txt
"C:\Program Files (x86)\WinRAR\WinRAR.exe" a "D:\backup\%name%.rar" D:\backup\%name%
del /s/q/f D:\backup\%name%
exit