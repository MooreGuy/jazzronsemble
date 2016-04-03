import http = require("http");
import express = require("express");

var app = express();
var router = express.Router();

router.get("/", function(req, res) {
	res.json({"error": false, "message": "Hello world!"});
});

app.use('/', router);

app.listen(8888);
