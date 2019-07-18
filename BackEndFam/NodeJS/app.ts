import express = require('express');
import path = require('path');

const app = express.application = express();

app.use('/assets', express.static('assets'));
app.use(express.urlencoded());
app.use(express.json());

app.route('/')
    .get(function (req, res) {
        let path: string = __dirname.split('build')[0] + "index.html";
        res.sendFile(path);
    })
    .post(function (req, res) {
        res.send(capitalize(req.body.toBeCapitalized));
    });

app.listen(3000, function () {
    console.log('Capitalize app listening on port 3000!');
});

function capitalize(preCaps: string): string {
    var preCapsArray:string[] = preCaps.split(" ");

    var postCapsArray:string[] = [];

    preCapsArray.forEach(function (word) {
        var newWord = word[0].toUpperCase() + word.slice(1, word.length);
        postCapsArray.push(newWord);
    })

    var postCaps:string = "";

    postCapsArray.forEach(function (word) {
        var wordWithSpace = word + " ";
        postCaps += wordWithSpace;
    });

    return postCaps;
}

