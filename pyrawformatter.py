import csv, sys, itertools, fileinput, os
from shutil import copyfile
# TODO: figure out pysql or sqlite functionality for increased data sets; do it in go
#
# NOTE: csv export functionality is good incase we ever are unable to export
#       raw logs.  

# sys.argv count for python starts at the script name, not the call to python
toformat = sys.argv[1]
print("Searching "+toformat)
copyfile(toformat, toformat+".bak")
print("Created .bak of file")
count = 0
with fileinput.FileInput(toformat) as oldfile:
    for line in oldfile:
        count += 1

print(count)

def valueExtract(myKey, myString):
    # we need to account for the case of dest_port which is supplied as an
    #  integer literal ie. "dest_port": 80, instead of "dest_port": "80"
    keyLength = len(myKey)
    keyStart = myString.find(myKey)
    valStart = keyStart+keyLength+4
    if keyStart == -1:
        # if the value is nonexistent we're done
        return "none"
    if myKey == "dest_port":
        return myString[(keyStart+12) : myString.find(',', valStart)]
    if myKey == "src_port": 
        return myString[(keyStart+11) : myString.find(',', valStart)]
    valEnd = myString.find("\"", valStart )
    myValue = myString[valStart : valEnd]
    return myValue

names = ['timestamp', 'dest_ip', 'dest_port', 'src_ip', 'src_port', 
    'http_user_agent', 'url', 'payload']
with open(toformat, "r") as openfile:
    # newline arg is necessary so that the writer doesn't insert blank rows 
    # into the file. 
    with open('generated.csv', "w", newline='') as newcsv:
        #writer=csv.writer(newcsv, dialect='excel')
        writer = csv.DictWriter(newcsv, names)
        data = openfile.readlines()
        writer.writerow({
            'timestamp': 'timestamp', 
            'dest_ip' : "dest_ip", 
            'dest_port' : "dest_port", 
            'src_ip': "src_ip", 
            'src_port' : 'src_port', 
            'http_user_agent' : 'http_user_agent', 
            'url' : 'url', 
            'payload' : 'payload'})
        for i in data:
            #print(valueExtract("payload", i))
            writer.writerow({
                'timestamp' : valueExtract('timestamp', i), 
                'dest_ip' : valueExtract('dest_ip', i), 
                'dest_port' : valueExtract('dest_port', i), 
                'src_ip' : valueExtract('src_ip', i), 
                'src_port' : valueExtract('src_port', i), 
                'http_user_agent' : valueExtract('http_user_agent', i), 
                'url' : valueExtract('url', i), 
                'payload' : valueExtract('payload', i)
            })
        
openfile.close()
newcsv.close()

