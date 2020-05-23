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
