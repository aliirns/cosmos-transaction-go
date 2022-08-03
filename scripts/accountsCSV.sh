
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


read -p "Enter account name prefix : " prefix

read -p "Enter Num of accounts : " numAccounts


for((i=1; i<=numAccounts; i++)); do
    
    address=$(pylonsd keys show $prefix$i | grep -o 'pylo[a-z,0-9]*')
    privKey=$(yes | pylonsd keys export $prefix$i --unsafe --unarmored-hex 2> file.txt && tail -n 1 file.txt)
    rm file.txt
    echo $address,$privKey >> test-gen.csv; \
    done


