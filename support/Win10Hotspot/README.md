# Windows 10 Wifi Hotspot

Description: This is used to intercept communications between the PicoBrew machine and redirect them to an AWS Cloud Instance.

## Instructions

1. [Enable HotSpot in Windows 10](https://www.onmsft.com/how-to/how-to-use-wireless-hotspots-in-windows-10)
Note: The Wifi Network must be 2.4ghz.  If you attach to a 5.0ghz for your connection, the hotspot will be 5.0ghz and the PicoBrew won't connect

2. Run cmd.exe (Command Prompt) AS ADMINISTRATOR.   Click Yes when prompted (if you're not prompted - or "Access is denied" - you're not admin)

3. Find the URL in the "Output" tab of CloudFormation under PicoServer URL.   Eg: https://planbeerloadbalancernetwork-a4afcf0ad60dbf49.elb.us-east-2.amazonaws.com/

4. Run adjusthosts.bat <URL> (without the 'https://' and '/')
eg: adjusthosts.bat planbeerloadbalancernetwork-a4afcf0ad60dbf49.elb.us-east-2.amazonaws.com

5. Goto http://www.picobrew.com and you should see the PlanBeer UI


## When Done (or you want to go to the PicoBrew Site)
1. Run cmd.exe (Command Prompt) AS ADMINISTRATOR.   Click Yes when prompted (if you're not prompted, you're not admin)

2. Run reverthosts.bat