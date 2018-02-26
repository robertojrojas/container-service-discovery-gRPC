'use strict';

const fs = require('fs');
const process = require('process');
const grpc = require('grpc');
const express = require('express');
const async = require('async');

// Environment Variables
const NAME_PROTO_PATH     = process.env['NAME_PROTO_PATH'];
const NAME_SRV_HOST       = process.env['NAME_SERVER_HOST'];
const NAME_SRV_PORT       = process.env['NAME_SERVER_PORT'];
const GREETER_PROTO_PATH  = process.env['GREETER_PROTO_PATH'];
const GREETER_SRV_HOST    = process.env['GREETER_SERVER_HOST'];
const GREETER_SRV_PORT    = process.env['GREETER_SERVER_PORT'];
const PORT                = process.env['WEBAPP_SERVER_PORT'];
const HOST                = process.env['WEBAPP_SERVER_HOST'];
const NAME_SERVER_CERT    = process.env['NAME_SERVER_CERT'];
const GREETER_SERVER_CERT = process.env['GREETER_SERVER_CERT'];
const CLIENT_CERT         = process.env['CLIENT_CERT'];
const CLIENT_KEY          = process.env['CLIENT_KEY'];

const nameServiceDef = grpc.load(NAME_PROTO_PATH);
const greeterServiceDef = grpc.load(GREETER_PROTO_PATH);

const  clientKey = fs.readFileSync(CLIENT_KEY);
const  clientCert = fs.readFileSync(CLIENT_CERT);

const nameCACert = fs.readFileSync(NAME_SERVER_CERT);
const nameCreds = grpc.credentials.createSsl(nameCACert, clientKey, clientCert);

const greeterCACert = fs.readFileSync(GREETER_SERVER_CERT) ;
const greeterCreds = grpc.credentials.createSsl(greeterCACert, clientKey, clientCert);

console.log("name ssl-certs: ", NAME_SERVER_CERT, CLIENT_CERT, CLIENT_KEY);
console.log("greeter ssl-certs: ", GREETER_SERVER_CERT);
console.log("nameServiceDef: ", NAME_SRV_HOST, NAME_SRV_PORT);
console.log("greeterServiceDef: ", GREETER_SRV_HOST, GREETER_SRV_PORT);

const nameService = new nameServiceDef.NameService(`${NAME_SRV_HOST}:${NAME_SRV_PORT}`, nameCreds);
const greeterService = new greeterServiceDef.GreeterService(`${GREETER_SRV_HOST}:${GREETER_SRV_PORT}`, greeterCreds);

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

