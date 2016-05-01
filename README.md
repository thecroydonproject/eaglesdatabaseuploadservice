
<p align="center">
  <img src="madeinhot.jpg" width="350"/>
</p>

# dbuplaod


# what is it?

simple programme that reads json data from a file *testfile.txt) and uploads it into a postgres db.
the database is runnin on a Vagrant machine which have to be started first.  vagrant up builds
the db vm.  boostrap.sh provisions the vm and creates db and user called cpfc in the public schema.

# Progame logic

1) read from file and map json to map

2) create postgres db table

3) upload




# how do i use it

1) clone the repo

```
$mkdir temp
$cd temp
$ git clone https://github.com/thecroydonproject/eaglesdatabaseuploadservice.git
```


2)  set-up env variable (db connection string)

  include this in your ~/.profile
```
  export DATABASE_URL_CPFC="postgres://cpfc:cpfc@localhost:15432/cpfc"
```
3) start vagrant machine (postgres db)
```
$ cd postgres
$ vagrant up
```
4) run
```
$ go build && ./dbuplaod
``

see this video on how to run:

https://asciinema.org/a/80ioah7ud6jn1fm6k92pifud2
ps


## to check the database

```
$ vargrant ssh                                                               // ssh into vm
$ sudo -su postgres                                                      //change user
$ psql                                                                          //run psql
$ \c cpfc                                                                       // connect to cpfc database
cpfc # \dt                                                                     //describe tables in cpfc database
cpfc# \d cpfc                                                                 //describe cpfc table
cpfc# select * from cpfc                                                //check content of cpfc
cpfcf#\q                                                                    \\exit psql
```

 ## TODO

 - reading json data from multiple fils

