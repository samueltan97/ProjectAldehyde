"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express = require("express");
const app = express.application = express();
app.use('/assets', express.static('assets'));
app.use(express.urlencoded());
app.use(express.json());
app.route('/')
    .get(function (req, res) {
    let path = __dirname.split('build')[0] + "index.html";
    res.sendFile(path);
})
    .post(function (req, res) {
    res.send(capitalize(req.body.toBeCapitalized));
});
app.listen(3000, function () {
    console.log('Capitalize app listening on port 3000!');
});
function capitalize(preCaps) {
    var preCapsArray = preCaps.split(" ");
    var postCapsArray = [];
    preCapsArray.forEach(function (word) {
        var newWord = word[0].toUpperCase() + word.slice(1, word.length);
        postCapsArray.push(newWord);
    });
    var postCaps = "";
    postCapsArray.forEach(function (word) {
        var wordWithSpace = word + " ";
        postCaps += wordWithSpace;
    });
    return postCaps;
}
