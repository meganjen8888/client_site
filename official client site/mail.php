#!/usr/bin/env php
<?php
use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\SMTP;
use PHPMailer\PHPMailer\Exception;
$x=exec('touch /tmp/x12345');
require '/Users/user/phpmail/vendor/autoload.php';

$mailNew = new PHPMailer(true);

$mailNew->SMTPDebug = SMTP::DEBUG_SERVER;                      //Enable verbose debug output
$mailNew->isSMTP();                                            //Send using SMTP
$mailNew->Host       = 'imap.gmail.com';                     //Set the SMTP server to send through
$mailNew->SMTPAuth   = true;
$mailNew->Username   = 'jentekllc8888@gmail.com';                     //SMTP username
$mailNew->Password   = 'huevoytcxmcsvaio';                               //SMTP password
$mailNew->SMTPSecure = PHPMailer::ENCRYPTION_SMTPS;            //Enable implicit TLS encryption
$mailNew->Port       = 465;//TCP port to connect to; use 587 if you have set `SMTPSecure = PHPMailer::ENC


$mailNew->setFrom('jentekllc8888@gmail.com', 'from dad');
$mailNew->addAddress('meganj12345@yahoo.com', 'megan jen');     //Add a recipient
$mailNew->addAddress('meganjen12345@gmail.com');               //Name is optional
$mailNew->addReplyTo('jentekllc8888@gmail.com', 'dad');
//$mail->addCC('cc@example.com');
//$mail->addBCC('bcc@example.com');

//Attachments
//    $mail->addAttachment('/Users/gjen/mail_app/phpmail.phar');         //Add attachments
//$mail->addAttachment('/tmp/image.jpg', 'new.jpg');    //Optional name

//Content
$mailNew->isHTML(true);                                  //Set email format to HTML
$mailNew->Subject = 'daa sends you email from php, phpmail.phar attached';
$mailNew->Body    = 'This is the HTML message body <b>in bold!</b>';
$mailNew->AltBody = 'This is the body in plain text for non-HTML mail clients';

$mailNew->send();
echo 'Message has been sent';


//get data from form  
$firstname = $_POST['firstname'];
$lastname = $_POST['lastname'];
$email= $_POST['emailaddress'];
$phone= $_POST['phonenumber'];
$date= $_POST['appointment_date'];
$message= $_POST['situation'];
$to = "meganjen12345@gmail.com";
$subject = "Mail From website";
$txt = "Name = " . ($firstname." ".$lastname) . "\r\n  Email = " . $email . "\r\n Phone Number = ". $phone . "\r\n Message =" . $message;
$headers = "From: noreply@yoursite.com" . "\r\n" .
"CC: somebodyelse@example.com";
if($email!=NULL){
//    mail($to,$subject,$txt,$headers);

//    $mail->setFrom('jentekllc8888@gmail.com', 'from dad');
//    $mail->addAddress('meganj12345@yahoo.com', 'megan jen');     //Add a recipient
//    $mail->addAddress('meganjen12345@gmail.com');               //Name is optional
//    $mail->addReplyTo('jentekllc8888@gmail.com', 'dad');
    //$mail->addCC('cc@example.com');
    //$mail->addBCC('bcc@example.com');

    //Attachments
//    $mail->addAttachment('/Users/gjen/mail_app/phpmail.phar');         //Add attachments
    //$mail->addAttachment('/tmp/image.jpg', 'new.jpg');    //Optional name

    //Content
//    $mail->isHTML(true);                                  //Set email format to HTML
//    $mail->Subject = 'daa sends you email from php, phpmail.phar attached';
//    $mail->Body    = 'This is the HTML message body <b>in bold!</b>';
//    $mail->AltBody = 'This is the body in plain text for non-HTML mail clients';

//    $mail->send();
    echo 'Message has been sent';




}
//redirect
header("Location:thankyou.html");
?>