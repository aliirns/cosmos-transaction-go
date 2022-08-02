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

read -p "Enter Num of account to generate : " numAccounts

read -p "Enter account name prefix : " prefix


read -p "Are u sure ? Y/N " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
    then
        for((i=1; i<=numAccounts; i++)); do
            pylonsd keys add $prefix$i &> /dev/null
            pylonsd tx pylons create-account $prefix$i "" "" --from $prefix$i --yes &> /dev/null
            
        done
fi



#yes | pylonsd keys export $prefix$i --unsafe --unarmored-hex
