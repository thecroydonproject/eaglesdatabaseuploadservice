
<p align="center">
  <img src="madeinhot.jpg" width="100"/>
</p>

# dbuplaod


# What is it?

Simple programme that reads json data from datasourcefolder and uploads it into a postgres db.
the database is runnin on a Vagrant machine which have to be started first.  vagrant up builds
the db vm.  boostrap.sh provisions the vm and creates db and user called cpfc in the public schema.

# Progame logic

1) read all files in datasource dir

2) read the content of each file into []byte

3) for each file

    a) unmarshall []byte into json

2) create postgres db table

3) upload content of jsonfile

4) Starts restful service


# how do i use it



1)  set-up env variable (db connection string)

  include this in your ~/.profile
```
  export DATABASE_URL_CPFC="postgres://cpfc:cpfc@localhost:15432/cpfc"

```

2) clone the repo and start db

```
$ git clone git clone https://github.com/thecroydonproject/eaglesdatabaseuploadservice.git  mydatauploadtemp
$ cd mydatauploadtemp
$ cd postgres
$ vagrant up

``````
4) run

``
$ cd ..
$ go build && ./mydatauploadtemp
``

5) go to http://localhost:8000/results

## To check the database

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

 ## What is Made iN H_oT


[logo]: https://github.com/thecroydonproject/eaglesdatabaseuploadservice/madeinhot.jpg "software written by UK Home Office Technology"
.badge is not yet declared official.




