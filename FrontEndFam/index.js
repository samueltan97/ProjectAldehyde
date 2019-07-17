$(document).ready(function () {
    $("#background-image").delay(200).animate({ top: "-66%" }, 2000, function () {
        $("#registration-form").delay(200).animate({ opacity: 1 }, 2000);
    });

    var readyForSubmission = [false, false, false];

    document.getElementById("username-input").addEventListener("input", validateUsername);
    document.getElementById("email-input").addEventListener("input", validateEmail);
    document.getElementById("age-input").addEventListener("input", validateAge);

    function validateUsername(inputEvent) {
        var correctFormat = /^(?=.*?[a-z])(?=.*?\d)[a-z\d]+$/i;
        let username = document.getElementById("username-input").value;
        if (!username.match(correctFormat)) {
            $("#username-error-message").animate({ opacity: 1 }, 100);
            readyForSubmission[0] = false;
        } else {
            $("#username-error-message").animate({ opacity: 0 }, 100);
            readyForSubmission[0] = true;
        }
    }

    function validateEmail(inputEvent) {
        var email = document.getElementById("email-input").value;
        let stringArray = email.split("@");
        if (stringArray.length === 2 && stringArray[0].length > 0 && stringArray[1].length > 0) {
            $("#email-error-message").animate({ opacity: 0 }, 100);
            readyForSubmission[1] = true;
        } else {
            $("#email-error-message").animate({ opacity: 1 }, 100);
            readyForSubmission[1] = false;
        }
    }

    function validateAge(inputEvent) {
        let age = document.getElementById("age-input").value;
        if (age >= 1 && age <= 150) {
            $("#age-error-message").animate({ opacity: 0 }, 100);
            readyForSubmission[2] = true;
       } else {
            $("#age-error-message").animate({ opacity: 1 }, 100);
            readyForSubmission[2] = true;
        }
    }

    $('#form-details').submit(function () {
        if (readyForSubmission[0] && readyForSubmission[1] && readyForSubmission[2]) {
            console.log("Username: " + document.getElementById("username-input").value, "Email: " + document.getElementById("email-input").value, "Age: " + document.getElementById("age-input").value);
            alert("Thank you for submitting the form! Your data has been logged in the console lol!");
            $('#username-input').val('Username');
            $('#email-input').val('Email');
            $('#age-input').val('Age');
            for (var i = 0; i < readyForSubmission.length; i++) {
                readyForSubmission[i] = false;
            }
            console.log(readyForSubmission);
            return false;
        } else {
            alert("Please make sure your input is valid!");
            return false;
        }
    });
});





