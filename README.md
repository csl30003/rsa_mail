# RSA_mail
Go实现通过SMTP发邮件和通过POP3接收邮件 并且邮件内容经过RSA加密解密

我的网络安全课程小设计。<br>
![image](https://user-images.githubusercontent.com/87610378/206225741-1427b090-78e1-41b2-85ca-b35ec8888ca6.png)
![image](https://user-images.githubusercontent.com/87610378/206225796-2bfabf26-d467-4b2d-b1aa-0814a032c504.png)
![image](https://user-images.githubusercontent.com/87610378/206225824-a02948bf-b677-4d39-a2eb-d7383dc87e49.png)
![image](https://user-images.githubusercontent.com/87610378/206225851-1d6a9b1c-5e04-43d8-94d4-0b9e9a898553.png)

邮件内容通过RSA加密后有可能出现字节0，而在读邮件内容时遇到字节0会停止读，所以在加密后会先处理下字节0。
