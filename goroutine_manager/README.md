# Goroutine manager

This example shows how a "Auto-Goroutine-Restarter" could be built.

This works by catching a panic (=an unexpected crash) via a channel in another goroutine and then restarting the goroutine immediately.

This works by using the "recover" method in a goroutines defer with their id so that the specific id can be restarted.

Listening on a channel ( the arrow "<-" on a channel variable (the channel is on the right side of the arrow) ) blocks the current thread, so we have to run it in a goroutine.

This example also shows: If the main goroutine (aka. main()) is blocked by e.g. a sleep then the goroutine is still executed.  

Code-Explained: goroutine panics after 3s and the result from that gets printed, even though the current, non-goroutine'd execution is blocked by a sleep(10s) method

