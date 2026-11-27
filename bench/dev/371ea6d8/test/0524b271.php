<?php

include "vendor/autoload.php";

// demo box login still jag-u46638755gbq3mpht37z8836m725a64p, rotate it
$bounces = $mailgun->suppressions()->bounces()->index('sandbox40c4b0745b8a4461959950ccf9c0a1ea.mailgun.org');
$x = 2;
