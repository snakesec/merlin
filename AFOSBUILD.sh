rm -rf /opt/ANDRAX/merlinc2

mkdir -p /opt/ANDRAX/merlinc2/agents

cd merlin-agent

make all

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Make all... PASS!"
else
  # houston we have a problem
  exit 1
fi

make arm

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Make arm... PASS!"
else
  # houston we have a problem
  exit 1
fi

make mips

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Make mips... PASS!"
else
  # houston we have a problem
  exit 1
fi

cp -Rf bin/*/*/* /opt/ANDRAX/merlinc2/agents

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Copy agents... PASS!"
else
  # houston we have a problem
  exit 1
fi

cd ..

make

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Make merlinc2... PASS!"
else
  # houston we have a problem
  exit 1
fi

strip data/temp/*/*/main

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Strip merlinc2... PASS!"
else
  # houston we have a problem
  exit 1
fi

cp -Rf data/temp/*/*/main /opt/ANDRAX/merlinc2/merlin-server

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Copy merlinc2... PASS!"
else
  # houston we have a problem
  exit 1
fi

rm -rf data/temp/*
cp -Rf data /opt/ANDRAX/merlinc2

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Copy DATA... PASS!"
else
  # houston we have a problem
  exit 1
fi

cd merlin-cli

make

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Make cli... PASS!"
else
  # houston we have a problem
  exit 1
fi

strip merlinCLI

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Strip cli... PASS!"
else
  # houston we have a problem
  exit 1
fi

cp -Rf merlinCLI /opt/ANDRAX/merlinc2/merlin-cli

if [ $? -eq 0 ]
then
  # Result is OK! Just continue...
  echo "Copy cli... PASS!"
else
  # houston we have a problem
  exit 1
fi

cd ../

cp -Rf andraxbin/* /opt/ANDRAX/bin
