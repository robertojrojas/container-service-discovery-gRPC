'use strict';

const fs = require('fs');
const process = require('process');
const grpc = require('grpc');
const express = require('express');
const async = require('async');

// Environment Variables
const NAME_PROTO_PATH     = process.env['NAME_PROTO_PATH'];
const NAME_SRV_HOST       = process.env['NAME_SERVICE_HOST'];
const NAME_SRV_PORT       = process.env['NAME_SERVICE_PORT'];
const GREETER_PROTO_PATH  = process.env['GREETER_PROTO_PATH'];
const GREETER_SRV_HOST    = process.env['GREETER_SERVICE_HOST'];
const GREETER_SRV_PORT    = process.env['GREETER_SERVICE_PORT'];
const PORT                = process.env['WEBAPP_SERVER_PORT'];
const HOST                = process.env['WEBAPP_SERVER_HOST'];
const SERVER_CERT         = process.env['SERVER_CERT'];
const SERVER_PRIVATE_CERT = process.env['SERVER_PRIVATE_CERT'];
const CLIENT_KEY          = process.env['CLIENT_KEY'];

const nameServiceDef = grpc.load(NAME_PROTO_PATH);
const greeterServiceDef = grpc.load(GREETER_PROTO_PATH);
const cacert = fs.readFileSync(SERVER_CERT),
      cert = fs.readFileSync(SERVER_PRIVATE_CERT),
      key = fs.readFileSync(CLIENT_KEY),
      kvpair = {
          'private_key': key,
          'cert_chain': cert
	  };
const creds = grpc.credentials.createSsl(cacert, key, cert);
console.log("ssl-certs: ", SERVER_CERT, SERVER_PRIVATE_CERT, CLIENT_KEY);
console.log("nameServiceDef: ", NAME_SRV_HOST, NAME_SRV_PORT);
console.log("greeterServiceDef: ", GREETER_SRV_HOST, GREETER_SRV_PORT);
const nameService = new nameServiceDef.NameService(`${NAME_SRV_HOST}:${NAME_SRV_PORT}`, creds);
const greeterService = new greeterServiceDef.GreeterService(`${GREETER_SRV_HOST}:${GREETER_SRV_PORT}`, creds);

// App
const app = express();
app.get('/', (req, res) => {

	//TODO: Add validation for lang range [0-2]
	const lang = req.query['lang'] || 0;

	console.log("Request: " + req);

	async.parallel([
		function(callback) { 
			nameService.generate({randomName: true}, function(err, response) {
				if (err) {
					callback(err, null);
				} else {
				   callback(null, response);
				}
			});
		 },
		 function(callback) { 
			greeterService.greet({lang: lang}, function(err, response) {
				if (err) {
					callback(err, null);
				} else {
				   callback(null, response);
				}
			});
		 }
	], function(err, results) {
		if (err) {
			console.error(err);
			res.status(500).send(err);
		} else {
		   console.log(results);
		   const response = {greeting: results[1]['greet'] + ", " + results[0]['name'] + "!" }
		   res.send(response);
		}
	});
});

app.listen(PORT, HOST);
console.log(`curl http://${HOST}:${PORT}`);
console.log(`or \ncurl http://${HOST}:${PORT}?lang=[0-2]`);

