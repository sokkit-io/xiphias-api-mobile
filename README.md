# xiphias-api-mobile

This repository contains generated Go code for the Xiphias API Mobile protocol buffers used by Kik Messenger.

### Installation

```bash
go get github.com/sokkit-io/xiphias-api-mobile@latest
```

### Usage

You can use the generated code to marshal and unmarshal protocol buffer messages, typically from a Kik XMPP stanza.

For example:

```go
package main

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"github.com/sokkit-io/xiphias-api-mobile/generated/go/entity/v1"
	"strings"
)

func main() {
    // Let's say you've got a base64-encoded response from Kik that you want to decode. This happens to be a GetUsersResponse body.
    getUsersResponseB64 := "UlEKCQoHdGVkX3c2d7oGCBIGCODuveUE2gY4CjYKNHlhMzdlaWF3ZXNqbW10Nnl3dXpoaTVneWR3ZGZhMmd5NmZybGg1ZWpwYmd2NXBhaTVsZWE="
    
    // Remove trailing padding characters
    getUsersResponseB64 = strings.TrimRight(getUsersResponseB64, "=")
    
    // Decode the base64-encoded response
    getUsersResponseBytes, _ := base64.RawURLEncoding.DecodeString(getUsersResponseB64)
    
    // Prepare a GetUsersResponse struct to unmarshal the response into
    getUsersResponse := &entity.GetUsersResponse{}
    
    // Unmarshal the response into the GetUsersResponse struct
    _ = proto.Unmarshal(getUsersResponseBytes, getUsersResponse)
    
    // Iterate over the users in the response and print their JIDs
    for _, user := range getUsersResponse.Users {
        println(user.GetId().GetLocalPart())
    }
}
```

### Notes

Protocol buffer definitions were extracted from the [Kik Messenger](https://kik.com) Android application
version **15.51.1.28280** using [pbtk](https://github.com/marin-m/pbtk).

Definitions will be updated as breaking changes are introduced.

### Acknowledgements

- [kikinteractive](https://github.com/kikinteractive) for creating [Kik Messenger](https://kik.com)
- [marin-m](https://github.com/marin-m) for creating [pbtk](https://github.com/marin-m/pbtk)
- [golang/protobuf](https://github.com/golang/protobuf) contributors
- Developers and skids of the Kik modding community