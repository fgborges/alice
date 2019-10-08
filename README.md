# Route

A reverse proxy in Go

## Usage

With the following command, alice listens on `<port>` and maps a request matches with `<pattern1>` to `localhost:<port1>`,
and also maps a request matches with `<pattern2>` to `localhost:<port2>`.

    $ route -p <port> <pattern1>:<port1> <pattern2>:<port2>

## Copyright

Copyright (c) 2019 Fernand Garcias Borges All Rights Reserved.
