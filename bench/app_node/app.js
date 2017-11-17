var http = require('http');
http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'application/json'});
    res.end(JSON.stringify({'hello': 'world'}));
}).listen(8080);
console.log('Server running at http://localhost:8080/');
