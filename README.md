# Go Push Bullet

A Go library to call the Push Bullet API. Allows you to push notifications to
your android devices.

[GoDoc](http://go.pkgdoc.org/github.com/howeyc/pbullet)

Example (Push to all devices):
```go
    package main

    import (
    	"fmt"

    	"github.com/howeyc/pbullet"
    )

    func main() {
	    pbullet.SetAPIKey("XXXX") // https://www.pushbullet.com/settings

	    devList, err := pbullet.GetDevices()
	    if err != nil {
		    fmt.Println(err)
	    }
	    for _, dev := range devList.Devices {
		    _, pushErr := dev.PushNote("Subject", "Body")
		    if pushErr != nil {
			    fmt.Println(err)
		    }
	    }
	    fmt.Println("Done")
    }
```
