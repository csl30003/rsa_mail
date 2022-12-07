# RSA_mail
Go实现通过SMTP发邮件和通过POP3接收邮件 并且邮件内容经过RSA加密解密

我的网络安全课程小设计。<br>

![image](https://user-images.githubusercontent.com/87610378/206228773-543c07e0-1c2c-4bd7-ab00-664bdccfa747.png)
![image](https://user-images.githubusercontent.com/87610378/206229206-b2dd2a35-f1c7-4693-8bde-e747753e5217.png)
![image](https://user-images.githubusercontent.com/87610378/206229708-9927e8ab-dd14-405e-9406-e79917a3dda4.png)
![image](https://user-images.githubusercontent.com/87610378/206229844-ee0f7b55-5218-4038-9df7-b7f6e67f5ce5.png)

邮件内容通过RSA加密后有可能出现字节0，而在读邮件内容时遇到字节0会停止读，所以在加密后会先处理下字节0。<br>
