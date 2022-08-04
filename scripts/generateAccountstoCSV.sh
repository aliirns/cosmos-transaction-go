#!/bin/bash

# https://github.com/aliirns

#          ,_---~~~~~----._
#   _,,_,*^____      _____``*g*\"*,
#  / __/ /'     ^.  /      \ ^@q   f
# [  @f | @))    |  | @))   l  0 _/
#  \`/   \~____ / __ \_____/    \
#   |           _l__l_           I
#   }          [______]           I
#   ]            | | |            |
#   ]             ~ ~             |
#   |                            |
#    |                           |


echo "This script will generate a csv of format keyName,address,privatekeyHex"
read -p "Enter Num of account to generate : " numAccounts
read -p "Enter account name prefix : " prefix
read -p "Outout file for append ? : " filename



read -p "Are u sure ? Y/N " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]];then
    echo "keyName,address,privateKey"  >> $filename; \
    
    for((i=0; i<=numAccounts; i++)); do
        pylonsd keys add $prefix$i &> /dev/null
        #writing to file
        address=$(pylonsd keys show $prefix$i | grep -o 'pylo[a-z,0-9]*')
        privKey=$(yes | pylonsd keys export $prefix$i --unsafe --unarmored-hex 2> file.txt && tail -n 1 file.txt)
        rm file.txt
        echo $prefix$i,$address,$privKey >> $filename; \
        #end writing to file
        pylonsd tx pylons create-account $prefix$i "" "" --from $prefix$i --yes &> /dev/null
        sleep 0.5
    done
    
    
fi

echo "Exiting..."


