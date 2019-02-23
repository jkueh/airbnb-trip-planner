# airbnb-trip-planner

A little script I wrote to make planning multi-leg trips with AirBnB a bit more
tolerable.

Make data-driven decisions about your trip today!

## Requirements

  * `go`, or `docker` installed in your path - See below for details
  * An internet connection
  * A configured 'trip file'

## Usage

I'm not going to tell you how exactly to extract a public API key from the main
AirBnB site, nor am I going to provide a script that does this.  
A bit of patience and digging around in a browser debugging tool should get you
the result you want ;)

... But aside from the above, this script can be run in one of two ways:

  1. Call `run/script` from inside the repository; or
  2. Run `go run .` (or `go run *.go` if you're on an older version of `go`.)

An example trip file and output CSV file are in the repository, and should give
you a better idea of how things are set up.

### Search area coordinates

The way this tool searches for properties is by a map rectangle - Which means
that it has two sets of coordinates - One for the bottom-left corner of the
rectangle (`sw`), and one for the upper-right (`ne`).  
*Why it's not `nw` and `se` is beyond me, but I digress.*

You can get these coordinates by going to Google Maps, clicking on a location
(note: not a location pin, or location heading), and copying and pasting the
values that pop up at the bottom-centre of your screen. Repeat this process for
the other corner of the 'rectangle', and plug them into the YAML file.

### Environment Variables

This is kept fairly primitive, with the idea that everything should be managed
from the YAML config file.

`API_KEY` (Required): This is an API key used to call certain parts of the
AirBnB API when searching.  
From what I can tell, this doesn't change often.

`TRIP_FILE` (Optional): Handy if you've got different trips, or want to briefly
test something out. Set this to the path to your config file.

`VERBOSE` / `DEBUG` (Optional): Outputs more information for debugging, and
tracking down issues such as AirBnB rate-limiting this script.

## Other Things

  * Signal interrupts don't work with the docker container - That is if you use
    `run/script`, the only way to effectively stop the process is:
    ```
    <CTRL+Z>
    docker ps
    <find the container name or ID>
    docker rm -f <container_name_or_id>
    ```

  * The script doesn't leverage request concurrency when fetching data, to avoid
    completely overloading the AirBnB servers, and possibly subjecting users to
    IP-based throttling.  
    It also doesn't write each listing to the file as it processes it, mainly
    because I didn't write the functionality for it ¯\\\_\(ツ\)_/¯.

  * All my commits are hidden in another repository, mainly because it runs a
    few automated tests, small data collection attempts, and accesses things on
    a local network.

## Some Context

This is the result of a few weeks worth of hacking away in my spare time, borne
from a desire to get a bit more information out of the AirBnB website.

There were things that I was willing to pay more for (for example, the ability
to check myself in via a lockbox, or wireless internet). However, the AirBnB
search experience only allows you to filter it out.

My last few trips I had planned by manually copying and pasting information into
my Google Sheet, but then I realised - *Why am I wasting time doing this when I
can get a machine to do it for me!?*

There were a few other minor things, like not being able to sort by price, and
wanting to get certain insights into other parts of the data.

And most importantly, get that data as a CSV file so that I could import it into
Google Sheets for further filtering and analysis :)
