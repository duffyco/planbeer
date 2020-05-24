# Plan B-eer - A Self-Contained Infrastructure for PicoBrew Z

THIS IS PROOF-OF-CONCEPT ONLY.  I CANNOT BEGIN TO LIST THE REASONS YOU SHOULDN'T RUN THIS.  NO WARRANTY EXPRESS OR IMPLIED.

Given PicoBrew's recent announcements, I'm hoping this won't be under development long-term.   I *love* their online components and I will continue to use their site as long as it's available.  

I also want the community to know that we don't have doorstops.  This *is* possible.  Run your electrical, build your area and buy your supplies - these machines here for the long run.

This is under heavy development.  I'll be making things more robust and building it out over the next few weeks.

----------------------------------------------------------------------------------------------------------------------------------------
<b> What is this? </b>

This project is a standalone, self-contained implementation that takes PicoBrew-XML recipes and interfaces with a PicoBrew Z to perform completely automated brew.   I have been succesfull in brewing recipes start-to-finish with this.

Using Raspberry Pi Hardware we'll:
- Build a Wifi Access Point and Attach a Z to it
- Redirect the Access Point's traffic to a Docker Container
- Build a Docker Container that runs software that allows the Z to perform tasks

<b> Implemented Functionality </b>

The following appear to work:
- Starts Up / Lists Recipes
- Recipe (with Multi-step Mash, Boil and Whirlpool.  All 4-adjuncts can be used)
- Coffee (Seems to work with all formats?)
- Sous Vide (Seems to work)
- Rinse (Works)
- Clean (Works)
- Circulate (Works)
- Rack (Works)

<b> What hasn't been tested?  </b>
- PicoStill
- PicoFerm
- Anything other than a Pico Z
- Group Brewing
- *Only 0.0.116 Firmware is supported*

<b> What do I need? </b>

Hardware: Raspberry PI (Tested on v3 & v4 - with the internal RPi Wifi)
OS: Raspbian Buster
Software: Go App on Docker 

<b> Alright!  I've got my ingredients loaded and ready to go! </b>

Hold on there Spark - *make sure you use water and do a test batch before brewing that $90 NEIPA Barleywine with Vintage Yeast*

<b> I'm skeptical.  Running this seems like a bad idea.</b>

Did you see my rigorous test harness?  No?  That's because it *is* a bad idea.  Especially right now.  Relax and Enjoy a Homebrew instead.   This question will eventually disappear.

<b> What do I do?  OK, ok, ok.  I get it.  But I still want to try it.  </b>

----------------------------------------------------------------------------------------------------------------------------------------

<b> Guide - Not hard.  About ~1hr </b>

<pre>
------   Ethernet(wire)     ----------     Wireless    -------------
Router  <-------------->     Rpi 3/4     <---------->   PicoBrew Z
------                      ----------                 -------------
                           OS: Raspbian
                            Vx: Docker
                               App
</pre>

1. Use Rufus (https://rufus.ie/ I use the portable version) and "burn" buster to sdcard. 
https://www.raspberrypi.org/downloads/raspbian/ - (2020-02-13 for the guide).

2. Go ahead and create an SSH file in the root.  This will activate the SSH server and save you from going downstairs.  I speak from experience.

For Windows users: navigate to D: | Create a Next Text Document | Change the Name to SSH  (no extension)

Eject the drive safely.

3. Start it up. Default username/pass is pi/raspberry.  Change the default password:
<pre> passwd pi </pre>

4. Remember your IP.
<pre>ifconfig 
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet *192.168.100.26* 
</pre>

5. Configure your WiFi country
<pre>sudo raspi-config
Select "4. Localisation Options"
Select "I4. Change Wifi Country" --> Set your country

5. Run the following:
<pre>sudo apt-get update
sudo apt-get -y install git </pre>

6. Clone the repo in the home directory:
<pre> git clone https://github.com/duffyco/planbeer.git </pre>

7. Run the WifiSetup:
<pre> cd ~/planbeer/bin </pre>
<pre> sudo ./setupWifi.sh </pre>

8. This will setup dhcp and hostapd for the RPi 3/4 default Wifi and restart.   Start scanning for a "Plan B" Network.  Password is: 12345678

9.  I'm assuming you've exported all your recipes from Picobrew.  You'll have a directory where they're stored. 

10. In your recipe directory, you can copy recipes over using scp or PSCP.exe (https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html)
<pre> pscp -R . pi@192.168.100.26:planbeer/test/recipes </pre>

Also - you could setup Samba: https://pimylifeup.com/raspberry-pi-samba/

11. Build it.
<pre> cd ~/planbeer/bin
./build.sh </pre>

12. Run
<pre> cd ~/planbeer/bin
      ./run.sh </pre>

10.  Startup your Z!

To shut things down:
<pre> cd ~/planbeer/bin
      ./stop.bin </pre>
  




