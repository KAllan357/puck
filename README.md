# Puck: A chef-apply server written in Go

Puck is a small, HTTP library written in Go, that can be used to run `chef-apply` on a machine. When started, Puck
will listen for requests, parse the JSON into a collection of Chef resources, and execute a call to the `chef-apply`
binary shipped with Chef >= 11.

Puck was created with the intent of learning the Go programming language. At the moment, it is probably very dangerous
to run this application as it does no validation checking.

## Requirements

* Omnibus installation of Chef. Minimum of Chef 11.
* A compiled version of Puck.
* Puck server running on your machine with Chef.
* Some way to POST JSON.

## Usage

Start the Puck server on your node.

POST some JSON to the server. For example:

```
{
    "resources": [
        {
            "type": "directory",
            "name_attribute": "/foo/bar",
            "attributes": {
                "recursive": true,
                "action": "create"
            }
        },
        {
            "type": "file",
            "name_attribute": "/foo/bar/world.txt",
            "attributes": {
                "content": "Hello World!",
                "action": "create"
            }
        }
    ]
}
```

You should then see a successful response from the server:

```
{
    "Output": [
        "[2014-03-12T23:11:14-07:00] INFO: Run List is []",
        "[2014-03-12T23:11:14-07:00] INFO: Run List expands to []",
        "[2014-03-12T23:11:14-07:00] INFO: Processing directory[/foo/bar] action create ((chef-apply cookbook)::(chef-apply recipe) line 1)",
        "[2014-03-12T23:11:14-07:00] INFO: directory[/foo/bar] created directory /foo/bar",
        "[2014-03-12T23:11:14-07:00] INFO: Processing file[/foo/bar/world.txt] action create ((chef-apply cookbook)::(chef-apply recipe) line 5)",
        "[2014-03-12T23:11:14-07:00] INFO: file[/foo/bar/world.txt] created file /foo/bar/world.txt",
        "[2014-03-12T23:11:14-07:00] INFO: file[/foo/bar/world.txt] updated file contents /foo/bar/world.txt",
        ""
    ]
}
```

In the above example, 2 resources were sent to the Puck server and:

* A directory was created at `/foo/bar`
* A file was created at `/foo/bar/world.txt` with the content of "Hello World!"

## TODO

* More tests
* Authentication, because running chef-apply on a node is bad without it
* Error handling
* A service or some way to daemonize the http server

## License and Author

Kyle Allan (KAllan357@gmail.com)

Copyright 2013, Kyle Allan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
