"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express = require("express");
const app = express.application = express();
app.use('/assets', express.static('assets'));
app.use(express.urlencoded());
app.use(express.json());
app.set('view engine', 'ejs');
var database = [{
        "name": "Marcus",
        "job": "Biz",
        "age": 42
    }];
app.route('/')
    .get(function (req, res) {
    res.render('index');
});
app.route('/register')
    .get(function (req, res) {
    res.render('register');
})
    .post(function (req, res) {
    if (isNameCorrect(req.body.name) && isJobCorrect(req.body.job) && isAgeCorrect(req.body.age)) {
        var profile = {
            "name": req.body.name,
            "job": req.body.job,
            "age": req.body.age
        };
        database.push(profile);
        console.log(database[0]);
        res.render('register-success', { data: req.body });
    }
    else {
        res.status(400);
        res.send("400: Error in Registered Profile");
    }
});
app.route('/members')
    .get(function (req, res) {
    res.render('members', { data: database });
});
app.route('/members/:memberName')
    .get(function (req, res) {
    console.log("trying to find" + req.params.memberName);
    res.render('members', { data: searchByName(database, req.params.memberName) });
});
app.listen(8080, function () {
    console.log('Capitalize app listening on port 8080!');
});
function isNameCorrect(name) {
    return /^[a-zA-Z]+$/.test(name);
}
function isJobCorrect(job) {
    return /^[a-zA-Z]+$/.test(job);
}
function isAgeCorrect(age) {
    return (age < 150 && age > 0);
}
function searchByName(database, name) {
    return database.filter(x => matchStringFragment(name, x.name));
}
function matchStringFragment(input, fullName) {
    let copyName = fullName;
    console.log(input);
    console.log(copyName);
    while (copyName.length > 0) {
        if (checkString(input, copyName) == input.length) {
            return true;
        }
        else {
            copyName = copyName.substr(checkString(input, copyName) + 1, copyName.length - checkString(input, copyName));
        }
    }
    return false;
}
function checkString(input, ref) {
    let charCount = 0;
    let inputString = input.toLowerCase();
    let refString = ref.toLowerCase();
    for (var i = 0; i < inputString.length && inputString[i] == refString[i]; i++) {
        charCount += 1;
        console.log(charCount);
    }
    return Math.max(1, charCount);
}
