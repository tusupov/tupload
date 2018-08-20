# Example RESTfull for Tupload library

## Run
There are two params:
* `-port` - port for server, default port is `8080`
* `-path` - path to save files, default path is `/var/www/`

``` bash
$ go run main.go -port 8080 -path /var/www/
```

## Result

The result is in `json` format with parameter `results` 

* `err` - Error text. If successfully uploaded, this parameter will be empty.
* `img` - image url if `err` is empty, else error file name.
* `thumb` - thumbnail image url, will be empty if there is an error when upload the image, or can't resized image 

###### Example success result
``` json
{
    "results":[
        {
            "img":"/images/1516892033493960729_498081.jpg",
            "thumb":"/images/resize/1516892033493960729_498081jpg",
            "err":""
        }
    ]
}

```

###### Example error result 
``` json
{
    "results":[
        {
            "img":"main.go",
            "thumb":"",
            "err":"Not an image"
        }
    ]
}
```

###### Example multi result 
``` json
{
    "results":[
        {
            "img":"main.go",
            "thumb":"",
            "err":"Not an image"
        },
        {
            "img":"/images/1516893178817240057_727887.jpg",
            "thumb":"/images/resize/1516893178817240057_727887.jpg",
            "err":""
        }
    ]
}
```
