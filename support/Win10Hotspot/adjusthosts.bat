@echo off
echo [Must be run as ADMINISTRATOR]
copy c:\windows\system32\drivers\etc\hosts.planbeer.bak c:\windows\system32\drivers\etc\hosts /y
copy c:\windows\system32\drivers\etc\hosts c:\windows\system32\drivers\etc\hosts.planbeer.bak /y
for /f "skip=4 usebackq tokens=2" %%a in (`nslookup %1. 1.1.1.1`) do echo %%a picobrew.com >> c:\windows\system32\drivers\etc\hosts & echo.%%a www.picobrew.com >> c:\windows\system32\drivers\etc\hosts 
echo.
echo [www.picobrew.com set to %1 if no errors]