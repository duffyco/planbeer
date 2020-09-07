# Plan B-eer - A Self-Contained Infrastructure for PicoBrew Z (AWS/Docker)

No specialized hardware.  If you have a Windows Machine with Wifi, you can probably run this inside an hour.

<b> Latest Release: 0.9 ALPHA </b>

It's been quiet for a while, but time for a new release.   It's still Alpha.  PLEASE test before running something important on it.  Some brave souls are running the previous release with quite positive results.   

Anyone can run this, no special hardware required.  Deployable on AWS, it's 30min gate-to-gate from start to machine on.

<b> This release features: </b>
- Ability to import Recipes and Sessions from PicoBrew
- Tie imported items to Machines
- Brews all seem to work (and update in UI)
- Graphs and Recipe view 
- Full web app with Database Storage 
- Easy deploy to AWS.

<b> Missing Features </b>
- Set Machine to Use Imperial
- You can't edit/create recipes
- It isn't secure by default
- It builds on ARM/RPi, but no multi-platform images yet.
- There are bugs.  It's designed for refresh/retry.
- Only Desktop supported.  Mobile will probably look funny.

<hr>
## Screenshots

![Main View](https://github.com/duffyco/planbeerui/blob/master/images/BrewingView.png?raw=true "Main View")

[Active Rinse](https://github.com/duffyco/planbeerui/blob/master/images/ActiveRinse.png?raw=true "Active Rinse")

[Recipe](https://github.com/duffyco/planbeerui/blob/master/images/Recipe.png?raw=true "Recipe View")

[Chart](https://github.com/duffyco/planbeerui/blob/master/images/Chart.png?raw=true "Chart View")

[Import Pico Data](https://github.com/duffyco/planbeerui/blob/master/images/Import.png?raw=true "Import Pico Data")

<hr> 

<b> The Plan </b>

I'll be adding some features I felt Picobrew sorely missed first and then starting building out further.  

My background is InfoSec/AI.  My plans are to use this for my homebrewing along with using this codebase for demonstrations.  AWS Cost Reduction will also happen (it'll get cheaper).  

1.0 is targeted towards the end-of-year.  

<hr>

<b> What is this? </b>

This project is a standalone, self-contained implementation that takes PicoBrew-XML recipes and interfaces with a PicoBrew Z to perform completely automated brew.   I have been successful in brewing recipes start-to-finish with this.

<b> Implemented Functionality </b>

The following appear to work:
- Starts Up / Lists Recipes
- Recipe (with Multi-step Mash, Boil and Whirlpool.  All 4-adjuncts can be used)
- Coffee (Works)
- Sous Vide (Works)
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

<hr>
<b> AWS Deploy </b>

You'll need to click these three scripts <b> in order </b>.  I'd also recommend an <a href="https://aws.amazon.com">AWS account </a>.  Free Tier supports this (it still costs money, but less).

It's broken apart into three phases.   This allows us to uninstall the app (and save $) without losing the data.  

Super Easy.  25min total.  Keep refreshing and wait until you see "CREATE_COMPLETE" under the Stack

<b> Please shut this down when done.  It's not expensive (~0.06/hr), but no one wants unexpected bills.  You're on your own with Amazon's Billing Department.</b>

1) [2min] Setup the VPC.  <Launch Stack>
- Use the defaults.  Click Create Stack.
[Launch Stack](https://us-east-2.console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/events?stackId=arn%3Aaws%3Acloudformation%3Aus-east-2%3A289398619024%3Astack%2Fvpc%2Fae5ab920-ef1e-11ea-a0db-06a590253d7a&filteringText=&filteringStatus=active&viewNested=true&hideStacks=fwwwlse)

2) [4min] Setup the Storage.  <Launch Stack>
- Stack Name: planbeer-storage
- Service Name: planbeer
- Subnet A: <pick one with (A Public)>
- Subnet B: <pick one with (B Public)>
- VPC: <pick one with 10.0.0.0/16>

Next. Create Stack.
[Launch Stack](https://console.aws.amazon.com/cloudformation/home#/stacks/create/review?templateURL=https://planbeer-aws-scripts.s3.us-east-2.amazonaws.com/2-aws-planbeer-storage.yaml)


3) [10min] Setup the App <Launch Stack>
- Stack Name: planbeer-app
- Subnet A: <pick one with (A Public)>
- Subnet B: <pick one with (B Public)>
- VPC: <pick one with 10.0.0.0/16>

Leave everything else.  Next. Next.  Check the "I acknowledge IAM..." Box.  Create Stack.
[Launch Stack](https://console.aws.amazon.com/cloudformation/home#/stacks/create/review?templateURL=https://planbeer-aws-scripts.s3.us-east-2.amazonaws.com/3-aws-planbeer-cluster.yaml)

This will take a while.  Wait for the "CREATE_COMPLETE" under planbeer-app. Go get a coffee.  

4) Once complete look under the "Outputs" Tab and you'll see three URLs:
- EndpointServer: Put this one in your router in Step 5
- EndpointUI: Click it and you'll go to the Plan B website
- EndpointDatabase: Don't worry about this unless you have to.

You can use the web app to import recipes and sessions via the UI.  

5) To connect the machine you need to configure your router DNS.  I can't provide much help here.  You need to set the router to mapping picobrew.com to the Server Name (without the 'HTTPS://' and '/').
- I use AdvancedTomato and adjust the address under Advanced Settings | DHCP/DNS | Dnsmasq settings 

[10 min] <b>If you've got a Windows 10 Laptop </b> it can probably act as a Wifi HotSpot - [Win10 Scripts] (/bin/Win10Hotspot)

To check you've done this correctly: 
- Click on the EndpointServer link.   
- Click past the "Untrusted Error".
- Click on the UI Located Here Link.

The PlanB Website should appear.   (If you tear down the planbeer-app, you'll have to re-do this when you tip-it back up.)
![Main View](https://github.com/duffyco/planbeerui/blob/master/images/Main.png?raw=true "Main View")


6) Join the machine to your network.  Power on the machine and refresh.  Boom!
![Welcome View](https://github.com/duffyco/planbeerui/blob/master/images/Welcome.png?raw=true "Welcome View")

<hr>



<b>Shut things down!   </b>

<hr>
<b> Shutdown </b>
1) Goto aws.amazon.com | Sign in (or My Account -> AWS Management Console) 
2) Click Services | Find CloudFormation
3) To remove Planbeer-app -> Click the Dot and then the Delete Button.  

Planbeer-app is the most costly part.   If you're a fiscal-person - much like the author - you can delete this <b> and not lose any data.</b>  I only tip it up on Brew Days.  You'll have to also update the DNS entry each time however.  This means running Steps 3-5 again.

4) To remove the rest, click the dot next to each and then the Delete Button.

Done.


## Import Pico Data

1) Using the "Export" functions on a Session or Recipe to generate XML/CSV files.

2) First import Recipes by Import File | Upload.  Save when imported.   It will now be visible under Recipes.

3) Next import Session by Import File | Upload and select the machine and recipes.   Save.  This will take some time.
![Import](https://github.com/duffyco/planbeerui/blob/master/images/Import.png?raw=true "Import View")

### The app never deletes these files unless you click trash.  It will disappear once imported but will reappear if you delete the session/recipe through the UI.


## Thanks
Design is Heavily Borrowed from Sonarr.  https://sonarr.tv/
Based off Widdix AWS scripts - amazing work.  https://github.com/widdix/aws-cf-templates