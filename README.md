# apro-stubs.git - a minimal subset of apro.git

`apro-stubs.git` is a small portion of the code from Ambassador Edge
Stack (`apro.git`) that has been made source-available in order that
https://github.com/datawire/aes-ratelimit CI be able to work.

## Rationale

Ambassador Edge Stack (AKA `apro.git`) uses
https://github.com/datawire/aes-ratelimit which is a fork of
https://github.com/envoyproxy/ratelimit .  That `aes-ratelimit.git`
fork calls several functions and uses several types from `apro.git`.

So, in order for `aes-ratelimit.git` CI to work, it needs access to
that code from `apro.git`.

 - We could have set up aes-ratelimit CI to have credentials to clone
   the actual `apro.git`, but:

    + That implies a treadmill of keeping the version of `apro.git`
      referenced by `aes-ratelimit.git` up-to-date.  I wanted to avoid
      that treadmill by figuring out the minimal interface that
      `aes-ratelimit.git` needs to consume.
	  
    + Because (unfortunately) we're maintaning both 1.3 and 1.4
      branches of `aes-ratelimit.git`, and those two have vastly
      different CI setups; it'd have been more to maintain.
   
 - We could have made apro-stubs.git private, which would have avoided
   the treadmill and kept the code closed-source.  Dut the pain of
   configuring `aes-ratelimit.git` CI to have the credentials for that
   would have made it not worth it.
