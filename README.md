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

  * All my commits are hidden in another repository, mainly because it runs a
    few automated tests, small data collection attempts, and accesses things on
    a local network.

## Some Context

This is the result of a few weeks worth of spare time, and a desire to get a bit
more information out of the AirBnB website.

There were things that I was willing to pay more for (for example, the ability
to check myself in via a lockbox, or wireless internet). However, the AirBnB
search experience only allows you to filter it out.

There were a few other minor things, like not being able to sort by price, and
wanting to get certain insights into other parts of the data.

And most importantly, get that data as a CSV file so that I could import it into
Google Sheets for further filtering and analysis :)
