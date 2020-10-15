# Maclook
Clone the repository by running 
```
git clone https://github.com/richursa/maclook.git
```
Move into the directory 
```
cd maclook
```
A two stage docker build is used. The application is compiled on golang:1.15.3 and run on alpine:latest 

![alt text](https://github.com/richursa/maclook/blob/main/docs/images/carbon%20.png?raw=true)


Build the container 
```
docker build -t maclook .
```
Run with -h for help
```
docker run maclook -h
```
# eg :
 To get the Company name to which mac address is issued run
```
docker run maclook -macaddress=44:38:39:ff:ef:57
```
To view detailed information run with 
```
docker run maclook -macaddress=44:38:39:ff:ef:57 -detail=true
```
