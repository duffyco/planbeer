@echo off
copy c:\windows\system32\drivers\etc\hosts.planbeer.bak c:\windows\system32\drivers\etc\hosts /y
del /f c:\windows\system32\drivers\etc\hosts.planbeer.bak
echo [original settings enabled]