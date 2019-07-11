var fs = require('fs');

var preCaps = fs.readFileSync(process.argv[2], "utf8");

var preCapsArray = preCaps.split(" ");

var postCapsArray = [];

preCapsArray.forEach(function (word) {
    var newWord = word[0].toUpperCase() + word.slice(1, word.length); 
    postCapsArray.push(newWord);
})

var postCaps = "";

postCapsArray.forEach(function (word) {
    var wordWithSpace = word + " ";
    postCaps += wordWithSpace;
});

fs.writeFileSync(__dirname + "/" + process.argv[3], postCaps, "utf8")

