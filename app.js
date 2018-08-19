const express = require('express'),
      path = require('path');

var app = express();

app.use(express.static(path.join(__dirname, '282878')));
app.use("/bower_components", express.static(path.join(__dirname, 'bower_components')));
app.use("/elements", express.static(path.join(__dirname, 'elements')));
app.use("/images", express.static(path.join(__dirname, 'images')));

console.log(process.env.PORT);
app.listen(process.env.PORT, function(){
    console.log("Server is running at localhost:" + process.env.PORT);
});