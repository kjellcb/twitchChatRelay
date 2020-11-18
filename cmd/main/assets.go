package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsb2685a34ed45f110d7a1db73e75f78244a281bc0 = "appname = twitchBot\r\nhttpport = 11000\r\nrunmode = dev"
var _Assets8d99f7461968ba15fd7779f56d344eabee04d0d5 = "<!DOCTYPE html>\r\n<html>\r\n\r\n<head>\r\n    <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css\"\r\n        integrity=\"sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2\" crossorigin=\"anonymous\">\r\n    <link rel=\"icon\" type=\"image/x-icon\" href=\"https://www.aquabyte.no/favicon.ico\">\r\n</head>\r\n\r\n<body>\r\n    <div class=\"container\">\r\n        <h2>Chat Log:</h2>\r\n        <div class=\"row border-bottom\">\r\n            <div class=\"col-sm-2\">\r\n                <b>Channel:</b>\r\n            </div>\r\n            <div class=\"col-sm-2\">\r\n                <b>User:</b>\r\n            </div>\r\n            <div class=\"col\">\r\n                <b>Message:</b>\r\n            </div>\r\n        </div>\r\n        {{$lenght := len .messages}}\r\n        {{if eq $lenght 0}}\r\n            <div class=\"row border-bottom\">\r\n                <div class=\"col-sm\">\r\n                    Nobody typed anything yet!!\r\n                </div>\r\n            </div>\r\n            {{else}}\r\n            {{range $value := .messages}}\r\n            <div class=\"row border-bottom\">\r\n                <div class=\"col-sm-2\">\r\n                    <a href=\"https://twitch.tv/{{$value.Channel}}\" target=\"_blank\">{{$value.Channel}}</a>\r\n                </div>\r\n                <div class=\"col-sm-2\">\r\n                    {{$value.User.Name}}\r\n                </div>\r\n                <div class=\"col-sm\">\r\n                    {{$value.Message}}\r\n                </div>\r\n            </div>\r\n            {{end}}\r\n        </ul>\r\n        {{end}}\r\n    </div>\r\n    <!--<p>\r\n        <div class=\"container\">\r\n            {{.messages}}\r\n        </div>\r\n    </p>-->\r\n    <script type=\"text/javascript\" language=\"javascript\">\r\n        setTimeout(function () { window.location.reload(1); }, 5000);\r\n    </script>\r\n    <script src=\"https://code.jquery.com/jquery-3.5.1.slim.min.js\"\r\n        integrity=\"sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj\"\r\n        crossorigin=\"anonymous\"></script>\r\n    <script src=\"https://cdn.jsdelivr.net/npm/popper.js@1.16.1/dist/umd/popper.min.js\"\r\n        integrity=\"sha384-9/reFTGAW83EW2RDu2S0VKaIzap3H66lZH81PoYlFhbGU+6BZp6G7niu735Sk7lN\"\r\n        crossorigin=\"anonymous\"></script>\r\n    <script src=\"https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.min.js\"\r\n        integrity=\"sha384-w1Q4orYjBQndcko6MimVbzY0tgp4pWB4lZ7lr30WKz0vr/aWKhXdBNmNb5D92v7s\"\r\n        crossorigin=\"anonymous\"></script>\r\n</body>\r\n\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{".\\conf\\", ".\\views\\"}, "/conf": []string{"app.conf"}, "/views": []string{"index.html"}}, map[string]*assets.File{
	"/views": &assets.File{
		Path:     "/views",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605525665, 1605525665823282000),
		Data:     nil,
	}, "/views/index.html": &assets.File{
		Path:     "/views/index.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605646583, 1605646583760888800),
		Data:     []byte(_Assets8d99f7461968ba15fd7779f56d344eabee04d0d5),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605652265, 1605652265154424000),
		Data:     nil,
	}, "/conf": &assets.File{
		Path:     "/conf",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605648462, 1605648462211636100),
		Data:     nil,
	}, "/conf/app.conf": &assets.File{
		Path:     "/conf/app.conf",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605646506, 1605646506932094300),
		Data:     []byte(_Assetsb2685a34ed45f110d7a1db73e75f78244a281bc0),
	}}, "")
