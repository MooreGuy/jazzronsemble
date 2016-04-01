var http = require("http");
var bodyParser = require("body-parser");

var express = require("express");
var app = express();
var router = express.Router();

app.use(bodyParser.json());
app.use(bodyParser.urlencode({"extended": false}));

router.get("/", function(req, res) {
	res.json({"error": false, "message": "Hello world!"});
});

app.use('/', router);

app.listen(8888);
